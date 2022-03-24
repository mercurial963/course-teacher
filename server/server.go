package server

import (
	"course-teacher/config"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := NewRouter()
	gin.SetMode(gin.ReleaseMode)
	r.Run(config.GetListenAddr())
}
