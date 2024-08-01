package service

import "banking/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}
