package bootstrap

import (
	"database/sql"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"net/http"

	_ "github.com/lib/pq"
)

func InitApp(db *sql.DB) {
	// init Repository
	productRepo := repositories.NewProductRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)

	// init Service
	productService := services.NewProductService(productRepo, categoryRepo)
	categoryService := services.NewCategoryService(categoryRepo, productRepo)

	// init Controller
	handlers.InitProductHandler(productService)
	handlers.InitCategoryHandler(categoryService)

	// init Routes
	handlers.RegisterProductRoutes()
	handlers.RegisterCategoryRoutes()

	// Transaction
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	http.HandleFunc("/api/checkout", transactionHandler.HandleCheckout) // POST

	// Report
	reportRepoRepo := repositories.NewReportRepository(db)
	reportService := services.NewReportService(reportRepoRepo)
	reportHandler := handlers.NewReportHandler(reportService)

	http.HandleFunc("/api/report/hari-ini", reportHandler.GetTodayReport) // GET
	http.HandleFunc("/api/report", reportHandler.GetReport)               // GET with query params
}
