package request

import (
	"net/url"
)

type CheckPasswordRequest struct{}

func (r *CheckPasswordRequest) Validate() url.Values {
	errs := url.Values{}
	//if r.Name == "" {
	//errs.Add("name", "The name field is required!")
	//}
	return errs
}
