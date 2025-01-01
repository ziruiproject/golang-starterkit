## ANSWERS
___
## **Indexing strategy**
**1. Fetch a user by username:**

- **Index:** Membuat index pada kolom `username` 
- **Alasan:** Karena username akan sering dicari (dengan WHERE clause) maka ada baiknya untuk mengindex `username` sehingga cepat dalam mengambil data user saat misalnya login

**2. Fetch users who signed up after a certain date (`created_at > '2023-01-01'`):**

- **Index:** Dengan membuat non unique index pada `created_at` 
- **Alasan:** Index pada `created_at` akan mempercepat proses filtering berdasarkan tanggal 

**3. Fetch a user by email:**

- **Index:** Membuat index pada kolom `email`
- **Alaasan:** Sama dengan kolom username, membuat index pada kolom `email` akan mempercepat pencarian by email yang dimana kolom ini biasanay sering dilakukan lookup

**Apakah Perlu Composite Index?**

- **Jawab:** Jika memang sering mencari / memfilter dengan 2 kolom, misal nya mencari dengan kolom `email` dan `created_at` untuk mencari user yang mendaftar pada rentang waktu tertentu, maka composite index akan bermanfaat

**Konsekuensi:**

- **Read Performance:** Index (biasanya) bekerja dengan cara membuat Balanced Tree untuk kolom yang di index. Hal ini dapat mempercepat proses baca data karena databse tidak lagi mencari secara linear / sequential
- **Write Performance:** Namun Balanced Tree tersebut akan dibuat ulang setiap kali ada data yang masuk. Karena itu jika kita menerapkan index, maka performa saat memasukan / mengubah data (`INSERT`, `UPDATE`) akan berkurang
- **Storage Overhead:** Karena index menghadirkan struktur data baru, pastinya akan memakan lebih banyak tempat. Semakin banyak index yang dibuat, semakin banyak storage yg digunakan
___
## **Optimized SQL Query to Find Top 5 Customers Who Spent the Most Money in the Past Month**
```postgresql
    SELECT customer_id, SUM(amount) AS total_spent
    FROM orders
    WHERE order_date >= CURRENT_DATE - INTERVAL '1 month'
    GROUP BY customer_id
    ORDER BY total_spent DESC
    LIMIT 5;
```
**Penjelasan:**
- `SELECT customer_id, SUM(amount) AS total_spent`: Mengambil `customer_id` dan menjumlah total pengeluaran customer
- `FROM orders`: Memilih tabel `orders` sebagai data sumber.
- `WHERE order_date >= CURDATE() - INTERVAL 1 MONTH`: Memfilter hanya data pada 
- `GROUP BY customer_id`: Mengelompokan hasil berdasarkan `customer_id` agar dapat menjumlah pengeluaran per customer
- `ORDER BY total_spent DESC`: Mengurutkan data berdasarkan kolom total_spent dan diurutkan secara descending (menurun, dari yang terbesar)
- `LIMIT 5`: Hanya ambil 5 data pertama

**Improvement:**

Karena `customer_id` query bisa dengan mudah digrup / kelompokan bersadarkan kolom ini. Namun kolom `amount` dan `order_date` tidak terindex, yang akan memengaruhi peforma di dataset yang besar. Untuk lebih optimal lagi:

1. Buat index pada kolom `order_date` agar mempercepat filtering.
   ```sql
   CREATE INDEX idx_order_date ON orders(order_date);
   ```
2. Jika dataset terlalu besar, bisa dicoba dengam melakukan partisi table `orders` dengan `order_date`.
___
## **Approach to Refactoring a Monolithic Service into Microservices**

Cara merefactor Monolith ke Arsitektur yang baru:

1. **Mengelompokan Service:** Dengan menggunakan DDD (Domain Driven Design) kita bisa mengelompokan tanggung jawab masing-masing service di Monolith yang sudah ada untuk kemudian dipisah di arsitektur yang baru
2. **Cicil Pembuatan:** Dengan menerapkan Strangler Pattern, yakni dengan cara menulis ulang sebagian service, kemudian me-replace service tersebut dengan service yang baru. Sehingga service yang lama akan berjalan seiiringan dengan service baru sampai pada akhirnya semua service lama telah selesai ditulis ulang / direfactor
3. **Backward Compatible:** Agar backward compatible, saat masa transisi perlu menerapkan API Versioning (/api/v1, /api/v2) agar bisa rollback jika dibutuhkan