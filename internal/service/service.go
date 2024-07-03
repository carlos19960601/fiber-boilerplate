package service

import (
	"github.com/carlos19960601/fiber-boilerplate/internal/pkg/jwt"
	"github.com/carlos19960601/fiber-boilerplate/internal/pkg/sid"
	"github.com/carlos19960601/fiber-boilerplate/internal/repository"
)

type Service struct {
	sid *sid.Sid
	jwt *jwt.JWT
	tm  repository.Transaction
}

func NewService(
	tm repository.Transaction,
	sid *sid.Sid,
	jwt *jwt.JWT,
) *Service {
	return &Service{
		tm:  tm,
		sid: sid,
		jwt: jwt,
	}
}
