# frontend/Dockerfile
FROM node:18 as build

WORKDIR /app

# Копируем файлы зависимостей и устанавливаем их
COPY package.json package-lock.json* ./
RUN npm install

# Копируем исходный код и выполняем сборку
COPY . .
RUN npm run build

# Используем nginx для раздачи собранной статики
FROM nginx:alpine
COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]