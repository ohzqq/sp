package sp

import (
	"net/url"
)

func Decode(v url.Values, t any) error {
	b := DefaultBinder{}
	return b.Bind(v, t)
}
