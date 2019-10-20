package request

import (
	"net/url"

	"github.com/n7down/microservices/internal/utils"
)

type DeleteRequest struct {
	ID string
}

func (r *DeleteRequest) Validate() url.Values {
	errs := url.Values{}

	if r.ID == "" {
		errs.Add("id", "The id field is required!")
	}

	if !utils.IsValidUUID(r.ID) {
		errs.Add("id", "The id is not a valid format!")
	}

	return errs
}
