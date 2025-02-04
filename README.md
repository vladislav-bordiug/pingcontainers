# Функционал

Сервис Pinger получает все контейнеры, пингует их и отправляет записи в Backend в виде IP, время пинга и дата пинга (в случае если не было успешных пингов значение null). Backend содержит два REST маршрута для добавления/изменения Pingerом записи в бд и для получения всей информации о pingах. Frontend (React/TS/Bootstrap) представляет собой таблицу, которая обновляется с определенным интервалом. База данных PostgreSQL.

Периодичность пинга Pingerом можно менять как PINGER_INTERVAL_SECONDS в [.env](/.env).

Периодичность обновления фронтенда можно менять как VITE_PINGER_INTERVAL_SECONDS в [frontend/.env](frontend/.env).

![image](https://github.com/user-attachments/assets/e55012ca-ce0b-443c-8cd5-8c60904a2b95)

# Как запустить

Нужно выполнить в терминале команды, если на машине установлены Docker и Docker Compose:

```git clone https://github.com/vladislav-bordiug/pingcontainers```

```cd pingcontainers```

```docker-compose up --build```

Frontend доступен будет на http://localhost:3000/ 
