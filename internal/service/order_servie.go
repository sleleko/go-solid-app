package service

import (
	"context"
	"fmt"

	"go-order-solid/internal/domain"
	"go-order-solid/internal/notifier"
	"go-order-solid/internal/repository"
)

type OrderService struct {
	repo     repository.RepositoryWriter
	notifier notifier.Notifier
}

func NewOrderService(repo repository.RepositoryWriter, n notifier.Notifier) *OrderService {
	return &OrderService{
		repo:     repo,
		notifier: n,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, customer string, products []string, total float64) error {
	order := domain.Order{
		Customer: customer,
		Products: products,
		Total:    total,
		Status:   "pending",
	}

	if err := s.repo.SaveOrder(ctx, order); err != nil {
		return fmt.Errorf("create order: %w", err)
	}

	message := domain.Message{
		To:   customer,
		Text: fmt.Sprintf("Ваш заказ на сумму %.2f создан и ожидает обработки", total),
	}

	if err := s.notifier.Send(ctx, message); err != nil {
		return fmt.Errorf("send notification: %w", err)
	}

	return nil
}
