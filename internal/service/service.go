package service

import (
	"context"
	"fmt"
	"log"

	"github.com/kingledion/ent-demo/internal"
	"github.com/kingledion/ent-demo/internal/ent"
	"github.com/kingledion/ent-demo/internal/ent/user"
)

// Service is holder for functionality of the ent-demo
type Service struct {
	repo *ent.Client
}

// NewService is a factory function for the service
func New(repo *ent.Client) Service {
	return Service{repo: repo}
}

// AddOrder creates an order in response to an external call
func (s Service) AddOrder(ctx context.Context, order internal.Order) error {

	// create Merchant
	mrch, err := s.repo.Merchant.
		Create().
		SetUUID(order.Merchant.UUID).
		SetDba(order.Merchant.DBA).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating merchant: %v", err)
	}

	// create Order
	_, err = s.repo.User.
		Create().
		SetUUID(order.User.UUID).
		SetFirstname(order.User.Firstname).
		SetLastname(order.User.Lastname).
		AddOrder(mrch).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating user: %v", err)
	}

	return nil

}

// GetOrderedAtByUser returns all merchants a user ordered at
func (s Service) Recommend(ctx context.Context, userID string) ([]internal.Merchant, error) {

	// Get shopper with given id
	mrchs, err := s.repo.User.
		Query().
		Where(user.UUIDEQ(userID)). // get user
		QueryOrder().               // get all merchants ordered at
		QueryMOrder().              // get all shoppers who ordered at merchants
		QueryOrder().               // get all merchants those shoppers ordered at
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying merchants: %v", err)
	}
	log.Printf("merchant list returned: len(%d)", len(mrchs))

	var merchants = make([]internal.Merchant, len(mrchs))
	offset := 0
	for i, m := range mrchs {
		if m == nil {
			offset = offset + 1
		} else {
			merchants[i-offset] = internal.MerchantFromEnt(*m)
		}
	}

	return merchants, nil

}
