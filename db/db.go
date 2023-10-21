package db

import (
	"nativeApi/models"
)

// set up a database dummy
var (
	Moviedb = make(map[string]models.Movie)
)
