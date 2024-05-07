package handler

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"ormuco.go/internal/util"
	"strconv"
	"strings"
)

func (server *HTTPServer) GetVersion(w http.ResponseWriter, r *http.Request) {
	v1 := chi.URLParam(r, "v1")
	v2 := chi.URLParam(r, "v2")
	result := CompareVersions(v1, v2)
	if result == "The version 1 is not a number" || result == "The version 2 is not a number" {
		util.ResponseWithError(w, 400, result)
		return
	}
	util.RespondWithJSON(w, 201, result)
	return
}

func CompareVersions(v1, v2 string) string {
	parts1 := strings.Split(v1, ".")
	parts2 := strings.Split(v2, ".")

	maxLen := len(parts1)
	if len(parts2) > maxLen {
		maxLen = len(parts2)
	}

	for i := 0; i < maxLen; i++ {
		var num1, num2 int
		var err error
		if i < len(parts1) {
			num1, err = strconv.Atoi(parts1[i])
			if err != nil {
				return fmt.Sprintf("The version 1 is not a number")
			}
		}
		if i < len(parts2) {
			num2, err = strconv.Atoi(parts2[i])
			if err != nil {
				return fmt.Sprintf("The version 2 is not a number")
			}
		}

		if num1 < num2 {
			return fmt.Sprintf("The version %s is lower than version. %s", v1, v2)
		} else if num1 > num2 {
			return fmt.Sprintf("The version %s is greater than version. %s", v1, v2)
		}
	}
	return "they are equal"
}
