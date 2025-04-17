package auth

import "github.com/fasthttp/router"

func RegisterAuthRoutes(r *router.Router, h *HandlerAuth) {
	r.POST("/api/signup", h.SignUpHandler)
}
