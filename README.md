# delos-intern

## Frameworks
- HTTP Web : echo
- ORM : Gorm
- Database : PostgreSQL
- Unit Test : testify

## How to run
```bash
# clone the project
$ git clone -b local https://github.com/ilhamfzz/delos-intern.git

# enter the project directory
$ cd delos-intern

# Isi file .env sesuai dengan konfigurasi database
# buat 2 database, 1 untuk testing dan 1 untuk development

# download dependencies package
$ go mod download

# run unit test
$ go test -v .\testing\

# run the project
$ go run main.go
```

## API Documentation
- Postman : [https://documenter.getpostman.com/view/23908351/2s8YYJp2qj](https://documenter.getpostman.com/view/23908351/2s8YYJp2qj)

## Kendala
- Masih kurang memahami penggunaan docker
- Masih belajar mengenai unit test
- Masih berusaha menggunakan arsitektur yang lebih baik lagi