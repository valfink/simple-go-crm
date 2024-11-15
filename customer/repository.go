package customer

var allCustomers = PrepareMockData()

func ListAllCustomers() []Customer {
	customerList := make([]Customer, len(allCustomers))
	i := 0
	for _, c := range allCustomers {
		customerList[i] = c
		i++
	}

	return customerList
}
