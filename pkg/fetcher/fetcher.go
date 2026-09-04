package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	BaseURL    = "https://web.whatsapp.com"
	UserAgent  = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36"
	MaxWorkers = 8
)

var (
	scriptSrcRe      = regexp.MustCompile(`<script[^>]+src=["']([^"']+)["']`)
	preloadRe        = regexp.MustCompile(`<link[^>]+(?:rel=["'](?:preload|modulepreload)["'][^>]+as=["']script["']|as=["']script["'][^>]+rel=["'](?:preload|modulepreload)["'])[^>]+href=["']([^"']+)["']`)
	manifestRe       = regexp.MustCompile(`assets-manifest-([0-9.]+)\.json`)
	clientRevisionRe = regexp.MustCompile(`client_revision\\?":([0-9.]+)`)
	versionStrRe     = regexp.MustCompile(`(?:appVersion:|VERSION_STR=)"([0-9.]+)"`)
)

// BundleResult contains the fetched WhatsApp Web version and downloaded bundle sources.
type BundleResult struct {
	Version string
	URLs    []string
	Sources []string
}

// FetchBundles discovers and downloads all WhatsApp Web JavaScript bundles.
func FetchBundles(ctx context.Context, client *http.Client) (*BundleResult, error) {
	if client == nil {
		client = &http.Client{
			Timeout: 60 * time.Second,
		}
	}

	result := &BundleResult{}

	// 1. Fetch main page HTML
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, BaseURL+"/", nil)
	if err != nil {
		return nil, err
	}
	setBrowserHeaders(req)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed fetching whatsapp web: %w", err)
	}
	defer resp.Body.Close()

	htmlBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading whatsapp web html: %w", err)
	}

	// 2. Discover script URLs
	urlSet := make(map[string]bool)

	for _, m := range scriptSrcRe.FindAllSubmatch(htmlBody, -1) {
		if len(m) > 1 {
			u := resolveURL(string(m[1]))
			if isAllowedBundle(u) {
				urlSet[u] = true
			}
		}
	}
	for _, m := range preloadRe.FindAllSubmatch(htmlBody, -1) {
		if len(m) > 1 {
			u := resolveURL(string(m[1]))
			if isAllowedBundle(u) {
				urlSet[u] = true
			}
		}
	}

	// 3. Check serviceworker.js / sw.js for client_revision and asset manifests
	for _, swPath := range []string{"/sw.js", "/serviceworker.js"} {
		swReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, BaseURL+swPath, nil)
		setBrowserHeaders(swReq)
		if swResp, err := client.Do(swReq); err == nil {
			swData, _ := io.ReadAll(swResp.Body)
			swResp.Body.Close()

			if m := clientRevisionRe.FindSubmatch(swData); len(m) > 1 {
				result.Version = fmt.Sprintf("2.3000.%s", string(m[1]))
			}

			// Check for asset manifest
			if mm := manifestRe.FindSubmatch(swData); len(mm) > 1 {
				manifestURL := fmt.Sprintf("%s/assets-manifest-%s.json", BaseURL, string(mm[1]))
				mReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, manifestURL, nil)
				setBrowserHeaders(mReq)
				if mResp, err := client.Do(mReq); err == nil {
					var manifest map[string]any
					if err := json.NewDecoder(mResp.Body).Decode(&manifest); err == nil {
						for filename := range manifest {
							if strings.HasSuffix(filename, ".js") {
								u := fmt.Sprintf("%s/%s", BaseURL, filename)
								urlSet[u] = true
							}
						}
					}
					mResp.Body.Close()
				}
			}
		}
	}

	var urls []string
	for u := range urlSet {
		urls = append(urls, u)
	}
	sort.Strings(urls)
	result.URLs = urls

	if len(urls) == 0 {
		return nil, fmt.Errorf("no JavaScript bundle URLs discovered from %s", BaseURL)
	}

	// 4. Download bundles concurrently
	sources := make([]string, len(urls))
	var wg sync.WaitGroup
	errCh := make(chan error, len(urls))
	sem := make(chan struct{}, MaxWorkers)

	for i, bundleURL := range urls {
		wg.Add(1)
		go func(idx int, targetURL string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			data, err := downloadWithRetry(ctx, client, targetURL, 3)
			if err != nil {
				errCh <- fmt.Errorf("failed downloading %s: %w", targetURL, err)
				return
			}
			sources[idx] = string(data)

			// Try to extract version if not found yet
			if result.Version == "" {
				if vm := versionStrRe.FindSubmatch(data); len(vm) > 1 {
					result.Version = string(vm[1])
				}
			}
		}(i, bundleURL)
	}

	wg.Wait()
	close(errCh)

	if len(errCh) > 0 {
		// Log errors but return available sources if any
		for e := range errCh {
			fmt.Printf("Warning: %v\n", e)
		}
	}

	result.Sources = sources
	return result, nil
}

func setBrowserHeaders(req *http.Request) {
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Sec-Fetch-Dest", "script")
	req.Header.Set("Sec-Fetch-Mode", "no-cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Referer", BaseURL+"/")
}

func resolveURL(raw string) string {
	raw = strings.TrimSpace(raw)
	if strings.HasPrefix(raw, "//") {
		return "https:" + raw
	}
	if strings.HasPrefix(raw, "/") {
		return BaseURL + raw
	}
	return raw
}

func isAllowedBundle(rawURL string) bool {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return false
	}
	host := strings.ToLower(parsed.Hostname())
	if host == "web.whatsapp.com" || host == "static.whatsapp.net" || strings.HasSuffix(host, ".whatsapp.net") {
		return strings.HasSuffix(parsed.Path, ".js")
	}
	return false
}

func downloadWithRetry(ctx context.Context, client *http.Client, rawURL string, retries int) ([]byte, error) {
	var lastErr error
	for attempt := 1; attempt <= retries; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
		if err != nil {
			return nil, err
		}
		setBrowserHeaders(req)

		resp, err := client.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {
			data, readErr := io.ReadAll(resp.Body)
			resp.Body.Close()
			if readErr == nil && len(data) > 0 {
				return data, nil
			}
			lastErr = readErr
		} else if err != nil {
			lastErr = err
		} else {
			lastErr = fmt.Errorf("HTTP status %d", resp.StatusCode)
			resp.Body.Close()
		}

		time.Sleep(time.Duration(attempt*300) * time.Millisecond)
	}
	return nil, lastErr
}
