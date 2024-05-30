package authn

import (
	"github.com/goexl/authn/internal"
)

type Config struct {
	Prefix string          `json:"prefix,omitempty" yaml:"prefix" xml:"prefix" toml:"prefix"`
	Login  internal.Login  `json:"login,omitempty" yaml:"login" xml:"login" toml:"login"`
	Logout internal.Logout `json:"logout,omitempty" yaml:"logout" xml:"logout" toml:"logout"`
}
