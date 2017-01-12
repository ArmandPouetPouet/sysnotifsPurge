package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
import (
	"fmt"

	credentials "github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/sha1sum/aws_signing_client"
	elastic "gopkg.in/olivere/elastic.v3"
)

//const url = "https://search-core-sysnotifs-prod-7tlhh72bx4aasz7we2d3cjvybi.eu-west-1.es.amazonaws.com"

var client *elastic.Client

func newElasticClient(creds *credentials.Credentials) (*elastic.Client, error) {
	signer := v4.NewSigner(creds)
	awsClient, err := aws_signing_client.New(signer, nil, "es", "eu-west-1")
	if err != nil {
		return nil, err
	}
	return elastic.NewClient(
		elastic.SetURL("https://search-core-sysnotifs-prod-7tlhh72bx4aasz7we2d3cjvybi.eu-west-1.es.amazonaws.com"),
		elastic.SetScheme("https"),
		elastic.SetHttpClient(awsClient),
		elastic.SetSniff(false), // See note below
	)
}
*/
func initConnection() {

}

func clusterStat() {
	res, err := http.Get("search-core-sysnotifs-beta-tsrx3dffhismokydie5vubirdu.eu-west-1.es.amazonaws.com/_cluster/stats?human&pretty")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(body))

	/*	var err error
		creds := credentials.NewStaticCredentials("", "", "")
		client, err = newElasticClient(creds)
		if err != nil {
			// Handle error
		}
		cluster := elastic.NewClusterStatsService(client)

		rep, err := cluster.Do()
		if err != nil {
			return
		}
	*/
}
