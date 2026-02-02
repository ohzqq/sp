package sp

import (
	"net/http"
	"net/url"
)

func Decode(v url.Values, t any) error {
	b := DefaultBinder{}
	return b.Bind(v, t)
}

func DecodeRequest(v *http.Request, t any) error {
	b := DefaultBinder{}
	return b.BindRequest(v, t)
}
