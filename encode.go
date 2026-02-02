package sp

import (
	"net/url"

	"github.com/sonh/qs"
)

func Encode(v any) (url.Values, error) {
	return qs.NewEncoder().Values(v)
}
