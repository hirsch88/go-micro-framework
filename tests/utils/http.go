package testUtils

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func Get(app *gin.Engine, path string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", path, nil)
	app.ServeHTTP(w, req)
	return w
}
