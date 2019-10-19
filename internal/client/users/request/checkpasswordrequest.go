package request

import (
	"net/url"
)

type CheckPasswordRequest struct {
	Username string `json: "username" binding: "required"`
}

func (r *CheckPasswordRequest) Validate() url.Values {
	errs := url.Values{}

	if r.Username == "" {
		errs.Add("username", "The name field is required!")
	}

	return errs
}
