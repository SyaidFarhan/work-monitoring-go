# Go Work Monitoring API

## Stack

- Go
- Gin
- GORM
- SQLite
- JWT akan ditambahkan nanti
- bcrypt akan ditambahkan nanti

## Progress

- [x] Membuat Go module
- [x] Install Gin
- [x] Install GORM
- [x] Install SQLite driver
- [x] Membuat server Gin
- [x] Membuat endpoint `GET /`
- [x] Membuat endpoint `GET /health`
- [x] Membuat model `User`
- [x] Menjalankan GORM `AutoMigrate`
- [x] Memindahkan model `User` ke `internal/models`
- [x] Membuat `UserRepository` interface
- [x] Membuat implementasi repository dengan GORM
- [x] Memahami dasar `*`, `&`, pointer, dan `models.User`
- [x] Memperbaiki JSON tag dan unique index pada model `User`
- [x] Membuat `FindByID` repository menjadi public
- [x] Menangani error dari `router.Run`
- [x] Menjalankan `gofmt` dan `go test ./...`
- [ ] Membuat `POST /users`
- [ ] Membuat service layer untuk user
- [ ] Membuat handler layer untuk user
- [ ] Mendaftarkan route user secara terpisah
- [ ] Register dan hashing password
- [ ] Login dan JWT
- [ ] Authentication middleware
- [ ] Authorization berdasarkan role
- [ ] Project CRUD
- [ ] Task CRUD
- [ ] Dependency task
- [ ] Optimistic locking
- [ ] Transaction
- [ ] Test concurrency

## Arsitektur Layer

Struktur yang sedang digunakan:

```text
Request
  ↓
Router
  ↓
Handler
  ↓
Service
  ↓
Repository
  ↓
GORM / Database
```

Repository yang sudah dibuat:

- File: `internal/repositories/user_repository.go`
- Interface: `UserRepository`
- Method: `Create`, `FindByEmail`, `FindByID`
- Database query menggunakan `context.Context` dan GORM `WithContext`

Model yang sudah diperbaiki:

- `ID` memiliki tag `json:"id"` dan `gorm:"primaryKey"`
- `Email` memiliki tag `gorm:"uniqueIndex"`
- Repository dapat menerima dan mengembalikan pointer `*models.User`

## Cara Melanjutkan

Langkah berikutnya adalah membuat **service layer** dan menghubungkan service
dengan `UserRepository`. Setelah itu baru membuat handler `POST /users`
menggunakan `c.ShouldBindJSON`, validasi input, dan service.

Validasi terakhir:

```bash
gofmt -w main.go internal/models/user.go internal/repositories/user_repository.go
go test ./...
```

Hasil terakhir: berhasil, semua package dapat di-compile.

Berikan kode sedikit demi sedikit agar kode dapat diketik dan dipahami sendiri.
Lanjutkan ke bagian berikutnya hanya setelah pengguna menulis `next`.
