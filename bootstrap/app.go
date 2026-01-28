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
	produkRepo := repositories.NewProdukRepository(db)
	kategoriRepo := repositories.NewKategoriRepository(db)

	// init Service
	produkService := services.NewProdukService(produkRepo, kategoriRepo)
	kategoriService := services.NewKategoriService(kategoriRepo, produkRepo)

	// init Controller
	controllers.InitProdukController(produkService)
	controllers.InitKategoriController(kategoriService)

	// init Routes
	controllers.RegisterProdukRoutes()
	controllers.RegisterKategoriRoutes()
}
