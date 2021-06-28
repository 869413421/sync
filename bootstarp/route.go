package bootstarp

import (
	"github.com/gin-gonic/gin"
	"sync"
	"sync/routes"
)

var Router *gin.Engine
var once sync.Once

func SetupRoute() *gin.Engine {
	//1.singleton start routing
	once.Do(func() {
		//1.1 init routes
		Router = gin.New()
		routes.RegisterWebRoutes(Router)
	})

	//2.return router object
	return Router
}
