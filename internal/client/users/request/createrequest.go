package request

import (
	"net/url"
)

type CreateRequest struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
}

func (r *CreateRequest) Validate() url.Values {
	errs := url.Values{}

	if r.Username == "" {
		errs.Add("username", "The username field is required!")
	}

	if r.Password == "" {
		errs.Add("password", "The password field is required!")
	}

	if r.Firstname == "" {
		errs.Add("firstname", "The firstname field is required!")
	}

	if r.Lastname == "" {
		errs.Add("lastname", "The lastname field is required!")
	}

	return errs
}
