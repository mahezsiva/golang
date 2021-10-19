package service

import "working/domain"

type DefaultService interface {
	GetAll() ([]domain.Customer, error)
}

type CustomerServicestub struct {
	repo domain.Customerrepo
}

func (c CustomerServicestub) DefaultService([]domain.Customer, error) {
	c.repo.FindAll()
}

func Newservice(Repository domain.Customerrepo) CustomerServicestub {
	return CustomerServicestub{Repository}
}
