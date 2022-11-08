package service

import (
	"Order-Service/internal/data/pg"
	address "Order-Service/internal/service/handlers/address"
	delivery "Order-Service/internal/service/handlers/delivery"
	order "Order-Service/internal/service/handlers/order"
	orderItem "Order-Service/internal/service/handlers/order_item"
	status "Order-Service/internal/service/handlers/status"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"

	"Order-Service/internal/service/helpers"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()
	log := s.log.WithFields(map[string]interface{}{
		"service": "Order-Service",
	})

	r.Use(
		ape.RecoverMiddleware(log),
		ape.LoganMiddleware(log),
		ape.CtxMiddleware(
			helpers.CtxLog(log),
			helpers.CtxStatusesQ(pg.NewStatusesQ(s.db)),
			helpers.CtxOrdersQ(pg.NewOrdersQ(s.db)),
			helpers.CtxOrderItemsQ(pg.NewOrderItemsQ(s.db)),
			helpers.CtxAddressesQ(pg.NewAddressesQ(s.db)),
			helpers.CtxDeliveriesQ(pg.NewDeliveriesQ(s.db)),
		),
	)
	r.Route("/integrations/Order-Service", func(r chi.Router) {
		r.Route("/statuses", func(r chi.Router) {
			r.Post("/", status.CreateStatus)
			r.Get("/", status.GetStatusList)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", status.GetStatus)
				r.Put("/", status.UpdateStatus)
				r.Delete("/", status.DeleteStatus)
			})
		})
		r.Route("/orders", func(r chi.Router) {
			r.Post("/", order.CreateOrder)
			r.Get("/", order.GetOrderList)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", order.GetOrder)
				r.Put("/", order.UpdateOrder)
				r.Delete("/", order.DeleteOrder)
			})
		})
		r.Route("/order_items", func(r chi.Router) {
			r.Post("/", orderItem.CreateOrderItem)
			r.Get("/", orderItem.GetOrderItemList)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", orderItem.GetOrderItem)
				r.Put("/", orderItem.UpdateOrderItem)
				r.Delete("/", orderItem.DeleteOrderItem)
			})
		})
		r.Route("/addresses", func(r chi.Router) {
			r.Post("/", address.CreateAddress)
			r.Get("/", address.GetAddressList)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", address.GetAddress)
				r.Put("/", address.UpdateAddress)
				r.Delete("/", address.DeleteAddress)
			})
		})
		r.Route("/deliveries", func(r chi.Router) {
			r.Post("/", delivery.CreateDelivery)
			r.Get("/", delivery.GetDeliveryList)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", delivery.GetDelivery)
				r.Put("/", delivery.UpdateDelivery)
				r.Delete("/", delivery.DeleteDelivery)
			})
		})
	})

	return r
}
