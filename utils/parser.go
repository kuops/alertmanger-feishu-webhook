package utils

import (
	"fmt"
	"log"
	"net/url"
)

func PromGraphURL(promUrl, promDomain string) string {
	u, err := url.Parse(promUrl)
	if err != nil {
		log.Printf("err parser prom graph url: %v\n", promUrl)
	}
	if promDomain == "" {
		promDomain = "http://" + u.Host
	}
	promGraphUrl := fmt.Sprintf("%v%v?%v", promDomain, u.Path, u.RawQuery)
	return promGraphUrl
}

