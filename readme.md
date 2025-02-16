# Goph Password Keeper

![example!](./docs/example.gif "example")

## Архитектура проекта
![C12!](./docs/c12.png "C12")
![C3!](./docs/c3.png "C3")

## Структура проекта

    .
    ├── cmd                 
    │   ├── client           # Приложение клиента 
    │   └── server           # Приложение сервера
    ├── internal             # Внутренний код
    │   ├── client           
    │   │   ├── app         
    │   │   └── infra
    │   ├── common        
    │   │   └── dto          # Domain слой
    │   └── server           
    │       ├── app        
    │       ├── infra         
    │       └── test
    └── ...

## Запуск проекта

Создать в корне .env.client

```
DATABASE_FILE=
GRPC_ADDRESS=
```

Создать в корне .env.server

```
DATABASE_DSN=
GRPC_ADDRESS=
JWT_SECRET=
```

Подготовка окружения и запуск

```bash
make up
```
```bash
make protoc
```
