package request

import (
	"net/url"
)

type LoginRequest struct{}

func (r *LoginRequest) Validate() url.Values {
	errs := url.Values{}
	//if r.Name == "" {
	//errs.Add("name", "The name field is required!")
	//}
	return errs
}
