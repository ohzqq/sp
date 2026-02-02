package sp

import (
	"net/url"

	"github.com/sonh/qs"
)

// Encode encodes a struct to url.Values.
// see https://pkg.go.dev/github.com/sonh/qs for documentation.
func Encode(v any) (url.Values, error) {
	return qs.NewEncoder().Values(v)
}
