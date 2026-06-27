package main

import (
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/auth"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/core"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/kitchen"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/menu"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/order"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/product"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/store"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/table"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/user"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/waiter"
)

func main() {
	routing, database := core.InitModule()
	
	auth.InitModule(routing, database)
	user.InitModule(routing, database)
	store.InitModule(routing, database)
	product.InitModule(routing, database)
	menu.InitModule(routing, database)
	kitchen.InitModule(routing, database)
	order.InitModule(routing, database)
	table.InitModule(routing, database)
	waiter.InitModule(routing, database)

	core.StartApplication()
}
