package payment

import (
	"context"

	"goCartAPI/domain/order"
	"goCartAPI/domain/payment"
	"goCartAPI/domain/shared"
)

func (s *service) Initiate(ctx context.Context, orderID shared.ID) (string, error) {
	p, err := s.payments.FindByOrderID(ctx, orderID)
	if err != nil {
		return "", err
	}

	o, err := s.orders.FindByID(ctx, orderID)
	if err != nil {
		return "", err
	}

	g, err := s.gateways.Resolve(p.Provider())
	if err != nil {
		return "", err
	}

	ref, err := g.Initiate(ctx, p.OrderID(), p.Amount())
	if err != nil {
		txErr := s.tx.WithinTransaction(ctx, func(txctx context.Context) error {
			if err := p.MarkFailed(); err != nil {
				return err
			}

			if err := o.Cancel(); err != nil {
				return err
			}

			for _, item := range o.Items() {
				inv, err := s.inventory.FindByProductID(txctx, item.ProductID())
				if err != nil {
					return err
				}

				if inv.Reserved() >= item.Quantity() {
					if err := inv.Release(item.Quantity()); err != nil {
						return err
					}

					if err := s.inventory.Save(txctx, inv); err != nil {
						return err
					}
				}
			}

			if err := s.payments.Save(txctx, p); err != nil {
				return err
			}

			return s.orders.Save(txctx, o)
		})

		if txErr != nil {
			return "", txErr
		}

		return "", err
	}

	if err := s.tx.WithinTransaction(ctx, func(txctx context.Context) error {
		if err := p.StartProcessing(ref); err != nil {
			return err
		}

		if err := o.TransitionTo(order.StatusAwaitingPayment); err != nil {
			return err
		}

		if err := s.payments.Save(txctx, p); err != nil {
			return err
		}

		return s.orders.Save(txctx, o)
	}); err != nil {
		return "", err
	}

	return ref, nil
}

func (s *service) Get(ctx context.Context, id shared.ID) (*payment.Transaction, error) {
	return s.payments.FindByID(ctx, id)
}
