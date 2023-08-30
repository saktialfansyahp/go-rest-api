# Go Rest Api

Selamat datang di Proyek Nama Proyek Anda! Proyek ini berisi implementasi backend untuk layanan Anda. Di bawah ini, Anda akan menemukan informasi tentang cara menjalankan proyek, endpoint yang tersedia, dan bagaimana cara menggunakannya.

## Menjalankan Proyek

1. Pastikan Anda memiliki Go (Golang) diinstal pada mesin Anda.
2. Clone repositori ini ke direktori lokal Anda.
3. Buka terminal dan navigasikan ke direktori repositori yang baru saja Anda klon.
4. Jalankan perintah berikut untuk menginstal dependensi:

   ```sh

   go mod download

   ```

5. Setelah dependensi terinstal, jalankan perintah berikut untuk memulai server:

   ```sh

   go run main.go
   Server akan berjalan pada alamat http://localhost:8080.

   ```

# Endpoint

1. GET /home: Rute untuk halaman beranda.
2. POST /login: Rute untuk melakukan login pengguna.
3. POST /register: Rute untuk mendaftarkan pengguna baru.
4. GET /logout: Rute untuk melakukan logout pengguna.
5. POST /role: Rute untuk membuat peran baru.

   Kemudian, Grup rute api untuk melindungi beberapa rute dengan middleware:

6. GET /api/admin/product: Rute yang hanya dapat diakses oleh pengguna dengan peran "admin". Di sini Anda menggunakan middleware untuk memverifikasi peran pengguna.
7. GET /api/user/productUser: Rute yang dapat diakses oleh pengguna dengan peran "admin" dan "user". Juga menggunakan middleware untuk verifikasi peran.
