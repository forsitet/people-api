package v1

import (
	"context"

	"people/internal/converter"
	peopleV1 "people/shared/pkg/openapi/people/v1"
)

func (h *peopleHandler) GetFriends(ctx context.Context, params peopleV1.GetFriendsParams) (peopleV1.GetFriendsRes, error) {
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
