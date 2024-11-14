package factories

import (
	"ecommerce/helpers"
	"ecommerce/models/mysql"
)

func newUser(companyId uint64) *mysql.User {
	return &mysql.User{
		CompanyId: companyId,
		Name:      helpers.GenerateLetter(8),
		Email:     helpers.GenerateString(12) + "_user@example.com",
		CommonColumn: mysql.CommonColumn{
			CreatedAt: helpers.GetTimeNow(),
			UpdatedAt: helpers.GetTimeNow(),
		},
	}
}

func DefinitionUser(companyId uint64) *mysql.User {
	return newUser(companyId)
}

func SeedUser(number int, companyId uint64) []*mysql.User {
	data := make([]*mysql.User, number)

	for i := 0; i < number; i++ {
		data[i] = DefinitionUser(companyId)
	}

	return data
}
