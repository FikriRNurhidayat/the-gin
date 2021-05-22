package config

import "os"

var SIGNING_SECRET []byte = []byte(os.Getenv("SIGNING_SECRET"))
