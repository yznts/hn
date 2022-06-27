package main

import (
	"strings"

	"github.com/kyoto-framework/kyoto"
	"github.com/kyoto-framework/zen"
)

type PStoryState struct {
	Navbar   *kyoto.ComponentF[CNavbarState]
	Story    *kyoto.ComponentF[CFeedStoryState]
	Comments *kyoto.ComponentF[CFeedCommentsState]
}

func PStory(ctx *kyoto.Context) (state PStoryState) {
	// Define rendering
	kyoto.Template(ctx, "page.story.html")
	// Determine id
	storyid := zen.Int(strings.ReplaceAll(ctx.Request.URL.Path, "/story/", ""))
	// Attach components
	state.Navbar = kyoto.Use(ctx, CNavbar)
	state.Story = kyoto.Use(ctx, CFeedStory(
		storyid, 0,
	))
	state.Comments = kyoto.Use(ctx, CFeedComments(CFeedCommentsState{
		StoryID: storyid,
	}))
	// Return
	return
}
