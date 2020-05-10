package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v3"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		if url, ok := pathsToUrls[r.URL.String()]; ok {
			http.Redirect(w, r, url, http.StatusPermanentRedirect)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// PathRedirect represents a mapping between some path and url
type PathRedirect struct {
	Path string
	URL  string
}

type UnmarshalFunc func([]byte, interface{}) error

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return encodingdeHandler(yml, fallback, yaml.Unmarshal)
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
//
//     [
//  	 {
// 			"path": "/some-path",
// 			"url": "https://www.some-url.com/demo"
// 		 },
// 	   ]
//
// The only errors that can be returned all related to having
// invalid JSON data.
func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return encodingdeHandler(jsn, fallback, json.Unmarshal)
}

func encodingdeHandler(data []byte, fallback http.Handler, unmarshal UnmarshalFunc) (http.HandlerFunc, error) {
	pathsDescriptions, err := readPathDescription(data, unmarshal)

	if err != nil {
		return nil, err
	}

	paths := createPathMapping(pathsDescriptions)
	return MapHandler(paths, fallback), nil
}

func readPathDescription(data []byte, unmarshal UnmarshalFunc) ([]PathRedirect, error) {
	pathsDescriptions := []PathRedirect{}
	err := unmarshal(data, &pathsDescriptions)

	if err != nil {
		return nil, err
	}
	return pathsDescriptions, nil
}

func createPathMapping(pathsDescriptions []PathRedirect) map[string]string {
	paths := make(map[string]string)
	for _, pathDescribe := range pathsDescriptions {
		paths[pathDescribe.Path] = pathDescribe.URL
	}
	return paths
}
