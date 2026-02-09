package routes

const (
	basePath = "/api"

	// Category
	Category     = basePath + "/categories"
	CategoryByID = Category + "/"

	// Product
	Product     = basePath + "/products"
	ProductByID = Product + "/"

	// Transaction
	Checkout = basePath + "/checkout"

	// Report
	ReportToday = basePath + "/report/hari-ini"
	Report      = basePath + "/report"
)
