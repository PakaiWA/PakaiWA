# PakaiWA - WhatsApp RESTful API

PakaiWA adalah sebuah RESTful API untuk mengintegrasikan layanan WhatsApp menggunakan Golang dan library **Whatsmeow** (Tulir). API ini memungkinkan pengguna untuk mengirim dan menerima pesan WhatsApp dengan mudah melalui endpoint yang dapat diakses.

## Fitur

- Kirim pesan teks, gambar, dokumen, dan media lainnya.
- Terima pesan teks, gambar, dan dokumen.
- Mendukung webhook untuk notifikasi masuk.
- Mudah diintegrasikan dengan aplikasi lain menggunakan HTTP API.

## Tech Stack

- **Golang** - Bahasa pemrograman utama untuk pengembangan.
- **Whatsmeow** - Library WhatsApp untuk Golang yang digunakan untuk komunikasi dengan WhatsApp.
- **Gin/Gorilla Mux** - Framework HTTP untuk routing API.
- **Docker** - Untuk containerization aplikasi.
- **PostgreSQL** - Database untuk menyimpan data pengiriman pesan (opsional, tergantung pada kebutuhan).
