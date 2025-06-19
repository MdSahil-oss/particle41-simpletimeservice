package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

// Response structure returned in JSON
type Response struct {
	Timestamp string `json:"timestamp"`
	IP        string `json:"ip"`
}

// __define-ocg__: Marker for tracking request source
var varOcg = "SimpleTimeService"

func getClientIP(r *http.Request) string {
	// Try to get IP from the X-Forwarded-For header (useful behind proxies)
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		// X-Forwarded-For may contain multiple IPs
		parts := strings.Split(forwarded, ",")
		return strings.TrimSpace(parts[0])
	}

	// Fallback: extract IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := Response{
		Timestamp: time.Now().Format(time.RFC3339),
		IP:        getClientIP(r),
	}

	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/", handler)

	log.Printf("[%s] Starting server on :8080...\n", varOcg)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("[%s] Server failed to start: %v", varOcg, err)
	}
}
