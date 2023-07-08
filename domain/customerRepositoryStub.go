package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"01", "Ed", "Aru", "88900000", "29/04/1995", "1"},
		{"02", "Edm", "Aru", "88900000", "29/04/1995", "1"},
	}
	return CustomerRepositoryStub{customers}
}
