package helpers

import (
	"context"
	"order-service/internal/data"

	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	statusesQCtxKey
	ordersQCtxKey
	orderItemsQCtxKey
	addressesQCtxKey
	deliveriesQCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxStatusesQ(entry data.StatusesQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, statusesQCtxKey, entry)
	}
}

func StatusesQ(r *http.Request) data.StatusesQ {
	return r.Context().Value(statusesQCtxKey).(data.StatusesQ).New()
}

func CtxOrdersQ(entry data.OrdersQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ordersQCtxKey, entry)
	}
}

func OrdersQ(r *http.Request) data.OrdersQ {
	return r.Context().Value(ordersQCtxKey).(data.OrdersQ).New()
}

func CtxOrderItemsQ(entry data.OrderItemsQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, orderItemsQCtxKey, entry)
	}
}

func OrderItemsQ(r *http.Request) data.OrderItemsQ {
	return r.Context().Value(orderItemsQCtxKey).(data.OrderItemsQ).New()
}

func CtxAddressesQ(entry data.AddressesQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, addressesQCtxKey, entry)
	}
}

func AddressesQ(r *http.Request) data.AddressesQ {
	return r.Context().Value(addressesQCtxKey).(data.AddressesQ).New()
}

func CtxDeliveriesQ(entry data.DeliveriesQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, deliveriesQCtxKey, entry)
	}
}

func DeliveriesQ(r *http.Request) data.DeliveriesQ {
	return r.Context().Value(deliveriesQCtxKey).(data.DeliveriesQ).New()
}
