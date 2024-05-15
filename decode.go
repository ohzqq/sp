package sp

import (
	"net/url"

	"github.com/sonh/qs"
)

func Marshal(v any) (url.Values, error) {
	return qs.NewEncoder().Values(v)
}
