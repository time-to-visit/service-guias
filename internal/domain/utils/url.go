package utils

import "net/url"

func UrlValuesToMap(values url.Values) map[string]interface{} {
	m := make(map[string]interface{})
	for k, v := range values {
		m[k] = v[0]
	}
	return m
}
