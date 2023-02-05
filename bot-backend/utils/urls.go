package utils

import (
	"net"
	"net/url"
	"strings"
)

func GetLambdaInvokeUrlDomain() (string, error) {
	if u, err := url.Parse(GetLambdaInvokeUrl()); err != nil {
		return "", err
	} else {
		if !strings.Contains(u.Host, ":") {
			return u.Host, nil
		}

		if host, _, err := net.SplitHostPort(u.Host); err != nil {
			return "", err
		} else {
			return host, nil
		}
	}
}
