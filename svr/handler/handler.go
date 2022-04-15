package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if redirectUrl, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, redirectUrl, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YamlHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := ParseYaml(yaml)
	if err != nil {
		fmt.Printf("Failed to parse YAML, err: %v", err)
		return nil, err
	}

	pathsToUrls := BuildMap(pathUrls)

	return MapHandler(pathsToUrls, fallback), nil
}

func JsonHandler(json []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := ParseJson(json)
	if err != nil {
		fmt.Printf("Failed to parse JSON, err: %v", err)
		return nil, err
	}

	pathsToUrls := BuildMap(pathUrls)

	return MapHandler(pathsToUrls, fallback), nil
}

func BuildMap(pathUrls []PathUrl) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pathToUrl := range pathUrls {
		pathsToUrls[pathToUrl.Path] = pathToUrl.URL
	}
	return pathsToUrls
}

func ParseYaml(data []byte) ([]PathUrl, error) {
	var pathUrls []PathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

type PathUrl struct {
	Path string
	URL  string
}

func ParseJson(data []byte) ([]PathUrl, error) {
	var pathsToUrls []PathUrl
	err := json.Unmarshal(data, &pathsToUrls)
	if err != nil {
		return nil, err
	}
	return pathsToUrls, nil
}
