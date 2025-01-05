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

func (s Server) GetApiV1Roles(w http.ResponseWriter, r *http.Request) {
	roles, err := s.rbacService.ListRoles(r.Context())
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

func (s Server) GetApiV1RolesRoleIdPermissions(w http.ResponseWriter, r *http.Request, roleId string) {
	permissions, err := s.rbacService.ListPermissionsByRoleID(r.Context(), roleId)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"status":500, "error": "InternalServerError"}`))
		return
	}

	permissionsRes, err := json.Marshal(permissions)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"status":500, "error": "InternalServerError"}`))
		return
	}

	w.Write(permissionsRes)
}
