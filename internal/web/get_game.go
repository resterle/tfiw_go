package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGame(r *gin.Context) {
	sid := r.Params.ByName("sid")

	if v, ok := GetGameView(sid); ok {
		r.JSON(http.StatusOK, v)
	} else {
		r.Status(http.StatusNotFound)
	}
}
