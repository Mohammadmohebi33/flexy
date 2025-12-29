package authhandler

import "flexy/service/authservice"

type Handler struct {
	authSvc authservice.Service
}

func New(authSvc authservice.Service) Handler {
	return Handler{
		authSvc: authSvc,
	}
}
