package Elasticsearch

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

//Get performs a http GET call and read body as byte[]
func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		errors.Wrap(err, "Get failed")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		return nil, errors.New("Get failed with statusCode :" + strconv.Itoa(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errors.Wrap(err, "Get response Body could not be read")
	}
	return body, err
}

//Delete performs a http DELETE call
func Delete(url string) error {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		errors.Wrap(err, "Delete request could not be built")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errors.Wrap(err, "Delete failed")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		return errors.New("Delete failed with statusCode :" + strconv.Itoa(resp.StatusCode))
	}
	return err
}
