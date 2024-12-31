package rest_api

import (
	"encoding/json"
	"net/http"

	rbac_app "github.com/bcdxn/garden-project/internal/app/rbac"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*Server)(nil)

type Server struct {
	rbacService *rbac_app.Service
}

func NewServer(rbacService *rbac_app.Service) Server {
	return Server{
		rbacService: rbacService,
	}
}

// (GET /ping)
func (Server) GetPing(w http.ResponseWriter, r *http.Request) {
	resp := Pong{
		Ping: "pong",
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s Server) GetApiV1Roles(w http.ResponseWriter, r *http.Request) {
	roles, err := s.rbacService.ListRoles()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"status":500, "error": "InternalServerError"}`))
		return
	}

	rolesRes, err := json.Marshal(roles)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"status":500, "error": "InternalServerError"}`))
		return
	}

	w.Write(rolesRes)
}
