package sp

import "net/url"

func Unmarshal(v url.Values, t any) error {
	b := DefaultBinder{}
	return b.Bind(t, v)
}
