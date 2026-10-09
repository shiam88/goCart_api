package order

import (
	"context"
	"errors"

	"goCartAPI/domain/order"
	"goCartAPI/domain/payment"
	"goCartAPI/domain/shared"
)

var ErrPaidOrderNeedsRefund = errors.New("paid or in-progress payment must be refunded instead of directly cancelled")

func (s *service) Cancel(ctx context.Context, orderID, userID shared.ID) error {
	return s.tx.WithinTransaction(ctx, func(txctx context.Context) error {
		o, err := s.orders.FindByID(txctx, orderID)
		if err != nil {
			return err
		}
		if o.UserID() != userID {
			return ErrOrderForbidden
		}
		previousStatus := o.Status()
		p, err := s.payments.FindByOrderID(txctx, o.ID())
		if err != nil {
			return err
		}
		if p.Status() == payment.StatusPaid || p.Status() == payment.StatusProcessing {
			return ErrPaidOrderNeedsRefund
		}
		if err := o.Cancel(); err != nil {
			return err
		}
		if p.Status() == payment.StatusPending {
			if err := p.Cancel(); err != nil {
				return err
			}
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
			} else if previousStatus == order.StatusProcessing {
				if err := inv.AddStock(item.Quantity()); err != nil {
					return err
				}
			}
			if err := s.inventory.Save(txctx, inv); err != nil {
				return err
			}
		}
		if err := s.payments.Save(txctx, p); err != nil {
			return err
		}
		return s.orders.Save(txctx, o)
	})
}
