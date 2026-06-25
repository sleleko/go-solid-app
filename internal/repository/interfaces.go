package repository

import (
	"context"

	"go-order-solid/internal/domain"
)

// RepositoryInitializer отвечает только за подготовку хранилища,
// например за создание таблиц.
type RepositoryInitializer interface {
	Init(ctx context.Context) error
}

// RepositoryWriter отвечает только за запись заказов.
// OrderService зависит именно от этого интерфейса, а не от SQLiteRepo.
type RepositoryWriter interface {
	SaveOrder(ctx context.Context, order domain.Order) error
}
