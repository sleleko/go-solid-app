# Go Order SOLID

Рефакторинг простого приложения заказов с применением SOLID.

## Что реализовано

- `OrderService` отвечает за бизнес-логику создания заказа.
- `SQLiteRepo` отвечает за работу с SQLite.
- `EmailSender` и `SMSSender` отвечают за отправку уведомлений.
- `OrderService` зависит от интерфейсов `RepositoryWriter` и `Notifier`, а не от конкретных реализаций.
- Интерфейсы для БД разделены:
  - `RepositoryInitializer` — инициализация хранилища;
  - `RepositoryWriter` — запись заказов.
- Добавлены тесты с моками без реальной базы данных.

## Структура проекта

```text
cmd/orders/main.go
internal/domain/order.go
internal/domain/message.go
internal/repository/interfaces.go
internal/repository/sqlite.go
internal/notifier/notifier.go
internal/notifier/email.go
internal/notifier/sms.go
internal/service/order_service.go
internal/service/order_service_test.go
```

## Запуск

```bash
go mod tidy
go run ./cmd/orders
```

Ожидаемый вывод:

```text
Email-уведомление отправлено клиенту Иван: Ваш заказ на сумму 10.50 создан и ожидает обработки
SMS-уведомление отправлено клиенту Мария: Ваш заказ на сумму 22.30 создан и ожидает обработки
```

## Тесты

```bash
go test ./...
```