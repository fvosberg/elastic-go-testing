package main

import (
	"github.com/fvosberg/elastic-go-testing/customer"
	"testing"
)

func TestRegistration(t *testing.T) {
	testCustomer := customer.Customer{Email: "testing@test.de"}
	testCustomer.Create()
	client := customer.CustomerClient()
	client.Flush()
	_, err := customer.FindByEmail("testing@test.de")
	if err != nil {
		t.Logf("Error occured: %+v\n", err)
		t.Fail()
	} else {
		t.Log("Found customer testing@test.de")
	}
}
