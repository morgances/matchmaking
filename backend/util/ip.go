package util

import (
	"strings"
	"net/http"
	"net"
)

func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if xffv := req.Header.Get(http.CanonicalHeaderKey("X-Forwarded-For")); xffv != "" {
		ips := strings.SplitN(xffv, ",", 3)
		remoteAddr = ips[0]
	} else if ip := req.Header.Get(http.CanonicalHeaderKey("XRealIP")); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}
