package home

import (
	"net/http"
	"todoproject/pkg/databasetools"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	databasetools.ReadSessions(databasetools.DB)
}
