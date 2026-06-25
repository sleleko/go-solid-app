package main

import (
	"context"
	"database/sql"
	"log"

	"go-order-solid/internal/notifier"
	"go-order-solid/internal/repository"
	"go-order-solid/internal/service"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "orders.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewSQLiteRepo(db)
	if err := repo.Init(ctx); err != nil {
		log.Fatal(err)
	}

	// Один и тот же OrderService работает с EmailSender.
	emailService := service.NewOrderService(repo, notifier.NewEmailSender())
	if err := emailService.CreateOrder(ctx, "Иван", []string{"apple", "banana"}, 10.5); err != nil {
		log.Fatal(err)
	}

	// Тот же OrderService без изменений работает с SMSSender.
	smsService := service.NewOrderService(repo, notifier.NewSMSSender())
	if err := smsService.CreateOrder(ctx, "Мария", []string{"orange", "milk"}, 22.3); err != nil {
		log.Fatal(err)
	}
}
