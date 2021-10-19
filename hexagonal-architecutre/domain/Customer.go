package domain


type Customer struct{
	Name string
	Age string
}

type CustomeService interface{
	FindAll()([]Customer,error)
}