# SIMAK-GO (Sistem Informasi Manajemen Akademik Kampus)

Backend API untuk sistem akademik kampus yang dibangun menggunakan Golang, Fiber, dan MongoDB dengan konsep **Clean Architecture**.

## 🚀 Tech Stack
- **Language:** Golang
- **Framework:** Fiber v2
- **Database:** MongoDB
- **Driver:** Official MongoDB Driver
- **Authentication:** JWT (JSON Web Token)
- **Architecture:** Clean Architecture

## 📂 Struktur Project
Project ini mengikuti aturan Clean Architecture, memisahkan dependency secara tegas.

```
app/
├── models/          # Entity / Domain Models
├── repository/      # Database Access Layer (Interface & Implementation)
├── service/         # Business Logic Layer
├── handler/         # HTTP Handler Layer (Controller)
├── middleware/      # Middleware (Auth, Role, Logging)
├── route/           # Route Definitions
├── helper/          # Utility functions (Response, Security)
├── config/          # Environment Configuration
├── database/        # Database Connection
main.go              # Entry point
```

## ⚙️ Setup & Installation

1. **Clone Repository**
   ```bash
   git clone https://github.com/yourusername/simak-go.git
   cd simak-go
   ```

2. **Install Dependencies**
   ```bash
   go mod tidy
   ```

3. **Setup Environment**
   Buat file `.env` di root folder:
   ```env
   APP_PORT=3000
   MONGO_URI=mongodb://localhost:27017
   MONGO_DB=simak_go
   JWT_SECRET=rahasia_negara_api
   ```

4. **Run Application**
   ```bash
   go run main.go
   ```

## 🧪 API Endpoints (Contoh)

### Auth
- `POST /api/auth/register` - Register user baru
- `POST /api/auth/login` - Login dan dapatkan token

### User (Protected)
- `GET /api/users/profile` - Get current user profile (Butuh Token)

### Admin (Role: admin)
- `GET /api/admin/dashboard` - Admin only page

## 🧱 Arsitektur
Alur data mengalir dari luar ke dalam:
**HTTP Request -> Handler -> Service -> Repository -> Database**

- **Handler**: Menerima request, validasi input dasar, memanggil Service.
- **Service**: Menjalankan business logic, memanggil Repository.
- **Repository**: Berinteraksi langsung dengan Database.
- **Models**: Defines struktur data.
