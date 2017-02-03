package Elasticsearch

import (
	config "SysnotifsPurge/Config"
	"encoding/json"
	"log"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

//CleanES checks current data load of the cluster, and perform purge if necessary
func CleanES() (ratio int, err error) {
	ratio, err = getRatio()
	for ratio < config.Threshold {
		err = purgeOldestSysnotifsIndex()
		ratio, err = getRatio()
		if err != nil {
			break
		}
	}
	return
}

//getRatio call cluster to get stats, and compute ratio CurrentlyAvailable / ClusterSize
func getRatio() (ratio int, err error) {
	var data clusterStats
	body, err := Get(config.URLClusterStats)
	if err != nil {
		return 0, errors.Wrap(err, "Can't get cluster stats")
	}
	err = json.Unmarshal(body, &data)
	ratio = int(float64(data.Nodes.Fs.AvailableInBytes) / float64(data.Nodes.Fs.TotalInBytes) * 100)
	return
}

//purgeOldestSysnotifsIndex find oldest index for sysnotifs-yyyy.mm.dd and ask cluster to delete it
func purgeOldestSysnotifsIndex() error {
	indices, err := getSysnotifsIndices()
	if err != nil {
		return errors.Wrap(err, "Can't purge oldest index")
	}
	if indices == nil || len(indices) <= 0 {
		return errors.New("No indices found")
	}
	err = purgeIndex(indices[0])
	return err
}

//getSysnotifsIndices call cluster and get only sysnotifs indices
func getSysnotifsIndices() (indices []string, err error) {
	var data indexStruct
	body, err := Get(config.URLStats)
	err = json.Unmarshal(body, &data)
	indices = filterAndOrderSysnotifsIndices(data)
	return
}

//filterAndOrderSysnotifsIndices remove non sysnotifs indices and order chronologically. NOT expecting any error here
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
func purgeIndex(indexName string) error {
	if strings.Contains(indexName, "kibana") {
		log.Fatal("kibana index shall not be touched")
	}
	err := Delete(config.URLCluster + "/" + indexName)
	if err == nil {
		errors.Wrap(err, "could not purge index : "+indexName)
	}
	return err
}
