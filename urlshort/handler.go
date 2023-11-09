package urlshort

import (
	"log/slog"
	"net/http"

	"gopkg.in/yaml.v3"
)

type UrlMapperEntry struct {
	path string `yaml:"path"`
	url  string `yaml:"url"`
}
type UrlMapper []UrlMapperEntry

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler, logger *slog.Logger) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Request url: " + r.URL.String())

		shortenedUrl, exists := pathsToUrls[r.URL.String()]
		if !exists {
			logger.Warn("No url in map")
			fallback.ServeHTTP(w, r)
		} else {
			logger.Info("Redirect...")
			http.Redirect(w, r, shortenedUrl, http.StatusMovedPermanently)
		}
	})
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler, logger *slog.Logger) (http.HandlerFunc, error) {
	emptyMap := map[string]string{}
	var urlMapper []struct {
		Path string `yaml:"path"`
		URL  string `yaml:"url"`
	}

	err := yaml.Unmarshal(yml, &urlMapper)
	if err != nil {
		logger.Error("error: " + err.Error())
	}

	for _, mapper := range urlMapper {
		emptyMap[mapper.Path] = mapper.URL
	}

	return MapHandler(emptyMap, fallback, logger), nil
}
