package app

import (
	"banking/mocks/service"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func Test_should_return_customers_with_status_code_200(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mockService := service.NewMockCustomerService(ctrl)
	mockService.EXPECT().GetAllCustomer("").Return()
	ch := CustomerHandlers{mockService}

	router := mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers)

	// Act

	// Assert

}
