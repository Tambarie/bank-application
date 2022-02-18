package domain

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}

type CustomerRepositoryStub struct {
	customers []Customer
}

func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	var customer = []Customer{
		{Id: "1001", Name: "Isaac", City: "Lagos", Zipcode: "122321", DateOfBirth: "12-12-12", Status: "1"},
		{Id: "1002", Name: "Sam", City: "Ph", Zipcode: "231223", DateOfBirth: "14-12-1", Status: "2"},
	}
	return CustomerRepositoryStub{customers: customer}
}
