# PakaiWA - WhatsApp RESTful API

PakaiWA adalah sebuah RESTful API untuk mengintegrasikan layanan WhatsApp menggunakan Golang dan library **Whatsmeow** (Tulir). API ini memungkinkan pengguna untuk mengirim dan menerima pesan WhatsApp dengan mudah melalui endpoint yang dapat diakses.

## Fitur

- Kirim pesan teks, gambar, dokumen, dan media lainnya.
- Terima pesan teks, gambar, dan dokumen.
- Mendukung webhook untuk notifikasi masuk.
- Mudah diintegrasikan dengan aplikasi lain menggunakan HTTP API.

## Teknologi

- **Golang** - Bahasa pemrograman utama untuk pengembangan.
- **Whatsmeow** - Library WhatsApp untuk Golang yang digunakan untuk komunikasi dengan WhatsApp.
- **Tulir** - Framework untuk menangani komunikasi dengan WhatsApp menggunakan Whatsmeow.
- **Gin/Gorilla Mux** - Framework HTTP untuk routing API.
- **Docker** - Untuk containerization aplikasi.
- **PostgreSQL** - Database untuk menyimpan data pengiriman pesan (opsional, tergantung pada kebutuhan).

## Instalasi

### Prasyarat

Sebelum memulai, pastikan Anda memiliki hal berikut:
- [Go](https://golang.org/dl/) versi 1.18 atau lebih tinggi
- [Docker](https://www.docker.com/get-started) (jika menggunakan container)
- [Whatsmeow Tulir](https://github.com/tulir/whatsmeow) sebagai dependensi utama

### Langkah-langkah Instalasi

1. **Clone repository**:

   ```bash
   git clone https://github.com/PakaiWA/PakaiWA.git
   cd PakaiWA
