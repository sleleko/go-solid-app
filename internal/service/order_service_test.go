package service_test

import (
	"context"
	"errors"
	"testing"

	"go-order-solid/internal/domain"
	"go-order-solid/internal/service"
)

type mockWriter struct {
	savedOrder domain.Order
	called     bool
	err        error
}

func (m *mockWriter) SaveOrder(_ context.Context, order domain.Order) error {
	m.called = true
	m.savedOrder = order
	return m.err
}

type mockNotifier struct {
	message domain.Message
	called  bool
	err     error
}

func (m *mockNotifier) Send(_ context.Context, message domain.Message) error {
	m.called = true
	m.message = message
	return m.err
}

func TestCreateOrderSuccess(t *testing.T) {
	writer := &mockWriter{}
	sender := &mockNotifier{}

	orderService := service.NewOrderService(writer, sender)
	err := orderService.CreateOrder(context.Background(), "Иван", []string{"apple", "banana"}, 10.5)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if !writer.called {
		t.Fatal("expected repository writer to be called")
	}

	if writer.savedOrder.Customer != "Иван" {
		t.Fatalf("expected customer Иван, got %s", writer.savedOrder.Customer)
	}

	if writer.savedOrder.Status != "pending" {
		t.Fatalf("expected status pending, got %s", writer.savedOrder.Status)
	}

	if !sender.called {
		t.Fatal("expected notifier to be called")
	}

	if sender.message.To != "Иван" {
		t.Fatalf("expected message receiver Иван, got %s", sender.message.To)
	}
}

func TestCreateOrderReturnsRepositoryError(t *testing.T) {
	writer := &mockWriter{err: errors.New("db error")}
	sender := &mockNotifier{}

	orderService := service.NewOrderService(writer, sender)
	err := orderService.CreateOrder(context.Background(), "Иван", []string{"apple"}, 10.5)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if sender.called {
		t.Fatal("notifier should not be called if repository returned an error")
	}
}
