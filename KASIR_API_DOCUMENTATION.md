# Kasir API Documentation

API documentation untuk Kasir API dengan endpoint untuk mengelola kategori dan produk.

---

## Base URL
```
http://localhost:8080/api
```

---

## Kategori Endpoints

### 1. Get All Kategori
**Endpoint:** `GET /api/categories`

**Description:** Mengambil semua data kategori

**Request Body:**
```json
{
  "search": "",
  "limit": 10,
  "offset": 0
}
```

**Request:**
```bash
curl -X GET http://localhost:8080/api/categories
```

**Response (200 OK):**
```json
{
	"payload": {
		"data": [
			{
				"id": 3,
				"name": "Minuman",
				"description": "Produk Minuman"
			},
			{
				"id": 4,
				"name": "Makanan",
				"description": "Produk Makanan"
			},
			{
				"id": 8,
				"name": "Sepeda",
				"description": "Produk Makanan"
			}
		],
		"totalRecordFiltered": 3,
		"totalRecords": 3
	},
	"responseCode": 200,
	"responseMessage": "success"
}
```

---

### 2. Get Kategori By ID
**Endpoint:** `GET /api/categories/{id}`

**Description:** Mengambil data kategori berdasarkan ID

**Parameters:**
| Name | Type | Description |
|------|------|-------------|
| id | integer | ID kategori |

**Request:**
```bash
curl -X GET http://localhost:8080/api/categories/1
```

**Response (200 OK):**
```json
{
	"payload": {
		"data": {
			"id": 3,
			"name": "Minuman",
			"description": "Produk Minuman"
		}
	},
	"responseCode": 200,
	"responseMessage": "success"
}
```

**Response (404 Not Found):**
```json
{
	"responseCode": 404,
	"responseMessage": "Kategori tidak ditemukan"
}
```

---

### 3. Create Kategori
**Endpoint:** `POST /api/categories`

**Description:** Membuat kategori baru

**Request Body:**
```json
{
		"name": "Furniture",
		"description": "Peralatan Rumah Tangga"
}
```

**Request:**
```bash
curl -X POST http://localhost:8080/api/categories \
  -H "Content-Type: application/json" \
  -d '{
		"name": "Furniture",
		"description": "Peralatan Rumah Tangga"
}'
```

**Response (201 Created):**
```json
{
	"payload": {
		"data": {
			"id": 11,
			"name": "Furniture",
			"description": "Peralatan Rumah Tangga"
		}
	},
	"responseCode": 201,
	"responseMessage": "success"
}
```

**Response (400 Bad Request):**
```json
{
  "responseCode": 400,
  "responseMessage": "Invalid Request Body"
}
```

---

### 4. Update Kategori
**Endpoint:** `PUT /api/categories/{id}`

**Description:** Memperbarui data kategori

**Parameters:**
| Name | Type | Description |
|------|------|-------------|
| id | integer | ID kategori |

**Request Body:**
```json
{
  "name": "Makanan Berat",
  "description": "Produk makanan berat"
}
```

**Request:**
```bash
curl -X PUT http://localhost:8080/api/categories/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Makanan Berat",
    "description": "Produk makanan berat"
  }'
```

**Response (200 OK):**
```json
{
	"payload": {
		"data": {
			"id": 1,
			"name": "Makanan Berat",
      "description": "Produk makanan berat"
		}
	},
	"responseCode": 201,
	"responseMessage": "success"
}
```

**Response (400 Bad Request):**
```json
{
  "responseCode": 400,
  "responseMessage": "Invalid Request Body"
}
```

---

### 5. Delete Kategori
**Endpoint:** `DELETE /api/categories/{id}`

**Description:** Menghapus kategori berdasarkan ID

**Parameters:**
| Name | Type | Description |
|------|------|-------------|
| id | integer | ID kategori |

**Request:**
```bash
curl -X DELETE http://localhost:8080/api/categories/1
```

**Response (200 OK):**
```json
{
  "responseCode": 200,
  "responseMessage": "Sukses delete"
}
```

**Response (404 Not Found):**
```json
{
	"responseCode": 404,
	"responseMessage": "Kategori tidak ditemukan"
}
```

---

## Produk Endpoints

### 1. Get All Produk (dengan Pagination)
**Endpoint:** `GET /api/product`

**Description:** Mengambil semua data produk dengan pagination

**Query Parameters:**
| Name | Type | Default | Description |
|------|------|---------|-------------|
| limit | integer | 10 | Jumlah data per halaman |
| offset | integer | 0 | Jumlah data yang di-skip |

**Request Body:**
```json
{
  "search": "",
  "limit": 10,
  "offset": 0
}
```

**Request:**
```bash
curl -X GET http://localhost:8080/api/product \
  -H "Content-Type: application/json" \
  -d '{
    "search": "",
    "limit": 10,
    "offset": 0
  }'
```

**Response (200 OK):**
```json
{
	"payload": {
		"data": [
			{
				"id": 1,
				"name": "Coca Cola 1L",
				"price": 3000,
				"stock": 20,
				"category": {
					"id": 3,
					"name": "Minuman",
					"description": "Produk Minuman"
				}
			},
			{
				"id": 2,
				"name": "Chiki balls",
				"price": 2000,
				"stock": 40,
				"category": {
					"id": 4,
					"name": "Makanan",
					"description": "Produk Makanan"
				}
			},
			{
				"id": 3,
				"name": "Taro Snack",
				"price": 2000,
				"stock": 40,
				"category": {
					"id": 4,
					"name": "Makanan",
					"description": "Produk Makanan"
				}
			},
			{
				"id": 4,
				"name": "Coca Cola 1L",
				"price": 3000,
				"stock": 20,
				"category": {
					"id": 3,
					"name": "Minuman",
					"description": "Produk Minuman"
				}
			},
			{
				"id": 5,
				"name": "Coca Cola 1L",
				"price": 3000,
				"stock": 20,
				"category": {
					"id": 3,
					"name": "Minuman",
					"description": "Produk Minuman"
				}
			},
			{
				"id": 6,
				"name": "Coca Cola 1L",
				"price": 3000,
				"stock": 20,
				"category": {
					"id": 3,
					"name": "Minuman",
					"description": "Produk Minuman"
				}
			}
		],
		"totalRecordFiltered": 6,
		"totalRecords": 6
	},
	"responseCode": 200,
	"responseMessage": "success"
}
```

---

### 2. Get Produk By ID
**Endpoint:** `GET /api/product/{id}`

**Description:** Mengambil data produk berdasarkan ID

**Parameters:**
| Name | Type | Description |
|------|------|-------------|
| id | integer | ID produk |

**Request:**
```bash
curl -X GET http://localhost:8080/api/product/1
```

**Response (200 OK):**
```json
{
  "id": 1,
  "name": "Nasi Goreng",
  "price": 25000,
  "stock": 10,
  "category": {
    "id": 1,
    "name": "Makanan",
    "description": "Produk makanan"
  }
}
```

**Response (404 Not Found):**
```json
{
  	"responseCode": 404,
	"responseMessage": "Product tidak ditemukan"
}
```

---

### 3. Create Produk
**Endpoint:** `POST /api/product`

**Description:** Membuat produk baru

**Request Body:**
```json
{
  "name": "Nasi Goreng",
  "price": 25000,
  "stock": 10,
  "category": {
    "id": 1
  }
}
```

**Request:**
```bash
curl -X POST http://localhost:8080/api/product \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Nasi Goreng",
    "price": 25000,
    "stock": 10,
    "category": {
      "id": 1
    }
  }'
```

**Response (201 Created):**
```json
{
  "productId": 1,
  "name": "Nasi Goreng",
  "price": 25000,
  "stock": 10,
  "category": {
    "categoryId": 1,
    "name": "Makanan",
    "description": "Produk makanan"
  }
}
```

**Response (400 Bad Request):**
```json
{
  	"responseCode": 404,
	"responseMessage": "Product tidak ditemukan"
}
```

---

### 4. Update Produk
**Endpoint:** `PUT /api/product/{id}`

**Description:** Memperbarui data produk

**Parameters:**
| Name | Type | Description |
|------|------|-------------|
| id | integer | ID produk |

**Request Body:**
```json
{
	"payload": {
		"data": [
			{
        "name": "Nasi Goreng Special",
        "price": 30000,
        "stock": 20,
        "category": {
          "categoryId": 1
        }
      }
		],
		"totalRecordFiltered": 3,
		"totalRecords": 3
	},
	"responseCode": 200,
	"responseMessage": "success"
}
```

**Request:**
```bash
curl -X PUT http://localhost:8080/api/product/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Nasi Goreng Special",
    "price": 30000,
    "stock": 20,
    "category": {
      "id": 1
    }
  }'
```

**Response (200 OK):**
```json
{
	"payload": {
		"data": [
			{
        "productId": 1,
        "name": "Nasi Goreng Special",
        "price": 30000,
        "stock": 20,
        "category": {
          "categoryId": 1,
          "name": "Makanan",
          "description": "Produk makanan"
        }
      }
		],
		"totalRecordFiltered": 3,
		"totalRecords": 3
	},
	"responseCode": 200,
	"responseMessage": "success"
}
```

**Response (404 Not Found):**
```json
{
  	"responseCode": 404,
	"responseMessage": "Product tidak ditemukan"
}
```

---

### 5. Delete Produk
**Endpoint:** `DELETE /api/product/{id}`

**Description:** Menghapus produk berdasarkan ID

**Parameters:**
| Name | Type | Description |
|------|------|-------------|
| id | integer | ID produk |

**Request:**
```bash
curl -X DELETE http://localhost:8080/api/product/1
```

**Response (200 OK):**
```json
{
  "responseCode": 200,
  "responseMessage": "Sukses delete"
}
```

**Response (404 Not Found):**
```json
{
  "responseCode": 404,
	"responseMessage": "Product tidak ditemukan"
}
```

## Transaction Endpoints

### 1. Checkout
**Endpoint:** `POST /api/checkout`

**Description:** Memproses permintaan checkout dengan menerima data item (product_id dan quantity), melakukan validasi stok, menghitung subtotal dan total transaksi, lalu membuat record transaksi dan detail transaksi.

**Query Parameters:**
| Name | Type | Default | Description |
|------|------|---------|-------------|
| product_id | integer | - | ID produk yang akan dibeli |
| quantity | integer | 1 | Jumlah unit produk yang dibeli |

**Request Body:**
```json
{
  "items": [
    {
      "product_id": 6,
      "quantity": 2
    },
    {
      "product_id": 7,
      "quantity": 2
    }
  ]
}
```

**Request:**
```bash
curl -X POST http://localhost:8080/api/checkout \
  -H "Content-Type: application/json" \
  -d '{
  "items": [
    {
      "product_id": 6,
      "quantity": 2
    },
    {
      "product_id": 7,
      "quantity": 2
    }
  ]
}
'
```

**Response (200 OK):**
```json
{
	"payload": {
		"data": {
			"id": 3,
			"total_amount": 10000,
			"created_at": "0001-01-01T00:00:00Z",
			"details": [
				{
					"id": 5,
					"transaction_id": 3,
					"product_id": 6,
					"product_name": "Coca Cola 1L",
					"quantity": 2,
					"subtotal": 6000
				},
				{
					"id": 6,
					"transaction_id": 3,
					"product_id": 7,
					"product_name": "Taro Snack",
					"quantity": 2,
					"subtotal": 4000
				}
			]
		}
	},
	"responseCode": 201,
	"responseMessage": "success"
}
```

**Response (500 Internal Server Error):**
```json
{
  	"responseCode": 500,
	"responseMessage": "Internal Server Error"
}
```

## Report Endpoints

### 1. Report Sales Summary Today
**Endpoint:** `GET /api/report/hari-ini`

**Description:** Mengambil ringkasan laporan penjualan untuk hari ini, termasuk total transaksi, total pendapatan, dan jumlah produk yang terjual berdasarkan data transaksi yang tercatat di sistem.

**Request:**
```bash
curl -X GET http://localhost:8080/api/report/hari-ini \
  -H "Content-Type: application/json" \
```

**Response (200 OK):**
```json
{
	"payload": {
		"data": {
			"total_revenue": 0,
			"total_transaksi": 0,
			"produk_terlaris": {
				"nama": "",
				"qty_terjual": 0
			}
		}
	},
	"responseCode": 200,
	"responseMessage": "success"
}
```

**Response (500 Internal Server Error):**
```json
{
  	"responseCode": 500,
	"responseMessage": "Internal Server Error"
}
```

### 2. Report Sales Summary Date Range Filter
**Endpoint:** `GET /api/report?start_date=2026-01-01&end_date=2026-02-01`

**Description:** Mengambil ringkasan laporan penjualan berdasarkan rentang tanggal tertentu, termasuk total transaksi, total pendapatan, dan jumlah produk yang terjual dari data transaksi yang tercatat dalam sistem.

**Query Parameters:**
| Name | Type | Default | Description |
|------|------|---------|-------------|
| start_date | string | - | Tanggal awal periode laporan penjualan (format: YYYY-MM-DD) |
| end_date | string | - | Tanggal akhir periode laporan penjualan (format: YYYY-MM-DD) |


**Request:**
```bash
curl -X GET http://localhost:8080/api/report?start_date=2026-01-01&end_date=2026-02-01 \
  -H "Content-Type: application/json" \
```

**Response (200 OK):**
```json
{
	"payload": {
		"data": {
			"total_revenue": 36000,
			"total_transaksi": 3,
			"produk_terlaris": {
				"nama": "Coca Cola 1L",
				"qty_terjual": 6
			}
		}
	},
	"responseCode": 200,
	"responseMessage": "success"
}
```

**Response (500 Internal Server Error):**
```json
{
  	"responseCode": 500,
	"responseMessage": "Internal Server Error"
}
```

---

## Error Response Format

Semua error response mengikuti format:

```json
{
  	"responseCode": "errorCode",
	"responseMessage": "Deskripsi error"
}
```

### Common Error Codes

| Status Code | Description |
|-------------|-------------|
| 200 | OK - Request berhasil |
| 201 | Created - Resource berhasil dibuat |
| 400 | Bad Request - Input tidak valid |
| 404 | Not Found - Resource tidak ditemukan |
| 500 | Internal Server Error - Error di server |

---

## Model Data

### Category Model
```json
{
	"categoryId": 1,
	"name": "Makanan",
	"description": "Produk makanan"
}
```

### Product Model
```json
{
	"productId": 1,
	"name": "Nasi Goreng",
	"price": 25000,
	"stock": 10,
	"category": {
		"categoryId": 1,
		"name": "Makanan",
		"description": "Produk makanan"
	}
}
```

### Transaction Model
```json
{
  	"id": 3,
	"total_amount": 10000,
	"created_at": "0001-01-01T00:00:00Z",
	"details": [
		{
			"id": 5,
			"transaction_id": 3,
			"product_id": 6,
			"product_name": "Coca Cola 1L",
			"quantity": 2,
			"subtotal": 6000
		},
		{
			"id": 6,
			"transaction_id": 3,
			"product_id": 7,
			"product_name": "Taro Snack",
			"quantity": 2,
			"subtotal": 4000
		}
	]
}
```

### Transaction Detail Model
```json
{
	"id": 5,
	"transaction_id": 3,
	"product_id": 6,
	"product_name": "Coca Cola 1L",
	"quantity": 2,
	"subtotal": 6000
}
```

---

## Endpoint Summary

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/categories` | Get all kategori |
| GET | `/api/categories/{id}` | Get kategori by ID |
| POST | `/api/categories` | Create kategori |
| PUT | `/api/categories/{id}` | Update kategori |
| DELETE | `/api/categories/{id}` | Delete kategori |
| GET | `/api/product` | Get all produk (with pagination) |
| GET | `/api/product/{id}` | Get produk by ID |
| POST | `/api/product` | Create produk |
| PUT | `/api/product/{id}` | Update produk |
| DELETE | `/api/product/{id}` | Delete produk |
| POST | `/api/checkout` | Create Checkout Transaction |
| GET | `/api/report/hari-ini` | Get Report Sales Summary current date|
| GET | `/api/report?start_date={start_date}&end_date={end_date}` | Get Report Sales Summary with date range filter|

---

## Pagination

Endpoint GET `/api/product` mendukung pagination dengan query parameters:

```bash
curl -X GET "http://localhost:8080/api/product?limit=5&offset=10"
```

Response akan include:
- `data`: Array of products
- `total`: Total number of products
- `limit`: Limit per page
- `offset`: Current offset

---

## Notes

- Semua request body harus dalam format JSON
- Semua response dalam format JSON
- Database menggunakan PostgreSQL
- Kategori ID harus valid saat membuat atau mengupdate produk
