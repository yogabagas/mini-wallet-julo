package controller

type AppController struct {
	CustomersController       interface{ CustomersController }
	OrderDetailsController    interface{ OrderDetailsController }
	EmployeesController       interface{ EmployeesController }
	ProductsController        interface{ ProductsController }
	ShippingMethodsController interface{ ShippingMethodsController }
	OrdersController          interface{ OrdersController }
	WalletsController         interface{ WalletsController }
}
