package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(s string) (string, error) {
	u, err := url.ParseRequestURI(s)
	if err != nil {
		return s, fmt.Errorf("Not an url.")
	}

	return u.Host + strings.TrimSuffix(u.Path, "/"), nil
}
