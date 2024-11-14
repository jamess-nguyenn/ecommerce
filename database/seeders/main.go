package main

import (
	"ecommerce/database/connection"
	"ecommerce/database/factories"
	"ecommerce/repositories/mysql"
	"fmt"
)

func main() {
	db, err := connection.ConnectMysql()
	if err != nil {
		fmt.Printf("Could not connect to the database: %v\n", err)
		return
	}

	defer func() {
		if err = db.Close(); err != nil {
			fmt.Printf("Error closing the database connections: %v\n", err)
		} else {
			fmt.Println("Closed the database connections")
		}
	}()

	// beginning seeding users
	companyId := uint64(100000)
	userNumber := 100
	users := factories.SeedUser(userNumber, companyId)

	userRepo := mysql.NewUserRepository(db)

	if err := userRepo.CreateMany(users); err != nil {
		fmt.Printf("Error seeding users: %+v\n", err)
	} else {
		fmt.Println("Seeded users successfully")
	}

	user1 := factories.DefinitionUser(companyId)
	outputUser1, _ := userRepo.Create(user1)
	userId1 := outputUser1.Id

	user2 := factories.DefinitionUser(companyId)
	outputUser2, _ := userRepo.Create(user2)
	userId2 := outputUser2.Id

	fmt.Println("User id:", userId1, "and", userId2)
	// end seeding users

	// beginning seeding products
	productNumber := 1000
	products := factories.SeedProduct(productNumber)

	productRepo := mysql.NewProductRepository(db)

	if err := productRepo.CreateMany(products); err != nil {
		fmt.Printf("Error seeding products: %+v\n", err)
	} else {
		fmt.Println("Seeded products successfully")
	}

	product1 := factories.DefinitionProduct()
	outputProduct1, _ := productRepo.Create(product1)
	productId1 := outputProduct1.Id

	product2 := factories.DefinitionProduct()
	outputProduct2, _ := productRepo.Create(product2)
	productId2 := outputProduct2.Id

	fmt.Println("Product id:", productId1, "and", productId2)
	// end seeding products

	//for i, user := range users {
	//	fmt.Printf("Item with index #%d: %+v\n\n", i, *user)
	//}

	fmt.Println("Seeding data ...")
}
