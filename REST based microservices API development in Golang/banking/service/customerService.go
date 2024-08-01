package service

import "banking/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return []domain.Customer{}, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
