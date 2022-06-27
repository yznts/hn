package main

import (
	"net/http"
	"os"

	"github.com/kyoto-framework/kyoto"
	"github.com/kyoto-framework/zen"
)

func main() {

	// Handle statics
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/dist"))))

	// Handle pages
	kyoto.HandlePage("/", PFeed)
	kyoto.HandlePage("/story/", PStory)

	// Handle actions
	kyoto.HandleAction(CFeed("", ""))
	kyoto.HandleAction(CFeedComment(0, 0))

	kyoto.Serve(":" + zen.Or(os.Getenv("PORT"), "25025"))
}
