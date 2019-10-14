package request

import (
	"net/url"
)

type HelloRequest struct {
	Name string `json: "name" binding:"required"`
}

func (r *HelloRequest) Validate() url.Values {
	errs := url.Values{}
	if r.Name == "" {
		errs.Add("name", "The name field is required!")
	}
	return errs
}
