package app

import (
	"banking/dto"
	"banking/mocks/service"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func Test_should_return_customers_with_status_code_200(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mockService := service.NewMockCustomerService(ctrl)
	dummyCustomers := []dto.CustomerResponse{
		{Id: "1001", Name: "Nam", City: "Thu Dau Mot", Zipcode: "70000", DateOfBirth: "1996-07-08", Status: "1"},
		{Id: "1002", Name: "Hung", City: "Ho Chi Minh", Zipcode: "100000", DateOfBirth: "1997-01-01", Status: "1"},
	}
	mockService.EXPECT().GetAllCustomer("").Return(dummyCustomers, nil)
	ch := CustomerHandlers{mockService}

	router := mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers)

	http.NewRequest(http.MethodGet, "/customers", nil)

}
