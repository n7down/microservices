package request

import (
	"net/url"

	"github.com/n7down/microservices/internal/utils"
)

type UpdateRequest struct {
	ID        string
	Username  string `json: "username" binding: "required"`
	Firstname string `json: "firstname" binding: "required"`
	Lastname  string `json: "lastname" binding: "required"`
}

func (r *UpdateRequest) Validate() url.Values {
	errs := url.Values{}

	if r.Username == "" {
		errs.Add("username", "The username field is required!")
	}

	if r.Firstname == "" {
		errs.Add("firstname", "The firstname field is required!")
	}

	if r.Lastname == "" {
		errs.Add("lastname", "The lastname field is required!")
	}

	if r.ID == "" {
		errs.Add("id", "The id field is required!")
	}

	if !utils.IsValidUUID(r.ID) {
		errs.Add("id", "The id is not a valid format!")
	}

	return errs
}
