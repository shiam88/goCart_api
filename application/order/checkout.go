package order

import (
	"context"
	"fmt"
	"strings"

	"goCartAPI/domain/cart"
	"goCartAPI/domain/order"
	"goCartAPI/domain/payment"
	"goCartAPI/domain/shared"
)

func (s *service) Checkout(ctx context.Context, in CheckoutInput) (*CheckoutResult, error) {
	if !in.PaymentProvider.IsValid() {
		return nil, payment.ErrInvalidProvider
	}

	itemsInput := in.Items //when the source is not "CART"

	if strings.EqualFold(in.Source, "CART") {
		c, err := s.carts.FindByUserID(ctx, in.UserID)
		if err != nil {
			return nil, err
		}
		if c.IsEmpty() {
			return nil, cart.ErrEmptyCart
		}

		itemsInput = make([]CheckoutItemInput, 0, len(c.Items()))
		for _, ci := range c.Items() {
			itemsInput = append(itemsInput, CheckoutItemInput{ProductID: ci.ProductID(), Quantity: ci.Quantity()})
		}
	}

	if len(itemsInput) == 0 {
		return nil, order.ErrEmptyOrder
	}

	var createdOrder *order.Order
	var createdPayment *payment.Transaction

	err := s.tx.WithinTransaction(ctx, func(txctx context.Context) error {

		orderItems := make([]order.Item, 0, len(itemsInput))

		for _, inputItem := range itemsInput {

			if inputItem.Quantity <= 0 {
				return order.ErrInvalidQuantity
			}

			p, err := s.products.FindByID(txctx, inputItem.ProductID)
			if err != nil {
				return err
			}
			if !p.IsActive() {
				return fmt.Errorf("product %s is inactive", p.ID())
			}

			inv, err := s.inventory.FindByProductID(txctx, inputItem.ProductID) // p.ID())
			if err != nil {
				return err
			}

			if err := inv.Reserve(inputItem.Quantity); err != nil {
				return err
			}
			if err := s.inventory.Save(txctx, inv); err != nil {
				return err
			}

			oi, err := order.NewItem(p.ID(), p.Name(), p.Price(), inputItem.Quantity)
			if err != nil {
				return err
			}
			orderItems = append(orderItems, oi)
		}

		zero, err := shared.ZeroMoney(s.shippingFee.Currency())
		if err != nil {
			return err
		}

		o, err := order.NewOrder(in.UserID, orderItems, s.shippingFee, zero, in.ShippingAddress)
		if err != nil {
			return err
		}

		pt, err := payment.NewTransaction(o.ID(), in.UserID, o.Total(), in.PaymentProvider)
		if err != nil {
			return err
		}
		if err := s.orders.Save(txctx, o); err != nil {
			return err
		}
		if err := s.payments.Save(txctx, pt); err != nil {
			return err
		}

		createdOrder, createdPayment = o, pt
		return nil
	})

	if err != nil {
		return nil, err
	}

	if in.PaymentProvider == payment.ProviderCOD {
		if err := s.tx.WithinTransaction(ctx, func(txctx context.Context) error {

			if err := createdOrder.TransitionTo(order.StatusProcessing); err != nil {
				return err
			}

			for _, item := range createdOrder.Items() {
				inv, err := s.inventory.FindByProductID(txctx, item.ProductID())
				if err != nil {
					return err
				}
				if err := inv.CommitReservation(item.Quantity()); err != nil {
					return err
				}
				if err := s.inventory.Save(txctx, inv); err != nil {
					return err
				}
			}

			return s.orders.Save(txctx, createdOrder)
		}); err != nil {
			return nil, err
		}

		s.clearCartAfterCheckout(ctx, in)

	}
	return &CheckoutResult{Order: createdOrder, Payment: createdPayment}, nil
}

// 	gateway, err := s.gateways.Resolve(in.PaymentProvider)
// 	if err != nil {
// 		return nil, err
// 	}
// 	reference, err := gateway.Initiate(ctx, createdOrder.ID(), createdOrder.Total())

// 	if err != nil {
// 		_ = s.tx.WithinTransaction(ctx, func(txctx context.Context) error {
// 			_ = createdPayment.MarkFailed()
// 			_ = createdOrder.Cancel()
// 			for _, item := range createdOrder.Items() {
// 				inv, e := s.inventory.FindByProductID(txctx, item.ProductID())
// 				if e != nil {
// 					return e
// 				}
// 				if e = inv.Release(item.Quantity()); e != nil {
// 					return e
// 				}
// 				if e = s.inventory.Save(txctx, inv); e != nil {
// 					return e
// 				}
// 			}
// 			if e := s.payments.Save(txctx, createdPayment); e != nil {
// 				return e
// 			}
// 			return s.orders.Save(txctx, createdOrder)
// 		})

// 		return nil, err
// 	}

// 	if err := s.tx.WithinTransaction(ctx, func(txctx context.Context) error {
// 		if err := createdPayment.StartProcessing(reference); err != nil {
// 			return err
// 		}
// 		if err := createdOrder.TransitionTo(order.StatusAwaitingPayment); err != nil {
// 			return err
// 		}
// 		if err := s.payments.Save(txctx, createdPayment); err != nil {
// 			return err
// 		}
// 		return s.orders.Save(txctx, createdOrder)
// 	}); err != nil {
// 		return nil, err
// 	}

// 	s.clearCartAfterCheckout(ctx, in)
//	return &CheckoutResult{Order: createdOrder, Payment: createdPayment, ProviderReference: reference}, nil

func (s *service) clearCartAfterCheckout(ctx context.Context, in CheckoutInput) {
	if !strings.EqualFold(in.Source, "CART") {
		return
	}
	if c, err := s.carts.FindByUserID(ctx, in.UserID); err == nil {
		c.Clear()
		_ = s.carts.Save(ctx, c)
	}
}
