package main

import (
	"net/http"
	"os"

	"github.com/2enhance/placeholdr/config"
	"github.com/2enhance/placeholdr/svg"
	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/text"
	"github.com/gorilla/pat"
)

// use JSON logging when run by Up (including `up start`)
func init() {
	if os.Getenv("UP_STAGE") == "" {
		log.SetHandler(text.Default)
	} else {
		log.SetHandler(json.Default)
	}
}

func main() {
	addr := ":" + os.Getenv("PORT")
	app := pat.New()

	app.Get("/avatar/{dimensions}", AvatarHandler)
	app.Get("/avatar", AvatarHandler)
	app.Get("/logo/{dimensions}", LogoHandler)
	app.Get("/logo", LogoHandler)
	app.Get("/{dimensions}", PlaceholderHandler)
	app.Get("/", PlaceholderHandler)

	if err := http.ListenAndServe(addr, app); err != nil {
		log.WithError(err).Fatal("error listening")
	}
}

// LogoHandler responds to GET @ /logo/{?dimensions}
func LogoHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.WriteHeader(http.StatusOK)

	query := req.URL.Query()
	options := config.NewLogoOptions(
		query.Get(":dimensions"),
		query.Get("id"),
	)

	svg.NewLogo(w, options)
}

// AvatarHandler responds to GET @ /avatar/{?dimensions}
func AvatarHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.WriteHeader(http.StatusOK)

	query := req.URL.Query()
	options := config.NewAvatarOptions(
		query.Get(":dimensions"),
		query.Get("id"),
	)

	svg.NewAvatar(w, options)
}

// PlaceholderHandler responds to GET @ /{dimensions}
func PlaceholderHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.WriteHeader(http.StatusOK)

	query := req.URL.Query()

	isGrey := false
	if _, ok := query["grey"]; ok {
		isGrey = true
	}

	options := config.NewPlaceholderOptions(
		query.Get(":dimensions"),
		query.Get("id"),
		isGrey,
	)

	svg.NewPlaceholder(w, options)

}
