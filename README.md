# Fibonacci
Fibonacci with REST and gRPC servers 

## Настройка
В фаиле конфигурации необходимо задать порты для REST и gRPC серверов,
а так же адрес memcache сервера

Фаил конфигурации (см. [configs](config/configs.toml)):

```
./config/configs.toml
```

## Сборка и запуск
Сервис можно собрать и запустить с помощью `make` (см. [Makefile](Makefile)).

Для сборки сервиса должен быть установлен компилятор golang.

Сборка сервиса:

```bash
make build
```
Для запуска должен быть запущен memcache сервер.

Команда запуска сервиса:

```bash
make start
```

## REST API
