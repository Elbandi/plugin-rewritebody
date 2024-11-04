// Package plugin_rewritebody a plugin to rewrite response body.
package plugin_rewritebody

import (
	"context"
	"net/http"

	"github.com/traefik/plugin-rewritebody/handler"
)

// CreateConfig creates and initializes the plugin configuration.
func CreateConfig() *handler.Config {
	return &handler.Config{}
}

// New creates and returns a new rewrite body plugin instance.
func New(context context.Context, next http.Handler, config *handler.Config, name string) (http.Handler, error) {
	return handler.New(context, next, config, name)
}
