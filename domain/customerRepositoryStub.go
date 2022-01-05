package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "123", Name: "ashish", City: "ATL", Zipcode: "30111", DateOfBirth: "2000-01-01", Status: "1"},
		{Id: "124", Name: "rob", City: "ATL", Zipcode: "30111", DateOfBirth: "2000-01-01", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
