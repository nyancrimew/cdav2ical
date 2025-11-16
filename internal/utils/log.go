package utils

import (
	"github.com/phuslu/log"
	"github.com/valyala/fasthttp"
)

func LogRequest(ctx *fasthttp.RequestCtx) {
	log.Info().Any("Client", ctx.RemoteIP()).Bytes("Path", ctx.Path()).Bool("Authenticated", IsAuthenticated(ctx)).Msg("Requested")
}
