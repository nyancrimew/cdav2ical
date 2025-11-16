package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var CalDAVHost = os.Getenv("CALDAV_HOST")
var CalDavUser = os.Getenv("CALDAV_USER")
var CalDavPassword = os.Getenv("CALDAV_PASSWORD")
var APIToken = []byte(os.Getenv("API_TOKEN"))
