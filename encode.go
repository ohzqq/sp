package sp

import (
	"github.com/ohzqq/qs"
)

// Encode encodes a struct to url.Values.
// see https://pkg.go.dev/github.com/sonh/qs for documentation.
func Encode[Q map[string][]string](v any) (Q, error) {
	return qs.NewEncoder().Values(v)
}
