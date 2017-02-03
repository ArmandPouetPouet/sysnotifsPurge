package Config

import (
	"fmt"
	"os"
	"strconv"
)

//"https://search-core-sysnotifs-prod-7tlhh72bx4aasz7we2d3cjvybi.eu-west-1.es.amazonaws.com"
// "https://search-core-sysnotifs-beta-tsrx3dffhismokydie5vubirdu.eu-west-1.es.amazonaws.com"
const urlClusterStatsSuffix = "/_cluster/stats"
const urlStatsSuffix = "/_stats"

//URLCluster is the url to the elasticsearch cluster
var URLCluster string

//Threshold is the limit under which cluster should be purged (a percentage)
var Threshold int

//URLClusterStats is the url to get cluster specific information (nodes, size available, etc.)
var URLClusterStats string

//SYSNOTIFSPURGE_ES_ENDPOINT ="https://search-core-sysnotifs-prod-7tlhh72bx4aasz7we2d3cjvybi.eu-west-1.es.amazonaws.com"
//SYSNOTIFSPURGE_THRESHOLD =30

//URLStats is the url to get stats on cluster's usage (indices, etc.)
var URLStats string

func init() {
	fmt.Println("*************************************************")
	URLCluster = os.Getenv("SYSNOTIFSPURGE_ES_ENDPOINT")
	Threshold, _ = strconv.Atoi(os.Getenv("SYSNOTIFSPURGE_THRESHOLD"))
	URLClusterStats = URLCluster + urlClusterStatsSuffix
	URLStats = URLCluster + urlStatsSuffix
	fmt.Println("calls will be performed against URL : " + URLCluster)
	fmt.Println("*************************************************")
}
