package ping

import (
	"net"
	"time"
)

// PingHost пингует указанный хост и возвращает результат в виде строки
func PingHost(host string) string {
	if host == "" {
		host = "google.com"
	}

	start := time.Now()
	conn, err := net.DialTimeout("tcp", host+":80", 2*time.Second)
	if err != nil {
		return "❌ Ping failed: " + err.Error()
	}
	defer conn.Close()
	elapsed := time.Since(start)
	return "✅ Ping successful: " + elapsed.String()
}
