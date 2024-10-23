
# CRUD gRPC Server

Этот проект представляет собой gRPC сервер, выполняющий CRUD операции (создание, чтение, обновление, удаление), а также консольную утилиту для клиента, позволяющую взаимодействовать с сервером. Сервер реализует методы, описанные в `.proto` файле, и использует PostgreSQL в качестве базы данных.

## Структура проекта

```plaintext
.
├── api
│   └── proto
│       └── crud.proto                 # Описание gRPC сервера и сообщений
├── cmd
│   ├── client
│   │   └── main.go                    # Точка входа клиента
│   └── server
│       └── main.go                    # Точка входа сервера
├── config.yaml                        # Конфигурационный файл
├── go.mod                             # Go модуль
├── go.sum                             # Контрольные суммы зависимостей
├── internal
│   ├── api
│   │   └── proto
│   │       ├── crud_grpc.pb.go        # Сгенерированный код для gRPC сервера
│   │       └── crud.pb.go             # Сгенерированные сообщения
│   ├── cli
│   │   └── cli.go                     # Логика работы с командной строкой клиента
│   ├── convert
│   │   └── convert.go                 # Конвертация данных между proto и структурой базы данных
│   └── server
│       └── server.go                  # Реализация gRPC сервера
└── readme.md                          # Описание проекта
```

## Описание компонентов

### 1. gRPC API

gRPC интерфейс и сообщения описаны в файле [crud.proto](https://github.com/notblinkyet/crud_gRPC/blob/master/api/proto/crud.proto). Сгенерированные из этого файла `.pb.go` файлы находятся в `internal/api/proto` и содержат код для взаимодействия с gRPC сервером.

### 2. Сервер

Реализация gRPC сервера находится в [internal/server/server.go](https://github.com/notblinkyet/crud_gRPC/blob/master/internal/server/server.go). Сервер реализует методы, описанные в `proto` файле, и использует интерфейс `storage.Storage` для взаимодействия с базой данных. Реализация базы данных взята из предыдущего проекта - [Crud](https://github.com/notblinkyet/Crud).

Серверная точка входа расположена в [cmd/server/main.go](https://github.com/notblinkyet/crud_gRPC/blob/master/cmd/server/main.go), где загружается конфигурация, подключается база данных (в текущей реализации - PostgreSQL).

### 3. Клиент

Консольная утилита клиента реализована в [cmd/client/main.go](https://github.com/notblinkyet/crud_gRPC/blob/master/cmd/client/main.go). Клиент загружает конфигурацию, инициализирует gRPC соединение с сервером, и предоставляет интерфейс командной строки для выполнения CRUD операций.

Логика обработки команд и взаимодействия с сервером описана в [internal/cli/cli.go](https://github.com/notblinkyet/crud_gRPC/blob/master/internal/cli/cli.go).

### 4. Конвертация данных

Для преобразования данных между форматами, используемыми в gRPC и базой данных, используется модуль [internal/convert/convert.go](https://github.com/notblinkyet/crud_gRPC/blob/master/internal/convert/convert.go).

## Конфигурация

Конфигурация проекта задается в файле `config.yaml`. В нем указываются параметры подключения к базе данных и настройки сервера.

## Зависимости

- [gRPC](https://grpc.io/) — фреймворк для межпроцессного взаимодействия.
- [Go](https://golang.org/) — язык программирования.
- [PostgreSQL](https://www.postgresql.org/) — база данных для хранения данных.

## Авторы

- [notblinkyet](https://github.com/notblinkyet)