package sp

import (
	"net/http"
	"net/url"
)

// Decode takes url.Values and decodes it to a struct.
// see https://pkg.go.dev/github.com/labstack/echo/v5 for documentation.
func Decode(v url.Values, t any) error {
	b := DefaultBinder{}
	return b.Bind(v, t)
}

// DecodeRequest takes url.Values and decodes it to a struct.
// see https://pkg.go.dev/github.com/labstack/echo/v5 for documentation.
func DecodeRequest(v *http.Request, t any) error {
	b := DefaultBinder{}
	return b.BindRequest(v, t)
}
