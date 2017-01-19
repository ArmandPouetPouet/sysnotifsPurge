package Elasticsearch

type indexStruct struct {
	Indices map[string]struct {
		Primaries string `json:"primaries"`
	} `json:"indices"`
}

type clusterStats struct {
	Nodes struct {
		Fs struct {
			TotalInBytes     int `json:"total_in_bytes"`
			AvailableInBytes int `json:"available_in_bytes"`
		} `json:"fs"`
	} `json:"nodes"`
}
