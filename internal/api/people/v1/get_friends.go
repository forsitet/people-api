package v1

import (
	"context"
	"database/sql"
	"errors"

	"people/internal/converter"
	peopleV1 "people/shared/pkg/openapi/people/v1"
)

func (h *peopleHandler) GetFriends(ctx context.Context, params peopleV1.GetFriendsParams) (peopleV1.GetFriendsRes, error) {
	if params.ID <= 0 {
		return &peopleV1.BadRequest{
			Code:    400,
			Message: "Invalid person ID",
		}, nil
	}

	_, err := h.Service.SearchByID(params.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &peopleV1.NotFound{
				Code:    404,
				Message: "Person not found",
			}, nil
		}
		return &peopleV1.InternalServer{
			Code:    500,
			Message: "Failed to check person existence",
		}, nil
	}

	friends, err := h.Service.GetFriends(params.ID)
	if err != nil {
		return &peopleV1.InternalServer{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	openAPIFriends := make(peopleV1.GetFriendsOKApplicationJSON, len(friends))
	for i, friend := range friends {
		openAPIFriends[i] = converter.PersonToApiModel(friend)
	}
	return &openAPIFriends, nil
}
