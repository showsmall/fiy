package router

import (
	"github.com/gin-gonic/gin"
	"fiy/app/admin/apis/system/sys_login_log"
	"fiy/app/admin/middleware"
	jwt "fiy/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysLoginLogRouter)
}

// 需认证的路由代码
func registerSysLoginLogRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &sys_login_log.SysLoginLog{}
	r := v1.Group("/sys-login-log").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetSysLoginLogList)
		r.GET("/:id", api.GetSysLoginLog)
		r.POST("", api.InsertSysLoginLog)
		r.PUT("/:id", api.UpdateSysLoginLog)
		r.DELETE("", api.DeleteSysLoginLog)
	}
}
