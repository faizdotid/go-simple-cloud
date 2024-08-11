package utils

import (
	"math/rand"
	"net"
)

// Convert bytes to kilobytes
func BytesToKb(bytes int64) int64 {
	return bytes / 1024
}

// Create random string
func CreateRandomString() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// Getting local address
//
// reference: https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
func GetLocalAddr() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()

		if err != nil {
			return ""
		}

		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				if v.IP.IsGlobalUnicast() {
					return v.IP.String()
				}
			}

		}
	}
	return ""
}
