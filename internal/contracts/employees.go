package contracts

import "protei/internal/entity"

type Employees interface {
	HandlerFindUserId(u entity.UserMailIn) (*entity.FindUserIdResponseOut, error)
	HandlerFindReasonId(u *entity.UserPersonId) (*entity.FindUserReasonIdResponse, error)
}
