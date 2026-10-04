# Berkontribusi ke Goku

Terima kasih sudah ingin membantu Goku. Proyek ini bertujuan menyediakan microframework Go yang kecil, mudah dipahami, dan tidak bergantung pada paket eksternal.

## Sebelum mulai

- Cari issue yang belum diambil, terutama yang berlabel `good first issue`.
- Untuk perubahan besar seperti ORM, autentikasi, atau perubahan API publik, buat issue diskusi terlebih dahulu agar arah implementasinya disepakati.
- Jangan menambahkan dependency eksternal tanpa alasan yang kuat. Prioritaskan library standar Go.

## Menyiapkan lingkungan kerja

Goku memerlukan Go 1.22 atau lebih baru.

```bash
git clone https://github.com/sirojmu/Goku.git
cd Goku
go test ./...
```

Untuk mencoba aplikasi contoh:

```bash
go run ./examples/toko
```

Lalu buka `http://localhost:8080` di peramban.

## Aturan perubahan

- Gunakan nama dan pesan galat yang jelas; pesan untuk pengguna diutamakan dalam Bahasa Indonesia.
- Jalankan `gofmt -w <berkas-go-yang-diubah>` sebelum mengirim perubahan.
- Tambahkan atau perbarui tes untuk perubahan perilaku router, context, middleware, atau renderer.
- Hindari mengubah API publik yang sudah ada tanpa mencatatkannya di README atau issue terkait.
- Jangan sertakan binary, kredensial, token, atau berkas konfigurasi lokal.

## Mengirim Pull Request

1. Fork repository ini dan buat branch dari `main`.
2. Buat perubahan yang fokus pada satu masalah.
3. Jalankan `go test ./...` dan pastikan seluruh tes lulus.
4. Jelaskan tujuan perubahan, cara mengujinya, dan issue yang ditangani pada deskripsi Pull Request.
5. Tunggu peninjauan; maintainer dapat meminta penyederhanaan, tes tambahan, atau pembaruan dokumentasi.

## Area kontribusi v0.1

- Tambahan metode HTTP dan penanganan `404`/`405` yang lebih baik.
- Middleware logging dan konfigurasi server.
- Validasi request dan penanganan formulir.
- SQLite, MariaDB, dan migration.
- Generator proyek yang dapat memakai Goku secara lokal/offline.
- Panduan reverse proxy nginx atau Apache untuk Laragon.
- Tes lintas Windows, Linux, dan macOS.

## Melaporkan masalah

Sertakan versi Go, sistem operasi, langkah reproduksi, hasil yang diharapkan, dan hasil yang terjadi. Jangan mempublikasikan token, kata sandi, atau data pribadi dalam issue.
