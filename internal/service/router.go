package service

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"gitlab.com/distributed_lab/ape"
	"order-service/internal/data/pg"
	address "order-service/internal/service/handlers/address"
	delivery "order-service/internal/service/handlers/delivery"
	order "order-service/internal/service/handlers/order"
	orderItem "order-service/internal/service/handlers/order_item"
	status "order-service/internal/service/handlers/status"
	"order-service/internal/service/middleware"

	"order-service/internal/service/helpers"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()
	log := s.log.WithFields(map[string]interface{}{
		"service": "order-service",
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
		middleware.BasicAuth(s.endpoints),
	)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Route("/integrations/order-service", func(r chi.Router) {
		r.Use(middleware.CheckManagerPosition())
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
			r.Post("/", order.CreateOrder(s.endpoints))
			r.Get("/", order.GetOrderList)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", order.GetOrder)
				r.Put("/", order.UpdateOrder(s.endpoints))
				r.Delete("/", order.DeleteOrder)
			})
		})
		r.Route("/order_items", func(r chi.Router) {
			r.Post("/", orderItem.CreateOrderItem(s.endpoints))
			r.Get("/", orderItem.GetOrderItemList)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", orderItem.GetOrderItem)
				r.Put("/", orderItem.UpdateOrderItem(s.endpoints))
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
			r.Post("/", delivery.CreateDelivery(s.endpoints))
			r.Get("/", delivery.GetDeliveryList)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", delivery.GetDelivery)
				r.Put("/", delivery.UpdateDelivery(s.endpoints))
				r.Delete("/", delivery.DeleteDelivery)
			})
		})
	})

	return r
}
