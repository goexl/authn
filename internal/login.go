package internal

import (
	"github.com/goexl/gox/http"
)

type Login struct {
	// nolint:lll
	Url     string        `default:"credentials" json:"url,omitempty" xml:"url," yaml:"url" toml:"url" validate:"startsnotwith=/"`
	Methods []http.Method `default:"['get','post']" json:"methods,omitempty" xml:"methods" yaml:"methods" toml:"methods"`
}
