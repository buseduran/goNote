# Node.js base image kullanıyoruz
FROM node:20-alpine

# Çalışma dizini ayarla
WORKDIR /app

# package.json ve package-lock.json dosyalarını kopyala ve bağımlılıkları yükle
COPY package.json package-lock.json ./
RUN npm install

# Kalan dosyaları kopyala
COPY . .

# React uygulamasını build et
RUN npm run build

# Nginx kullanarak servisi başlat
FROM nginx:alpine
COPY --from=0 /app/dist /usr/share/nginx/html

# Nginx'in çalışacağı port
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
