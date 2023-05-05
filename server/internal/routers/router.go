package routers

import (
	"github.com/cloakscn/share-word/server/internal/routers/v1"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	v1.RouterV1(engine.Group("/v1"))
}
