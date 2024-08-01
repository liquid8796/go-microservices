package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Nam", City: "Thu Dau Mot", Zipcode: "70000", DateOfBirth: "1996-07-08", Status: "1"},
		{Id: "1002", Name: "Hung", City: "Ho Chi Minh", Zipcode: "100000", DateOfBirth: "1997-01-01", Status: "1"},
	}

	return CustomerRepositoryStub{customers}
}
