package app

import (
	"banking/dto"
	"banking/errs"
	"banking/mocks/service"
	"net/http"
	"net/http/httptest"
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

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_status_code_500_with_error_message(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mockService := service.NewMockCustomerService(ctrl)
	mockService.EXPECT().GetAllCustomer("").Return(nil, errs.NewUnexpectedError("some database errors"))
	ch := CustomerHandlers{mockService}

	router := mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers)

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}
