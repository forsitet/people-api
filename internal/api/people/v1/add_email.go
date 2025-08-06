package v1

import (
	"context"

	peopleV1 "people/shared/pkg/openapi/people/v1"
)

func (h *peopleHandler) AddEmail(ctx context.Context, req *peopleV1.Email, params peopleV1.AddEmailParams) (peopleV1.AddEmailRes, error) {
	if err := h.Service.AddEmail(params.ID, req.Email); err != nil {
		return &peopleV1.InternalServer{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	return &peopleV1.AddEmailOK{}, nil
}
