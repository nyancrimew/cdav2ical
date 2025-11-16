package utils

import (
	"bytes"

	"github.com/nyancrimew/cdav2ical/internal/config"
	"github.com/valyala/fasthttp"
)

func IsAuthenticated(ctx *fasthttp.RequestCtx) bool {
	t := ctx.QueryArgs().Peek("token")
	return bytes.Equal(config.APIToken, t)
}
