package bootstrap

import (
	"kasir-api/controllers"
	"kasir-api/repositories"
	"kasir-api/services"
)

func InitApp() {
	// init Repository
	produkRepo := repositories.NewProdukRepository()
	kategoriRepo := repositories.NewKategoriRepository()

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
