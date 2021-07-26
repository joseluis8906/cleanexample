package order

import "context"

type (
	Creator struct {
		Saver interface {
			Save(context.Context, Order) error
		}
	}
)

func (creator Creator) Exec(ctx context.Context, customer, address, courier string, amount int) error {
	// validar que todas las entradas cumplen
	// si todo aplica bien

	newOrder := Order{
		Customer: customer,
		Address:  address,
		Courier:  courier,
		Amount:   amount,
	}

	return creator.Saver.Save(ctx, newOrder)
}
