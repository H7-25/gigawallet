package main

import (
	"fmt"

	giga "github.com/dogecoinfoundation/gigawallet/pkg"
	"github.com/dogecoinfoundation/gigawallet/pkg/broker"
	"github.com/dogecoinfoundation/gigawallet/pkg/chaintracker"
	"github.com/dogecoinfoundation/gigawallet/pkg/conductor"
	"github.com/dogecoinfoundation/gigawallet/pkg/core"
	"github.com/dogecoinfoundation/gigawallet/pkg/dogecoin"
	"github.com/dogecoinfoundation/gigawallet/pkg/receivers"
	"github.com/dogecoinfoundation/gigawallet/pkg/store"
	"github.com/dogecoinfoundation/gigawallet/pkg/webapi"
)

func Server(conf giga.Config) {

	rpc, err := dogecoin.NewL1Libdogecoin(conf, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(rpc.MakeAddress())

	c := conductor.NewConductor(
		conductor.HookSignals(),
		conductor.Noisy(),
	)

	// Start the MessageBus Service
	bus := giga.NewMessageBus()
	c.Service("MessageBus", bus)

	// Set up all configured receivers
	receivers.SetUpReceivers(c, bus, conf)

	// Set up the L1 interface to Core
	l1_core, err := core.NewDogecoinCoreRPC(conf)
	if err != nil {
		panic(err)
	}
	l1, err := dogecoin.NewL1Libdogecoin(conf, l1_core)
	if err != nil {
		panic(err)
	}

	// Setup a Store, SQLite for now
	store, err := store.NewSQLiteStore(conf.Store.DBFile)
	if err != nil {
		panic(err)
	}
	defer store.Close()

	// Start the Chain Tracker
	chaser, follower, err := chaintracker.StartChainTracker(c, conf, l1, store)
	if err != nil {
		panic(err)
	}

	// Start the PaymentBroker service (deprecated)
	pb := broker.NewPaymentBroker(conf, store)
	c.Service("Payment Broker", pb)

	// Start the Core listener service (ZMQ)
	corez, err := core.NewCoreZMQReceiver(bus, conf)
	if err != nil {
		panic(err)
	}
	corez.Subscribe(chaser)
	c.Service("ZMQ Listener", corez)

	api := giga.NewAPI(store, l1, bus, follower)

	// Start the Payment API
	p, err := webapi.NewWebAPI(conf, api)
	if err != nil {
		panic(err)
	}
	c.Service("Payment API", p)

	<-c.Start()
}