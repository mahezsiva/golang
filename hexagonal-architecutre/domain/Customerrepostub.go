package domain

type Customerrepo struct {
	Customers []Customer
}

func (c Customerrepo) FindAll() ([]Customer, error) {
	return c.Customers, nil
}

func NewcustomerStub() Customerrepo {
	Customers := []Customer{
		{Name: "mahesh", Age: "12"},
	}
	return Customerrepo{Customers}

}
