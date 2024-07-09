## Langkah-Langkah Menjalankan Program
Kloning atau Buat Proyek

```bash
    mkdir your_project
    cd your_project
```

Inisialisasi Proyek Go

```bash
    go mod init your_project
```

Buat File yang Diperlukan

Buat file main.go, go.mod, dan .env sesuai dengan yang diperlukan.

Instal Dependensi

```bash
    go mod tidy  
```

Atur file env

```bash
    export DB_CONN_STRING="user=youruser dbname=yourdb sslmode=disable"
    export PORT=8080
```

Jalankan program

```bash
    go run main.go
```

Akses API

```bash
    curl http://localhost:8080/summary
```