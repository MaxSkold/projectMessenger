package auth

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type HandlerAuth struct {
	s *ServiceAuth
}

func NewAuthHandler(service *ServiceAuth) *HandlerAuth {
	return &HandlerAuth{s: service}
}

func (h *HandlerAuth) SignUpHandler(ctx *fasthttp.RequestCtx) {
	var input CredsInput
	if err := json.Unmarshal(ctx.PostBody(), &input); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBody([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	if err := h.s.RegisterUser(&input); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBody([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.SetBody([]byte(`{"message": "User registered successfully"}`))
}

func (h *HandlerAuth) LoginHandler(ctx *fasthttp.RequestCtx) {
	var input CredsInput
	if err := json.Unmarshal(ctx.PostBody(), &input); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBody([]byte(`{"error": "` + err.Error() + `"}`))
	}

}
