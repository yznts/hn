package main

import (
	"github.com/kyoto-framework/kyoto"
)

type PFeedState struct {
	Navbar *kyoto.ComponentF[CNavbarState]
	Feed   *kyoto.ComponentF[CFeedState]
}

func PFeed(ctx *kyoto.Context) (state PFeedState) {
	// Define rendering
	kyoto.Template(ctx, "page.feed.html")
	// Attach components
	state.Navbar = kyoto.Use(ctx, CNavbar)
	state.Feed = kyoto.Use(ctx, CFeed(
		ctx.Request.URL.Query().Get("category"),
		ctx.Request.URL.Query().Get("query"),
	))
	// Return
	return
}
