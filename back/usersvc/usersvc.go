package usersvc

import (
	"github.com/champagneabuelo/openboard/back/pb"
)

var _ pb.UserServer = &UserSvc{}

type UserSvc struct{}
