package healthz

import (
	"mweather-go/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleHealthz(c *gin.Context) {
	c.String(http.StatusOK, "Hello,It works. "+config.Items().APIVersion)
}
