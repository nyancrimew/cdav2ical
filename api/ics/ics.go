package ics

import (
	"net/http"
	"path"

	"github.com/emersion/go-ical"
	"github.com/emersion/go-webdav"
	"github.com/emersion/go-webdav/caldav"
	"github.com/nyancrimew/cdav2ical/internal/config"
	"github.com/nyancrimew/cdav2ical/internal/utils"
	"github.com/phuslu/log"
	"github.com/valyala/fasthttp"
)

var c = &http.Client{}
var hc = webdav.HTTPClientWithBasicAuth(c, config.CalDavUser, config.CalDavPassword)

func GetICS(ctx *fasthttp.RequestCtx) {
	// Log Request and check for authentication
	utils.LogRequest(ctx)
	if !utils.IsAuthenticated(ctx) {
		ctx.SetStatusCode(http.StatusUnauthorized)
		log.Warn().Any("Client", ctx.RemoteIP()).Bytes("Path", ctx.Path()).Msg("Unauthenticated request rejected")
		return
	}

	href := ctx.UserValue("href").(string)

	cc, err := caldav.NewClient(hc, config.CalDAVHost)
	if err != nil {
		log.Error().Err(err).Any("Client", ctx.RemoteIP()).Str("href", href).Msg("Couldn't create CalDAV client")
		ctx.SetStatusCode(http.StatusInternalServerError)
		return
	}

	// Grab principal and home set for current user
	// TODO: we probably wanna simply make this a configuration option
	p, err := cc.FindCurrentUserPrincipal(ctx)
	if err != nil {
		log.Error().Err(err).Any("Client", ctx.RemoteIP()).Str("href", href).Msg("Couldn't find current user principal")
		ctx.SetStatusCode(http.StatusNotFound)
		return
	}

	chs, err := cc.FindCalendarHomeSet(ctx, p)
	if err != nil {
		log.Error().Err(err).Any("Client", ctx.RemoteIP()).Str("Principal", p).Str("href", href).Msg("Couldn't find calendar home set")
		ctx.SetStatusCode(http.StatusNotFound)
		return
	}

	calO, err := cc.GetCalendarObject(ctx, path.Join(chs, href))
	if err != nil {
		log.Error().Err(err).Any("Client", ctx.RemoteIP()).Str("Principal", p).Str("HomeSet", chs).Str("href", href).Msg("coulnt get calendar object")
		ctx.SetStatusCode(http.StatusNotFound)
		return
	}

	// Write ICS to context with appropriate content type
	ctx.SetContentType("text/calendar")
	ical.NewEncoder(ctx).Encode(calO.Data)
}
