package request

import (
	"net/url"

	"github.com/n7down/microservices/internal/utils"
)

type ByIDRequest struct {
	ID string
}

func (r ByIDRequest) Validate() url.Values {
	errs := url.Values{}

	if r.ID == "" {
		errs.Add("id", "The id field is required!")
	}

	if !utils.IsValidUUID(r.ID) {
		errs.Add("id", "The id is not a valid format!")
	}

	return errs
}
