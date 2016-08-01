package customer

import (
	"fmt"
	elastic "gopkg.in/olivere/elastic.v3"
	"log"
	"os"
)

type Customer struct {
	Email string `json:"email"`
}

func FindByEmail(email string) (customer Customer, err error) {
	client := CustomerClient()
	searchResult, err := client.Search().
		Index("customer").
		Query(elastic.NewBoolQuery().Filter(elastic.NewTermQuery("email", email))).
		From(0).Size(1).
		Do()
	if err != nil {
		err = ElasticsearchConnectionError{err, 504}
		log.Fatal(err)
	} else if searchResult.TotalHits() < 1 {
		err = NotFoundError{fmt.Sprintf("Could not find the customer with the email address %s.\n", email), 404}
	}
	log.Println("HEY")
	log.Printf("Searchresult from elasticsearch: %+v\n", searchResult)

	return customer, err
}

func (customer *Customer) Create() error {
	log.Printf("Create customer in elasticsearch: %+v\n", customer)
	customerClient := CustomerClient()
	put, err := customerClient.Index().
		Index("customer").
		Type("customer").
		BodyJson(customer).
		Do()

	if err == nil {
		log.Printf("Created customer in elasticsearch: %+v, with reponse: %+v.\n", customer, put)
	} else {
		log.Printf("Could not create customer %+v in elasticsearch: %s\n", customer, err)
	}
	return nil
}

func CustomerClient() *elastic.Client {
	elasticsearchURL := "http://elasticsearch:9200"

	log.Printf("Connecting to elasticsearch on %s.\n", elasticsearchURL)
	client, err := elastic.NewClient(
		elastic.SetURL(elasticsearchURL),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))

	if err != nil {
		log.Fatal(err)
	}
	EnsureCustomerIndex(client)
	return client
}

func EnsureCustomerIndex(elasticClient *elastic.Client) {
	exists, err := elasticClient.IndexExists("customer").Do()
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		log.Println("The elasticsearch index customer doesn't exist yet. Trying to create it")
		createElasticsearchIndexForCustomer(elasticClient)
	}
}

func createElasticsearchIndexForCustomer(elasticClient *elastic.Client) error {
	elasticsearchCustomersMapping := `{
		"settings": {
			"number_of_shards": 1,
			"number_of_replicas": 0,
			"analysis": {
				"analyzer": {
					"email": {
						"type": "custom",
						"tokenizer": "keyword",
						"filter": "lowercase"
					}
				}
			}
		},
		"mappings": {
			"customer": {
				"properties": {
					"username": {
						"type": "string"
					},
					"email": {
						"type": "string",
						"analyzer": "email"
					},
					"is_activated": {
						"type": "boolean"
					}
				}
			}
		}
	}`

	createdIndex, err := elasticClient.CreateIndex("customer").BodyString(elasticsearchCustomersMapping).Do()
	if err != nil {
		log.Fatal(err)
		return err
	}
	if createdIndex.Acknowledged {
		log.Println("Successfully created elasticsearch index customer")
	} else {
		log.Fatal("Could not create index customer")
		return err
	}
	return nil
}
