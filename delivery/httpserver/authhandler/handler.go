package authhandler

import "flexy/service/userservice"

type Handler struct {
	authSvc userservice.Service
}

func New(authSvc userservice.Service) Handler {
	return Handler{
		authSvc: authSvc,
	}
}
