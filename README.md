# Goku

Microframework Go tanpa dependency eksternal untuk proyek web lokal yang ringan.

## Fitur awal

- Router `GET` dan `POST`, termasuk parameter seperti `/produk/:id`.
- Middleware dan pemulihan panic.
- Respons teks dan JSON.
- Template HTML berbasis `html/template` dari Go.
- CLI pembuat proyek: `goku new <nama-proyek>`.

## Menjalankan contoh

```bash
go test ./...
go run ./examples/toko
```

Kemudian buka `http://localhost:8080`.

## Membuat proyek

```bash
go run ./cmd/goku new toko-saya
```

Generator menghasilkan proyek yang membutuhkan modul ini. Saat modul sudah dirilis ke GitHub, jalankan `go mod tidy` di proyek baru untuk mengambil versinya. Untuk pengembangan lokal, gunakan `replace` di `go.mod` proyek Anda menuju folder checkout `Goku`.

## Batasan v0.1

Belum ada middleware autentikasi, migrasi database, validasi formulir, maupun integrasi MariaDB/SQLite. Fitur tersebut adalah tahap berikutnya setelah API router stabil.
