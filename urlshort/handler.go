package urlshort

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

//takes in a map of short-urls to urls i.e path:dest
//returns a handler function which redirects to urls specified
func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

//constructs
func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	//parse Yaml
	var urlPaths []UrlPath
	err := yaml.Unmarshal(yamlBytes, &urlPaths)
	if err != nil {
		return nil, err
	}
	//construct map
	pathsToUrls := make(map[string]string)

	for _, u := range urlPaths {
		pathsToUrls[u.Path] = u.Url
	}

	//pass into MapHandler
	return MapHandler(pathsToUrls, fallback), nil
}

type UrlPath struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}
