package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/kyoto-framework/kyoto"
)

type CFeedCommentsState struct {
	StoryID    int
	CommentIDs []int

	Depth    int
	Comments []*kyoto.ComponentF[CFeedCommentState]
}

func CFeedComments(argstate CFeedCommentsState) kyoto.Component[CFeedCommentsState] {
	return func(ctx *kyoto.Context) (state CFeedCommentsState) {
		// Replace state with arguments' one
		state = argstate
		// Determine comment ids
		ids := []int{}
		// Fetch story comment ids, if story id provided
		if state.StoryID != 0 {
			// Create entrypoint
			entrypoint := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%v.json?print=pretty", state.StoryID)
			// Make request
			resp, err := http.Get(entrypoint)
			// Handle errors
			if err != nil {
				panic(err)
			}
			if resp.StatusCode != 200 {
				panic(errors.New("status code is not 200"))
			}
			// Close body after processing
			defer resp.Body.Close()
			// Unpack response
			var data struct {
				Title    string `json:"title"`
				URL      string `json:"url"`
				User     string `json:"by"`
				Time     int    `json:"time"`
				Points   int    `json:"score"`
				Comments int    `json:"descendants"`
				Kids     []int  `json:"kids"`
			}
			json.NewDecoder(resp.Body).Decode(&data)
			// Set
			ids = data.Kids
		} else { // Set ids from arguments
			ids = state.CommentIDs
		}
		// Init components
		for _, id := range ids {
			state.Comments = append(state.Comments, kyoto.Use(ctx, CFeedComment(id, state.Depth)))
		}
		// Return
		return
	}
}
