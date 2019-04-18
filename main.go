package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/ua-parser/uap-go/uaparser"
)

const defaultAppStorePage = "https://itunes.apple.com/us/app/id544007664"
const defaultPlayStorePage = "https://play.google.com/store/apps/details?id=com.google.android.youtube"
const defaultUnknownPage = "https://youtube.com"

func main() {
	port := flag.String("port", "8080", "Port for the web server to listen on")
	flag.Parse()

	asp := getEnv("APP_STORE_PAGE", defaultAppStorePage)
	psp := getEnv("PLAY_STORE_PAGE", defaultPlayStorePage)
	ukp := getEnv("UNKNOWN_PAGE", defaultUnknownPage)

	uap := uaparser.NewFromSaved()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		client := uap.ParseOs(r.UserAgent())

		if client.Family == "iOS" {
			http.Redirect(w, r, asp, http.StatusFound)
		} else if client.Family == "Android" {
			http.Redirect(w, r, psp, http.StatusFound)
		} else {
			http.Redirect(w, r, ukp, http.StatusFound)
		}
	})

	http.ListenAndServe(":"+*port, nil)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
