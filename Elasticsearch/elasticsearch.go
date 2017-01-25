package Elasticsearch

import (
	Helpers "SysnotifsPurge/Helpers"
	"encoding/json"
	"sort"
	"strings"
)

//CleanES checks current data load of the cluster, and perform purge if necessary
func CleanES() int {
	for getRatio() < threshold {
		purgeOldestIndex()
	}
	return getRatio()
}

func getRatio() int {
	var data clusterStats
	json.Unmarshal(Helpers.Get(urlClusterStats), &data)

	return int(float64(data.Nodes.Fs.AvailableInBytes) / float64(data.Nodes.Fs.TotalInBytes) * 100)
}

//getSysnotifsIndices call cluster and get only sysnotifs indices
func getSysnotifsIndices() []string {

	//Request ES cluster to get available indices
	var data indexStruct
	json.Unmarshal(Helpers.Get(urlStats), &data)

	return filterAndOrderSysnotifsIndices(data)

}

func purgeOldestIndex() {
	purgeIndex(getSysnotifsIndices()[0])
}

//filterAndOrderSysnotifsIndices remove non sysnotifs indices and order chronologically
func filterAndOrderSysnotifsIndices(data indexStruct) []string {
	var keys []string
	for key := range data.Indices {
		if !strings.Contains(key, "kibana") {
			keys = append(keys, key)
		}

	}
	sort.Strings(keys)
	return keys
}

//purgeIndex call cluster to delete index in parameter
func purgeIndex(indexName string) {
	if strings.Contains(indexName, "kibana") {
		panic("kibana index shall not be touched")
	}
	Helpers.Delete(urlCluster + "/" + indexName)

}
