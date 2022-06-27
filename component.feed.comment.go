package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/kyoto-framework/kyoto"
)

type CFeedCommentState struct {
	ID        int
	Depth     int
	Text      template.HTML
	User      string
	Time      string
	Points    int
	Deleted   bool
	NestedIDs []int
	Nested    *kyoto.ComponentF[CFeedCommentsState]
}

func CFeedComment(id, depth int) kyoto.Component[CFeedCommentState] {
	return func(ctx *kyoto.Context) (state CFeedCommentState) {
		// Preload action state
		kyoto.ActionPreload(ctx, &state)
		// Handle action
		_handled := kyoto.Action(ctx, "ToggleNested", func(args ...any) {
			if state.Nested == nil {
				state.Nested = kyoto.Use(ctx, CFeedComments(CFeedCommentsState{
					CommentIDs: state.NestedIDs,
					Depth:      state.Depth + 1,
				}))
			} else {
				state.Nested = nil
			}
		})
		// Prevent further execution if action handled
		if _handled {
			return
		}
		// Create entrypoint
		entrypoint := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%v.json?print=pretty", id)
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
			Text    string `json:"text"`
			User    string `json:"by"`
			Time    int    `json:"time"`
			Points  int    `json:"score"`
			Deleted bool   `json:"deleted"`
			Kids    []int  `json:"kids"`
		}
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			panic(err)
		}
		// Assign to fields
		state.ID = id
		state.Depth = depth
		state.Text = template.HTML(data.Text)
		state.User = data.User
		state.Time = "4h ago"
		state.Points = data.Points
		state.Deleted = data.Deleted
		state.NestedIDs = data.Kids
		return
	}
}
