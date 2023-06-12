package routers

import (
	controller "golang_api/app/controllers"
	middleware "golang_api/app/middleware"
	"golang_api/app/repositories"
	"golang_api/app/services"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Add Middleware
	logger := middleware.LoggerMiddleware{}
	r.Use(middleware.ErrorHandler())
	r.Use(logger.Logger())
	r.Use(middleware.CheckToken())

	// Load instance UserRepository
	userRepository := repositories.NewUserRepository()
	roleRepository := repositories.NewRoleRepository()
	moduleRepository := repositories.NewModuleRepository()

	// Load instance UserService
	userService := services.NewUserService(*userRepository)
	authService := services.NewAuthenticationService(*userRepository)
	roleService := services.NewRoleService(*roleRepository)
	moduleService := services.NewModuleService(*moduleRepository)

	// Load instance UserController
	userCtrl := controller.NewUserController(userService)
	authCtrl := controller.NewAuthenticationController(authService)
	roleCtrl := controller.NewRoleController(roleService)
	moduleCtrl := controller.NewModuleController(moduleService)

	// Create group routing endpoint "/api/v1"
	v1 := r.Group("/api/v1")
	{
		employee := v1.Group("/user")
		{
			employee.GET("/:id", userCtrl.GetUser)
			employee.GET("/", userCtrl.GetAllUser)
			employee.POST("/", userCtrl.PostUser)
			employee.PUT("/:id", userCtrl.PutUser)
			employee.DELETE("/:id", userCtrl.DeleteUser)
		}

		auth := v1.Group("/auth")
		{
			auth.POST("/requestToken", authCtrl.Login)
			auth.POST("/refreshToken", authCtrl.RefeshToken)
			auth.GET("/external/google", authCtrl.Oauth2Login)
			auth.GET("/external/google-callback", authCtrl.Callback)
		}

		role := v1.Group("role")
		{
			role.GET("/", roleCtrl.GetRoles)
			role.GET("/:id", roleCtrl.GetRoleById)
			role.GET("/module/:name", roleCtrl.GetRoleWithModule)
			role.GET("/name/:name", roleCtrl.GetRoleByName)
			role.POST("/", roleCtrl.PostRole)
			role.PUT("/:id", roleCtrl.PutRole)
			role.PUT("/module/set/:id", roleCtrl.PutRoleSetModule)
			role.DELETE("/:id", roleCtrl.DeleteRole)
		}

		module := v1.Group("module")
		{
			module.GET("/", moduleCtrl.GetModules)
			module.GET("/:id", moduleCtrl.GetModuleById)
			module.GET("/name/:name", moduleCtrl.GetModuleByName)
			module.POST("/", moduleCtrl.PostModule)
			module.PUT("/:id", moduleCtrl.PutModule)
			module.DELETE("/:id", moduleCtrl.DeleteModule)
		}

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
