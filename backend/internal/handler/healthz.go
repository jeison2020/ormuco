package handler

import (
	"net/http"

	"ormuco.go/internal/util"
)

func HandleReadiness(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, 200, struct{}{})
}
