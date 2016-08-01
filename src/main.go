package main

import (
	"github.com/fvosberg/elastic-go-testing/customer"
	"log"
)

func main() {
	testCustomer := customer.Customer{Email: "test@hop.de"}
	testCustomer.Create()
	log.Println("Created testcustomer")
}
