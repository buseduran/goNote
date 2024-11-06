# Go için base image
FROM golang:1.23.2-alpine

# Çalışma dizinini ayarla
WORKDIR /app

# go.mod ve go.sum dosyalarını kopyala
COPY go.mod go.sum ./

# Bağımlılıkları indir
RUN go mod download

# Kalan dosyaları kopyala
COPY . .

# Uygulamayı derle
RUN go build -o main .

# Uygulamanın çalışacağı port
EXPOSE 8080

# Uygulamayı başlat
CMD ["./main"]
