package user

import (
	"github.com/shripadmhetre/golang-learnings/wiredemo/domain"
	"net/http"
)

type handler struct {
	svc domain.UserService
}

func (h *handler) FetchByUsername() http.HandlerFunc {
	panic("implement me")
}
