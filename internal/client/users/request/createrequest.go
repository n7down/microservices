package request

import (
	"net/url"
)

type CreateRequest struct {
	Username  string `json: "username" binding: "required"`
	Password  string `json: "password" binding: "required"`
	Firstname string `json: "firstname" binding: "required"`
	Lastname  string `json: "lastname" binding: "required"`
}

func (r *CreateRequest) Validate() url.Values {
	errs := url.Values{}

	if r.Username == "" {
		errs.Add("username", "The name field is required!")
	}

	if r.Password == "" {
		errs.Add("password", "The name field is required!")
	}

	if r.Username == "" {
		errs.Add("firstname", "The name field is required!")
	}

	if r.Username == "" {
		errs.Add("lastname", "The name field is required!")
	}

	return errs
}
