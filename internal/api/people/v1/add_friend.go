package v1

import (
	"context"
	"database/sql"
	"errors"

	peopleV1 "people/shared/pkg/openapi/people/v1"
)

func (h *peopleHandler) AddFriend(ctx context.Context, req *peopleV1.AddFriendReq, params peopleV1.AddFriendParams) (peopleV1.AddFriendRes, error) {
	if params.ID <= 0 {
		return &peopleV1.BadRequest{
			Code:    400,
			Message: "Invalid person ID",
		}, nil
	}

	friendID, ok := req.FriendID.Get()
	if !ok {
		return &peopleV1.BadRequest{
			Code:    400,
			Message: "friend_id is required",
		}, nil
	}

	if friendID <= 0 {
		return &peopleV1.BadRequest{
			Code:    400,
			Message: "Invalid friend ID",
		}, nil
	}

	// Проверяем существование основного пользователя
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

	// Проверяем существование друга
	_, err = h.Service.SearchByID(friendID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &peopleV1.NotFound{
				Code:    404,
				Message: "Friend not found",
			}, nil
		}
		return &peopleV1.InternalServer{
			Code:    500,
			Message: "Failed to check friend existence",
		}, nil
	}

	// Проверяем, что пользователь не пытается добавить себя в друзья
	if params.ID == friendID {
		return &peopleV1.BadRequest{
			Code:    400,
			Message: "Cannot add yourself as a friend",
		}, nil
	}

	if err := h.Service.AddFriend(params.ID, friendID); err != nil {
		return &peopleV1.InternalServer{
			Code:    500,
			Message: "Failed to add friend",
		}, nil
	}

	return &peopleV1.AddFriendOK{}, nil
}
