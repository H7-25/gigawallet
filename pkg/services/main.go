package services

import (
	giga "github.com/dogecoinfoundation/gigawallet/pkg"
	"github.com/dogecoinfoundation/gigawallet/pkg/conductor"
)

func StartServices(cond *conductor.Conductor, bus giga.MessageBus, conf giga.Config, store giga.Store) {
	// BalanceKeeper sends "Balance Change" events.
	keeper := NewBalanceKeeper(store)
	cond.Service("NewBalanceKeeper", keeper)

	// InvoiceStamper sends "Invoice Paid" and "Invoice Partial Payment" events.
	stamper := NewInvoiceStamper()
	cond.Service("InvoiceStamper", stamper)

	// PayMaster sends "Payment Accepted" and "Payment Confirmed" events.
	master := NewPayMaster()
	cond.Service("PayMaster", master)
}