package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/kyoto-framework/kyoto"
	"github.com/yuriizinets/hn/hn"
)

type CFeedStoryState struct {
	ID       int
	Index    int
	Title    string
	URL      string
	URLFmt   string
	User     string
	Time     string
	Points   int
	Comments int
}

func CFeedStory(id, index int) func(*kyoto.Context) CFeedStoryState {
	return func(ctx *kyoto.Context) (state CFeedStoryState) {

		// Define formatters
		fmturl := func(u string) string {
			_u, _ := url.Parse(u)
			return _u.Host
		}
		fmttime := func(t int) string {
			passed := time.Since(time.Unix(int64(t), 0))
			if passed > 24*time.Hour {
				return fmt.Sprintf("%vd ago", int(passed.Hours()/24))
			} else if passed > 1*time.Hour {
				return fmt.Sprintf("%vh ago", int(passed.Hours()))
			} else if passed > 1*time.Minute {
				return fmt.Sprintf("%vm ago", int(passed.Minutes()))
			} else {
				return "just now"
			}
		}

		// Fetch story
		story, err := hn.FetchStory(id)
		if err != nil {
			panic(err)
		}

		// Assign state
		state.ID = id
		state.Index = index
		state.Title = story.Title
		state.URL = story.URL
		state.URLFmt = fmturl(story.URL)
		state.User = story.User
		state.Time = fmttime(story.Time)
		state.Points = story.Points
		state.Comments = story.Comments

		// No error
		return
	}
}
