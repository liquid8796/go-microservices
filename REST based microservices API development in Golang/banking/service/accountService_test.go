package service

import (
	realdomain "banking/domain"
	"banking/dto"
	"banking/mocks/domain"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/liquid8796/banking-lib/errs"
)

var mockRepo *domain.MockAccountRepository
var service AccountService

func setup(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo = domain.NewMockAccountRepository(ctrl)
	service = NewAccountService(mockRepo)
}

func Test_should_return_a_validation_error_response_when_the_request_is_not_validated(t *testing.T) {
	// Arrange
	request := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      0,
	}

	service := NewAccountService(nil)

	// Act
	_, appError := service.NewAccount(request)
	// Assert
	if appError == nil {
		t.Error("failed while testing the new account validation")
	}
}

func Test_should_return_an_error_from_server_side_if_the_new_account_cannot_be_created(t *testing.T) {
	// Arrange
	setup(t)

	req := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      6000,
	}
	account := realdomain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	mockRepo.EXPECT().Save(account).Return(nil, errs.NewUnexpectedError("Unexpected database error"))

	// Act
	_, appError := service.NewAccount(req)

	// Assert
	if appError == nil {
		t.Error("Test failed while validating error for new account")
	}
}

func Test_should_return_new_account_response_when_a_new_account_is_saved_sucessfully(t *testing.T) {
	// Arrange
	setup(t)

	req := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      6000,
	}
	account := realdomain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: dbTSLayout,
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	accountWithId := account
	accountWithId.AccountId = "201"
	mockRepo.EXPECT().Save(account).Return(&accountWithId, nil)

	// Act
	newAccount, appError := service.NewAccount(req)

	// Assert
	if appError != nil {
		t.Error("Test failed while creating new account")
	}
	if newAccount.AccountId != accountWithId.AccountId {
		t.Error("Test failed while matching new account")
	}
}
