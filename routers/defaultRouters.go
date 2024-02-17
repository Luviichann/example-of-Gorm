package routers

import (
	"gormnote/controllers/defaults"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", defaults.DefaultController{}.Index)
		defaultRouters.GET("/add", defaults.DefaultController{}.Add)
		defaultRouters.GET("/delete", defaults.DefaultController{}.Delete)
		defaultRouters.GET("/query", defaults.DefaultController{}.Query)
		defaultRouters.GET("/update", defaults.DefaultController{}.Update)

		defaultRouters.GET("/article/bt", defaults.DefaultController{}.Query_bt)
		defaultRouters.GET("/article/hm", defaults.DefaultController{}.Query_hm)
		defaultRouters.GET("/article/mtm", defaults.DefaultController{}.Query_mtm)
	}
}
