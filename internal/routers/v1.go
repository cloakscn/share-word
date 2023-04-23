package routers

import (
	"github.com/cloakscn/share-word/internal/services/ajax"
	"github.com/cloakscn/share-word/internal/vo/req"
	"github.com/gin-gonic/gin"
	"net/http"
)

func routerV1(group *gin.RouterGroup) {
	group.GET("/:id", GetByID)
	group.POST("/create", Create)
}

func GetByID(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		ajax.Error(context, ajax.ErrorNotFound)
		return
	}
	ajax.Success(context, id)
}

func Create(context *gin.Context) {
	var data req.CreateWorld
	if err := context.ShouldBindJSON(&data); err != nil {
		ajax.Error(context)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
