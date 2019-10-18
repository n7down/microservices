package request

import (
	"net/url"
)

type ListRequest struct{}

func (r *ListRequest) Validate() url.Values {
	errs := url.Values{}
	//if r.Name == "" {
	//errs.Add("name", "The name field is required!")
	//}
	return errs
}
