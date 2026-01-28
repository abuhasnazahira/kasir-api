package bootstrap

import (
	"database/sql"
	controllers "kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"

	_ "github.com/lib/pq"
)

func InitApp(db *sql.DB) {
	// init Repository
	productRepo := repositories.NewProdukRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)

	// init Service
	productService := services.NewProductService(productRepo, categoryRepo)
	categoryService := services.NewCategoryService(categoryRepo, productRepo)

	// init Controller
	controllers.InitProductHandler(productService)
	controllers.InitCategoryHandler(categoryService)

	// init Routes
	controllers.RegisterProductRoutes()
	controllers.RegisterCategoryRoutes()
}
