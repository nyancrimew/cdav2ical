package main

import (
	"os"

	"github.com/fasthttp/router"
	"github.com/nyancrimew/cdav2ical/api/ics"
	"github.com/nyancrimew/cdav2ical/internal/utils"
	"github.com/phuslu/log"
	"github.com/valyala/fasthttp"
)

var up = []byte{'O', 'K'}

func Index(ctx *fasthttp.RequestCtx) {
	utils.LogRequest(ctx)
	_, _ = ctx.Write(up)
}

func main() {
	// Setup logging
	if log.IsTerminal(os.Stderr.Fd()) {
		log.DefaultLogger = log.Logger{
			TimeFormat: "15:04:05",
			Caller:     1,
			Writer: &log.ConsoleWriter{
				ColorOutput:    true,
				QuoteString:    true,
				EndWithMessage: true,
			},
		}
	}

	log.Info().Msg("Starting server...")

	// Setup router and panic handling
	mux := router.New()
	mux.PanicHandler = func(ctx *fasthttp.RequestCtx, i interface{}) {
		ctx.SetStatusCode(500)
		log.Error().Any("p", i).Msg("PANIC")
	}

	// Routes
	mux.GET("/", Index)
	mux.GET("/ics/{href}.ics", ics.GetICS)

	// Start server
	log.Fatal().Err(fasthttp.ListenAndServe(":5505", mux.Handler)).Msg("fuck.")
}
