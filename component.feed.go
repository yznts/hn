package main

import (
	"github.com/kyoto-framework/kyoto"
	"github.com/yuriizinets/hn/hn"
)

type CFeedState struct {
	Category string
	Query    string

	Stories []*kyoto.ComponentF[CFeedStoryState]
}

// func CFeed(ctx *kyoto.Context) (state CFeedState) {
// 	// Default
// 	// Action state preload
// 	kyoto.ActionPreload(ctx, &state)
// }

func CFeed(category, query string) func(*kyoto.Context) CFeedState {
	return func(ctx *kyoto.Context) (state CFeedState) {
		// Set category and query
		state.Category = category
		state.Query = query
		// Defaults
		if state.Category == "" {
			state.Category = "top"
		}

		// Action state preload
		kyoto.ActionPreload(ctx, &state)

		// Determine load frame
		frame := [2]int{len(state.Stories), len(state.Stories) + 30}

		// Fetch stories
		storyids, err := hn.FetchStoryIds(
			state.Category,
			state.Query,
			frame,
		)
		if err != nil {
			panic(err)
		}

		// Init components
		for i, id := range storyids {
			state.Stories = append(state.Stories, kyoto.Use(ctx, CFeedStory(
				id, frame[0]+i+1,
			)))
		}
		// Return
		return
	}
}
