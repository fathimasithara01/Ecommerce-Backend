package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fathimasithara01/ecommerce/middleware"
	controllers "github.com/fathimasithara01/ecommerce/src/controllers/admin"
	repository "github.com/fathimasithara01/ecommerce/src/repository/admin"
	services "github.com/fathimasithara01/ecommerce/src/services/admin"
)

func AdminRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	//adminRoutes
	adminRepo := &repository.AdminRepository{DB: db}
	authService := &services.AuthService{Repo: adminRepo}
	authHandler := &controllers.AuthHandler{Services: authService}

	adminAuth := r.Group("/Authentication")
	{
		adminAuth.POST("/signup", authHandler.SignupAdmin)
		adminAuth.POST("/login", authHandler.LoginAdmin)
	}

	// product routes
	productRepo := repository.NewProductRepository(db)
	productUsecase := services.NewProductUsecase(productRepo)
	productHandler := controllers.NewProductHandler(productUsecase)

	product := r.Group("/products", middleware.AdminAuthMiddleware())
	{
		product.POST("/createProduct", productHandler.Create)
		product.GET("/getAllProducts", productHandler.GetAll)
		product.PUT("/:id", productHandler.Update)
		product.DELETE("/:id", productHandler.Delete)
	}

	// category routes
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := controllers.NewCategoryHandler(categoryService)
	admin := r.Group("/category", middleware.AdminAuthMiddleware())
	{
		admin.POST("/createCategory", categoryHandler.CreateCategory)
		admin.PUT("/update/:id", categoryHandler.UpdateCategory)
		admin.DELETE("/delete/:id", categoryHandler.DeleteCategory)
		admin.GET("/getAllCategories", categoryHandler.GetAllCategories)
		admin.GET("/:id", categoryHandler.GetCategoryByID)
	}

	// user routes
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserUsecase(userRepo)
	userHandler := controllers.NewUserHandler(userService)
	userGroup := r.Group("/users", middleware.AdminAuthMiddleware())
	{
		userGroup.GET("/getAllUsers", userHandler.GetAllUsers)
		userGroup.PUT("/:id/unblock", userHandler.UnblockUser)
		userGroup.PUT("/:id/block", userHandler.BlockUser)
		userGroup.DELETE("/delete/:id", userHandler.DeleteUser)
	}

	// order routes
	orderRepo := repository.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderHandler := controllers.NewOrderHandleer(orderService)

	order := r.Group("/orders", middleware.AdminAuthMiddleware())
	{
		order.GET("/getAllOrders", orderHandler.GetAllOrder)
		order.GET("/:id", orderHandler.GetOrderByID)
		order.PUT("/:id/status", orderHandler.UpdateOrderStatus)
		order.DELETE("/:id", orderHandler.DeleteOrder)
	}

	// dashboard routes
	dashboardRepo := repository.NewDashboardRepository(db)
	dashboardService := services.NewDashboardService(dashboardRepo)
	dashboardHandler := controllers.NewDashboardHandler(dashboardService)
	dashboard := r.Group("/dashboard", middleware.AdminAuthMiddleware())
	{
		dashboard.GET("/", dashboardHandler.GetDashboard)
	}

	return r
}
