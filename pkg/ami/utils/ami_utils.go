package utils

import (
	"encoding/json"
	"log"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
)

func VarsMap(values []string) map[string]string {
	r := make(map[string]string)

	for _, value := range values {
		k, v := VarsSplit(value)
		r[k] = v
	}

	return r
}

func VarsSplit(value string) (string, string) {
	s := strings.SplitN(value, "=", 2)
	k := s[0]

	if len(s) == 1 {
		return k, ""
	}

	return k, s[1]
}

func HasRawConnection(ip string, port int) (bool, error) {
	timeout := time.Second
	conn, err := net.DialTimeout(config.AmiNetworkTcpKey, net.JoinHostPort(ip, strconv.Itoa(port)), timeout)

	if err != nil {
		log.Printf("Connecting error: %v", err)
		return false, err
	}

	if conn != nil {
		defer conn.Close()
		log.Printf("Opened on: %s", net.JoinHostPort(ip, strconv.Itoa(port)))
		return true, nil
	}

	return false, nil
}

func HasRawConnectionWith(ip string, ports []int) (bool, error) {
	for _, port := range ports {

		if ok, err := HasRawConnection(ip, port); err != nil {
			return ok, err
		}
	}

	return true, nil
}

func IPDecode(ip string) (string, string, error) {
	u, err := url.Parse(ip)

	if err != nil {
		log.Printf("IP parse has error occurred = %v", err)
		return "", "", err
	}

	host, port, err := net.SplitHostPort(u.Host)
	return host, port, err
}

func ToJson(data interface{}) string {
	s, ok := data.(string)

	if ok {
		return s
	}

	result, err := json.Marshal(data)

	if err != nil {
		log.Printf(err.Error())
		return ""
	}

	return string(result)
}

func TakeKeyFromValue(collection map[string]string, value string) string {
	if len(collection) <= 0 {
		return value
	}

	for k, v := range collection {
		if strings.EqualFold(v, value) {
			return k
		}
	}

	return value
}

func Keys(in interface{}) (keys []string) {
	switch z := in.(type) {
	case map[string]int:
	case map[string]int32:
	case map[string]int64:
	case map[string]float32:
	case map[string]float64:
	case map[string]string:
	case map[string]bool:
		for k := range z {
			keys = append(keys, k)
		}
	case []int:
		for k := range z {
			keys = append(keys, strconv.Itoa(k))
		}
	default:
		return []string{}
	}
	return keys
}

func MergeMaps[K comparable, V any](m1 map[K]V, m2 map[K]V) map[K]V {
	merged := make(map[K]V)
	if len(m1) > 0 {
		for key, value := range m1 {
			merged[key] = value
		}
	}
	if len(m2) > 0 {
		for key, value := range m2 {
			merged[key] = value
		}
	}
	return merged
}
