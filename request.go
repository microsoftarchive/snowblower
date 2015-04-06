package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

var privateIPMasks = func() []net.IPNet {
	masks := []net.IPNet{}
	for _, m := range []string{
		"192.168.0.0/16",
		"10.0.0.0/8",
		"172.16.0.0/12",
		"fc00::/7",
	} {
		_, n, _ := net.ParseCIDR(m)
		masks = append(masks, *n)
	}
	return masks
}()

func isPublicIP(ip net.IP) bool {
	for _, m := range privateIPMasks {
		if m.Contains(ip) {
			return false
		}
	}
	return true
}

func realRemoteAddr(request *http.Request) string {
	if header := request.Header.Get("X-Forwarded-For"); header != "" {
		for _, addy := range strings.Split(header, ",") {
			addy = strings.TrimSpace(addy)
			ip := net.ParseIP(addy)
			if ip != nil && isPublicIP(ip) {
				return addy
			}
		}
	}

	return request.RemoteAddr
}

func requestHeadersAsArray(request *http.Request) []string {
	var headers []string
	for k, v := range request.Header {
		for _, i := range v {
			headers = append(headers, fmt.Sprintf("%s: %s", k, i))
		}
	}
	return headers
}
