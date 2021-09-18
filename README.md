# Fibonacci
Fibonacci with REST and gRPC servers
## Запуск в Docker
Сервис можно собрать и запустить с помощью `make` (см. [Makefile](Makefile)).


Первым делом соберём Docker-образ cервиса:
```bash
make docker-build
```
Далее скачиваем образ memcached сервера и запускаем его в контейнере
```bash
make docker-mc
```
Запускаем сервис fibonacci в отдельном контейнере:
```bash
make docker-run
```

## REST API
Чтобы получить все числа последовательности Фибоначчи с порядковыми номерами от *first* до *last*, необходимо отправить GET запрос вида:
```
http://server.address:port/fibonacci?x=first&y=last
```
В ответ придёт json. Пример ответа:
```json
{
  "sequence": [
    0,
    1,
    1
  ]
}
```

## gRPC
Proto фаил(см. [fibonacci.proto](pkg/api/fibonacci.proto)) находится по пути:
```
./pkg/api/fibonacci.proto
```

