package auth

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

type HandlerAuth struct {
	s *ServiceAuth
}

func NewAuthHeader(service *ServiceAuth) *HandlerAuth {
	return &HandlerAuth{s: service}
}

func (h *HandlerAuth) SignUpHandler(ctx *fasthttp.RequestCtx) {
	var input CredsInput
	if err := json.Unmarshal(ctx.PostBody(), &input); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(fmt.Sprintf("{\"error\": \"Invalid input: %v\"}", err))
		return
	}

	if err := h.s.RegisterUser(&input); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(fmt.Sprintf("{\"error\": \"%v\"}", err))
		return
	}

	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.SetBodyString("{\"message\": \"User registered successfully\"}")
}
