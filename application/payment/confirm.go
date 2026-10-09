package payment

import (
	"context"

	"goCartAPI/domain/order"
	"goCartAPI/domain/shared"
)

func (s *service) Confirm(ctx context.Context, orderID shared.ID, providerReference string, success bool) error {
	return s.tx.WithinTransaction(ctx, func(txctx context.Context) error {
		p, err := s.payments.FindByOrderID(txctx, orderID)
		if err != nil {
			return err
		}
		o, err := s.orders.FindByID(txctx, orderID)
		if err != nil {
			return err
		}
		if success {

			if err := p.MarkPaid(providerReference); err != nil {
				return err
			}
			if o.Status() == order.StatusAwaitingPayment {
				if err := o.TransitionTo(order.StatusPaid); err != nil {
					return err
				}
			}

			for _, item := range o.Items() {
				inv, err := s.inventory.FindByProductID(txctx, item.ProductID())
				if err != nil {
					return err
				}
				if inv.Reserved() >= item.Quantity() {
					if err := inv.CommitReservation(item.Quantity()); err != nil {
						return err
					}
					if err := s.inventory.Save(txctx, inv); err != nil {
						return err
					}
				}
			}

		} else {
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
		}

		if err := s.payments.Save(txctx, p); err != nil {
			return err
		}
		return s.orders.Save(txctx, o)
	})
}
