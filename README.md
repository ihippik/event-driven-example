# Microservices: Event-Driven communication

### Зависимости
Для демонстрации межсервисного взаимодействия нам потребуется NATS-сервер, 
в качестве событийной шины. Кстати написанный тоже на языке Golang.
https://nats.io/

Для быстрого старта нам потребуется запустить его в Docker-контейнере:

`docker run --name nats -d -p 4222:4222 nats`

https://www.docker.com/

### Запуск
Запускаем два сервиса, которые должны будут общаться между собой.

Заходим в папку user_service и запускаем в ней 
наш web-сервис `users` командой `go run main.go`

Далее в другом окне терминала заходим в папку audit_service и запускаем 
в ней наш второй `audit` сервис командой `go run main.go`

Далее делам POST-запрос в наш users-service:

`curl -d '{"name": "Ivan"}' -X POST http://localhost:8080/users`

_после окончания ваших экспериментов не забудьте удалить запущенный Docker-контейнер с NATS._

`docker rm -f nats` 



