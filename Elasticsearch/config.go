package Elasticsearch

const urlCluster string = "https://search-core-sysnotifs-prod-7tlhh72bx4aasz7we2d3cjvybi.eu-west-1.es.amazonaws.com"

//const urlCluster string = "https://search-core-sysnotifs-beta-tsrx3dffhismokydie5vubirdu.eu-west-1.es.amazonaws.com"
const urlClusterStats string = urlCluster + "/_cluster/stats"
const urlStats string = urlCluster + "/_stats"

const threshold = 20
