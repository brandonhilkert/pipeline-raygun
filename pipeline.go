package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const ApiUrl = "https://api.pipelinedeals.com"

type PipelineApi struct {
	ApiKey  string
	AppKey  string
	Page    int
	PerPage int
}

type PipelineApiResponse struct {
	Entries    []interface{} `json:"entries"`
	Pagination Pagination    `json:"pagination"`
}

type Pagination struct {
	Page    int `json:"page"`
	Pages   int `json:"pages"`
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
}

func (api *PipelineApi) PeopleTotal() (total int, err error) {
	var pRes PipelineApiResponse

	url := fmt.Sprintf("%s/api/v3/people.json?api_key=%s&app_key=%s&per_page=%d", ApiUrl, api.ApiKey, api.AppKey, api.PerPage)

	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return
	}

	err = json.Unmarshal(body, &pRes)

	if err != nil {
		return
	}

	total = pRes.Pagination.Total

	return
}

func (api *PipelineApi) People() (*PipelineApiResponse, error) {
	var pRes PipelineApiResponse

	url := fmt.Sprintf("%s/api/v3/people.json?api_key=%s&app_key=%s&page=%d", ApiUrl, api.ApiKey, api.AppKey, api.Page)

	res, err := http.Get(url)
	if err != nil {
		return &pRes, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return &pRes, err
	}

	err = json.Unmarshal(body, &pRes)

	if err != nil {
		return &pRes, err
	}

	return &pRes, nil
}
