package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1233", "Jae", "Berlin", "1110011", "2000-20-10", "1"},
		{"1234", "Alex", "Berlin", "1110011", "2000-20-10", "1"},
	}
	return CustomerRepositoryStub{customers}
}
