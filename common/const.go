package common

import "github.com/google/uuid"

const (
	DBTypeCategory = 1
	DBTypePost     = 2
	DBTypeUser     = 3
	DBTypeTag      = 4
)

const CurrentUser = "user"

type Requester interface {
	GetUserId() uuid.UUID
	GetEmail() string
	GetRole() string
}
