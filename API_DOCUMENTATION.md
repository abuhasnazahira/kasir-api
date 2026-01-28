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

**Request:**
```bash
curl -X GET http://localhost:8080/api/categories
```

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "name": "Makanan",
    "description": "Produk makanan"
  },
  {
    "id": 2,
    "name": "Minuman",
    "description": "Produk minuman"
  }
]
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
  "id": 1,
  "name": "Makanan",
  "description": "Produk makanan"
}
```

**Response (404 Not Found):**
```json
{
  "error": "kategori tidak ditemukan"
}
```

---

### 3. Create Kategori
**Endpoint:** `POST /api/categories`

**Description:** Membuat kategori baru

**Request Body:**
```json
{
  "name": "Makanan",
  "description": "Produk makanan"
}
```

**Request:**
```bash
curl -X POST http://localhost:8080/api/categories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Makanan",
    "description": "Produk makanan"
  }'
```

**Response (201 Created):**
```json
{
  "id": 1,
  "name": "Makanan",
  "description": "Produk makanan"
}
```

**Response (400 Bad Request):**
```json
{
  "error": "Invalid Request Body"
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
  "id": 1,
  "name": "Makanan Berat",
  "description": "Produk makanan berat"
}
```

**Response (400 Bad Request):**
```json
{
  "error": "Invalid Request Body"
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
  "message": "Sukses delete"
}
```

**Response (404 Not Found):**
```json
{
  "error": "kategori tidak ditemukan"
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

**Request:**
```bash
curl -X GET "http://localhost:8080/api/product?limit=10&offset=0"
```

**Response (200 OK):**
```json
{
  "data": [
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
    },
    {
      "id": 2,
      "name": "Mie Goreng",
      "price": 20000,
      "stock": 15,
      "category": {
        "id": 1,
        "name": "Makanan",
        "description": "Produk makanan"
      }
    }
  ],
  "total": 50,
  "limit": 10,
  "offset": 0
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
  "error": "Product tidak ditemukan"
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
  "error": "Kategori tidak ditemukan"
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
  "name": "Nasi Goreng Special",
  "price": 30000,
  "stock": 20,
  "category": {
    "categoryId": 1
  }
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
```

**Response (404 Not Found):**
```json
{
  "error": "Product tidak ditemukan"
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
  "message": "Sukses delete"
}
```

**Response (404 Not Found):**
```json
{
  "error": "Product tidak ditemukan"
}
```

---

## Error Response Format

Semua error response mengikuti format:

```json
{
  "error": "Deskripsi error",
  "status": "error"
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

### Kategori Model
```json
{
  "categoryId": 1,
  "name": "Makanan",
  "description": "Produk makanan"
}
```

### Produk Model
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
