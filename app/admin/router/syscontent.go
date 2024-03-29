package router

import (
	"github.com/gin-gonic/gin"
	"fiy/app/admin/apis/syscontent"
	"fiy/app/admin/middleware"
	jwt "fiy/pkg/jwtauth"
)


func init()  {
	routerCheckRole = append(routerCheckRole, registerSysContentRouter)
}

// 需认证的路由代码
func registerSysContentRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/syscontent").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/:id", syscontent.GetSysContent)
		r.POST("", syscontent.InsertSysContent)
		r.PUT("", syscontent.UpdateSysContent)
		r.DELETE("/:id", syscontent.DeleteSysContent)
	}

	l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		l.GET("/syscontentList", syscontent.GetSysContentList)
	}

}
