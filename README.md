# Go Deepseek AI Chat API (Clean Architecture)

Aplikasi ini adalah contoh implementasi Clean Architecture pada layanan chat dengan Deepseek AI berbasis Go. Mendukung dua mode interaksi:

1. **HTTP API** — Endpoint `/chat` untuk menerima permintaan chat dari aplikasi eksternal (misal: Postman).
2. **Command Line Interface (CLI)** — Chat langsung melalui terminal.

---

## Fitur Utama
- Arsitektur clean: terpisah domain, usecase, delivery, infrastructure.
- Endpoint HTTP `/chat` untuk menerima permintaan chat berbasis JSON.
- CLI interaktif untuk mengirim pesan ke AI dan menerima balasan.
- Menggunakan environment variable untuk menyimpan API Key.

---

## Struktur Folder
```
initial-project/
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   ├── domain/
│   ├── usecase/
│   ├── delivery/
│   │   ├── http/
│   │   └── cli/
│   └── infrastructure/
├── .env
├── .env.example
└── README.md
```

### Penjelasan Layer
- **cmd/main.go**: Entry point aplikasi.
- **internal/domain**: Entity & interface (pure business logic).
- **internal/usecase**: Usecase utama (application logic).
- **internal/delivery/http**: Handler HTTP.
- **internal/delivery/cli**: Handler CLI.
- **internal/infrastructure**: Implementasi akses eksternal (API Deepseek).
- **internal/config**: Load environment dan ambil API key.

---

## Cara Menjalankan

1. **Clone repository** dan masuk ke folder project.
2. **Set API Key** Deepseek di file `.env`:
   ```env
   DEEPSEK_API_KEY=masukkan_api_key_anda
   ```
3. **Jalankan aplikasi:**
   ```bash
   go run cmd/main.go
   ```
4. **Akses endpoint** di `http://localhost:8080/chat` menggunakan Postman atau aplikasi lain.
5. **Gunakan CLI** di terminal untuk chat langsung dengan AI.

---

## Contoh Request HTTP (Postman)
- **URL:** `http://localhost:8080/chat`
- **Method:** `POST`
- **Headers:** `Content-Type: application/json`
- **Body:**
  ```json
  {
    "prompt": "Apa itu AI?"
  }
  ```

### Contoh Response
- **Sukses:**
  ```json
  {
    "response": "Jawaban dari AI di sini"
  }
  ```
- **Error:**
  ```json
  {
    "error": "pesan error"
  }
  ```

---

## Contoh Penggunaan CLI
Saat aplikasi berjalan, kamu bisa langsung mengetik pesan di terminal:

```
Chat dengan Deepseek AI. Ketik pesan dan tekan Enter. (Ctrl+C untuk keluar)
Anda: Halo AI!
AI: Halo! Ada yang bisa saya bantu?
```

---

## Catatan
- Pastikan API Key Deepseek kamu aktif dan benar.
- Jika ada error parsing response, aplikasi akan menampilkan pesan error untuk membantu debugging.
- Struktur clean architecture memudahkan pengembangan fitur baru dan testing.

---

## Lisensi
Proyek ini dibuat untuk latihan dan pembelajaran.
