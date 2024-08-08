package home

import (
	"net/http"
	"todoproject/pkg/databasetools"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	sessionId, _ := r.Cookie("session_id")
	userId := cookie
	databasetools.ReadSessions(databasetools.DB)
}
