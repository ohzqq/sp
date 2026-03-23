package sp

// Decode takes url.Values and decodes it to a struct.
// see https://pkg.go.dev/github.com/labstack/echo/v5 for documentation.
func Decode[Q ~map[string][]string](v Q, t any) error {
	return BindQueryParams(v, t)
}
