# golang-final-project4-team2

Link Deploy API : https://golang-final-project4-team2-production.up.railway.app/

Repository Team 2 Untuk Final Project 4 Golang (Kampus Merdeka Hacktiv8)

Anggota Kelompok :

- JOVIN LIDAN (GLNG-KS04-023)
  Mengerjakan :
  - Setup Docker dan Init Project
  - Semua API User dan Postman user
  - Semua API Categories dan Postman categories
  - Semua API Products dan Postman products
  - Setup deployment menggunakan railway
  - Helper : GenerateToken, VerifyToken, ValidateRequest, ComparePass, HashPass
  - Helper_test : TestSuccessGenerateToken, TestFailedGenerateToken, TestSuccessComparePass, TestFailedComparePass, TestSuccessHashPass, TestFailedHashPass
  - Middleware auth dan Middleware only admin
  - Setup create 4 tabel di migrations.go
- GUSTIO NUSAMBA (GLNG-KS04-025)
  Mengerjakan :
  - Semua API Transaction History dan Postman transaction history

## Cara Install

1. Buka dan jalankan aplikasi docker.
2. Jalankan command `docker compose up --build` untuk menjalankan database postgres di dalam docker container , go dan air auto reload. Tunggu agar docker sudah berjalan dengan baik.
3. Setelah docker container semuanya berjalan dengan baik. Maka port default yang akan dibuka adalah `8080`

_Note : Memerlukan docker terinstall didalam perangkat anda_

_Nama File Postman : `Toko_belanja.postman_collection.json`_

```json
Akun Admin:
Email     : admin@gmail.com
Password  : admin12
```

## List Route
### Users
- **`POST`- Users Register `api/users/register`**, Digunakan untuk membuat user baru.
- **`POST`- Users Login `api/users/login`**, Digunakan untuk melakukan login atau autentikasi user.
- **`PATCH`- Users Update `api/users/topup`**, Digunakan untuk menambahkan balance user.

### Categories
- **`GET`- Categories Index `api/categories`**, Digunakan untuk mengambil seluruh data categories dari database.
- **`POST`- Categories Store `api/categories`**, Digunakan untuk membuat category baru.
- **`PATCH`- Categories Update `api/categories/:categoryId`**, Digunakan untuk mengubah data category berdasarkan idnya.
- **`DELETE`- Categories Delete `api/categories/:categoryId`**, Digunakan untuk menghapus data category berdasarkan idnya.

### Products
- **`GET`- Products Index `api/products`**, Digunakan untuk mengambil seluruh data products dari database.
- **`POST`- Products Store `api/products`**, Digunakan untuk membuat product baru.
- **`PUT`- Products Update `api/products/:productId`**, Digunakan untuk mengubah data product berdasarkan idnya.
- **`DELETE`- Products Delete `api/products/:productId`**, Digunakan untuk menghapus data product berdasarkan idnya.

### Transaction History
- **`GET`-Transaction History User `api/transaction/my-transactions`**, Berguna untuk user yang sedang login mendapatkan data-data transaksi pembeliannya.
- **`POST`-Transaction History Store `api/transaction`**, Digunakan untuk membuat transaction history baru.
- **`GET`-Transaction History Admin `api/transaction/user-transactions`**, Berguna untuk admin melihat seluruh data transaksi pembelian product oleh user.
