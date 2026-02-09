package bootstrap

import (
	"database/sql"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/routes"
	"kasir-api/services"
	"net/http"

	_ "github.com/lib/pq"
)

func InitApp(db *sql.DB) {
	// Category
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	http.HandleFunc(routes.CategoryByID, categoryHandler.HandleCategoryId) // GET / PUT / DELETE by ID
	http.HandleFunc(routes.Category, categoryHandler.HandleCategories)     // GET all categories / POST new product

	// Product
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo, categoryRepo)
	productHandler := handlers.NewProductHandler(productService)

	http.HandleFunc(routes.ProductByID, productHandler.HandleProductId) // GET / PUT / DELETE by ID
	http.HandleFunc(routes.Product, productHandler.HandleProducts)      // GET all products / POST new product

	// Transaction
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	http.HandleFunc(routes.Checkout, transactionHandler.HandleCheckout) // POST checkout transaction

	// Report
	reportRepoRepo := repositories.NewReportRepository(db)
	reportService := services.NewReportService(reportRepoRepo)
	reportHandler := handlers.NewReportHandler(reportService)

	http.HandleFunc(routes.ReportToday, reportHandler.GetTodayReport) // GET Report Today
	http.HandleFunc(routes.Report, reportHandler.GetReport)           // GET Report with date range filter
}
