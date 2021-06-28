package bootstarp

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var Router *gin.Engine
var once sync.Once

func SetupRoute() *gin.Engine {
	//1.singleton start routing
	once.Do(func() {
		Router = gin.New()
	})

	//2.return router object
	return Router
}
