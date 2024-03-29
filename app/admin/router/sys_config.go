package router

import (
	"github.com/gin-gonic/gin"
	"fiy/app/admin/apis/system/sys_config"
	"fiy/app/admin/middleware"
	"fiy/app/admin/models/system"
	"fiy/app/admin/service/dto"
	"fiy/common/actions"
	jwt "fiy/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysConfigRouter)
}

// 需认证的路由代码
func registerSysConfigRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	r := v1.Group("/config").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		model := &system.SysConfig{}
		r.GET("", actions.PermissionAction(), actions.IndexAction(model, new(dto.SysConfigSearch), func() interface{} {
			list := make([]system.SysConfig, 0)
			return &list
		}))
		r.GET("/:id", actions.PermissionAction(), actions.ViewAction(new(dto.SysConfigById), nil))
		r.POST("", actions.CreateAction(new(dto.SysConfigControl)))
		r.PUT("/:id", actions.PermissionAction(), actions.UpdateAction(new(dto.SysConfigControl)))
		r.DELETE("", actions.PermissionAction(), actions.DeleteAction(new(dto.SysConfigById)))
	}

	r1 := v1.Group("/configKey").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		model := &sys_config.SysConfig{}
		r1.POST("", model.GetSysConfigByKEYForService)
	}
}
