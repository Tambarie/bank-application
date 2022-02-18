package service

import "bank-application/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

type DefaultCustomerSrvice struct {
	repo domain.CustomerRepository
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerSrvice {
	return DefaultCustomerSrvice{
		repo: repository,
	}
}
