package urlshortner

import (
	"net/http"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return nil
}

func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return nil, nil
}
