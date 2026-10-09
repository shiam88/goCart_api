package payment

import (
	"context"
	"goCartAPI/domain/order"
	"goCartAPI/domain/shared"
)

func (s *service) Refund(ctx context.Context, transactionID shared.ID) error {
	p, err := s.payments.FindByID(ctx, transactionID)
	if err != nil {
		return err
	}

	g, err := s.gateways.Resolve(p.Provider())
	if err != nil {
		return err
	}

	if err := g.Refund(ctx, p.ProviderReference(), p.Amount()); err != nil {
		return err
	}

	return s.tx.WithinTransaction(ctx, func(txctx context.Context) error {
		p, err = s.payments.FindByID(txctx, transactionID)
		if err != nil {
			return err
		}
		o, err := s.orders.FindByID(txctx, p.OrderID())
		if err != nil {
			return err
		}
		if err := p.Refund(); err != nil {
			return err
		}
		if o.Status() != order.StatusRefunded {
			if err := o.TransitionTo(order.StatusRefunded); err != nil {
				return err
			}
		}
		if err := s.payments.Save(txctx, p); err != nil {
			return err
		}
		return s.orders.Save(txctx, o)
	})
}
