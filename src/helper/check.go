package helper

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type TokenResult struct {
	Token  string `json:"token"`
	Status string `json:"status"`
}

func IsValidLength(token string) bool {
	l := len(token)
	return l == 59 || l == 70
}

func CheckToken(tokens []string, concurrency int) []TokenResult {
	results := make([]TokenResult, len(tokens))
	client := &http.Client{Timeout: 10 * time.Second}

	sem := make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	for i, token := range tokens {
		wg.Add(1)
		sem <- struct{}{}
		go func(i int, token string) {
			defer wg.Done()
			defer func() { <-sem }()

			if !IsValidLength(token) {
				results[i] = TokenResult{Token: token, Status: "Bad token length"}
				return
			}

			req, _ := http.NewRequest("GET", "https://discord.com/api/v9/users/@me", nil)
			req.Header.Set("Authorization", token)

			resp, err := client.Do(req)
			if err != nil {
				results[i] = TokenResult{Token: token, Status: "error"}
				return
			}
			defer resp.Body.Close()

			switch resp.StatusCode {
			case 200:
				results[i] = TokenResult{Token: token, Status: "valid"}
			case 401:
				results[i] = TokenResult{Token: token, Status: "invalid"}
			case 429:
				results[i] = TokenResult{Token: token, Status: "rate limited"}
			default:
				results[i] = TokenResult{Token: token, Status: fmt.Sprintf("status %d", resp.StatusCode)}
			}
		}(i, token)
	}

	wg.Wait()
	return results
}

