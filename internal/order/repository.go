package order

import (
	"context"
	"encoding/json"
)

type (
	Repository struct {
		Saver interface {
			Save(context.Context, interface{}) error
		}
		Finder interface {
			Find(context.Context, interface{}) ([]byte, error)
		}
	}
)

func (repo Repository) Save(ctx context.Context, newOrder Order) error {
	return repo.Saver.Save(ctx, newOrder)
}

func (repo Repository) Find(ctx context.Context, criteria interface{}) ([]Order, error) {
	rawOrders, err := repo.Finder.Find(ctx, criteria)
	if err != nil {
		return nil, err
	}

	orders := make([]Order, 0)
	err = json.Unmarshal(rawOrders, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
