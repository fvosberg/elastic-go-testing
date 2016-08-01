package main

import (
	"github.com/fvosberg/elastic-go-testing/customer"
	"testing"
	"time"
)

func TestRegistration(t *testing.T) {
	testCustomer := customer.Customer{Email: "testing@test.de"}
	testCustomer.Create()
	time.Sleep(time.Second * 1)
	_, err := customer.FindByEmail("testing@test.de")
	if err != nil {
		t.Logf("Error occured: %+v\n", err)
		t.Fail()
	} else {
		t.Log("Found customer testing@test.de")
	}
}
