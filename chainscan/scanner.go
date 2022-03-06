package chainscan

import "sync"

type Scanner struct {
	mutex   sync.Mutex
	url     string
	baseUrl string
	apiKey  string
}

func NewScanner(url string, apiKey string) *Scanner {
	//es := &Provider{baseUrl: "https://api.bscscan.com/api", apiKey: "4KN3EGT11S1ZHXE9ZGFZXA9E4F5537FYIV"}
	es := &Scanner{baseUrl: url, apiKey: apiKey}

	return es
}
