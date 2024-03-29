package router

import (
	"github.com/gin-gonic/gin"
	"fiy/app/admin/apis/system/sys_opera_log"
	"fiy/app/admin/middleware"
	jwt "fiy/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysOperaLogRouter)
}

// 需认证的路由代码
func registerSysOperaLogRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &sys_opera_log.SysOperaLog{}
	r := v1.Group("/sys-opera-log").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetSysOperaLogList)
		r.GET("/:id", api.GetSysOperaLog)
		r.POST("", api.InsertSysOperaLog)
		r.PUT("/:id", api.UpdateSysOperaLog)
		r.DELETE("", api.DeleteSysOperaLog)
	}
}
