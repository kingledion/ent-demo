package service

import (
	"context"
	"fmt"

	"github.com/kingledion/ent-demo/internal/ent"
)

type Service struct {
	repo *ent.Client
}

func (s Service) AddOrder(ctx context.Context, order internal.Order) err {

	mrch := order.Merchant

	// create Order
	usr, err := s.repo.User.
		Create().
		SetUUID(order.User.UUID).
		SetFirstname(order.User.Firstname).
		SetLastname(order.User.Lastname).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating user: %v", err)
	}

	// create Merchant
	mrch, err := s.repo.Merchant.
		Create().
		SetUUID(order.Merchant.UUID).
		SetDBA(order.Merchant.DBA).
		AddMOrder(&user).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating merchant: %v", err)
	}
}
