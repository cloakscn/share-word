package https

import (
	"github.com/cloakscn/share-word/internal/routers"
	"github.com/gin-gonic/gin"
	"log"
	"sync"
)

type Config struct {
	Port string
}

func Server(cfg *Config, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Println("初始化 http 服务。")
	log.Println("http://localhost:" + cfg.Port)

	//gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	routers.Router(engine)

	err := engine.Run(":" + cfg.Port)
	if err != nil {
		log.Fatalln("初始化 http 服务失败。")
	}
	log.Println("初始化 http 服务成功。")
}
