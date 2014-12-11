package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"runtime"
	"time"
)

const (
	PerPage = 200
	Pages   = 5
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Fprintln(w, "Welcome!")
}

func PeopleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := r.URL.Query()
	apiKey := vars.Get("api_key")
	appKey := vars.Get("app_key")

	// api := &PipelineApi{
	// 	ApiKey:  apiKey,
	// 	AppKey:  appKey,
	// 	PerPage: 1,
	// }

	// fmt.Fprintf(w, "%v", apiKey)
	// fmt.Fprintf(w, "%v", appKey)
	//

	// total, _ := api.PeopleTotal()
	//
	// fmt.Fprintf(w, "%v", fmt.Sprintf("%d", total))

	c := make(chan *PipelineApiResponse)

	for page := 1; page <= Pages; page++ {
		go func(page int, co chan<- *PipelineApiResponse) {
			start := time.Now()

			a := &PipelineApi{
				ApiKey:  apiKey,
				AppKey:  appKey,
				Page:    page,
				PerPage: PerPage,
			}

			r, _ := a.People()

			log.Printf("Finished run for thread %d in %s", page, time.Since(start))

			co <- r
		}(page, c)
	}

	var pRes PipelineApiResponse
	// fmt.Fprintf(w, "%v", body)
	for i := 1; i <= Pages; i++ {
		res := <-c
		pRes.Entries = append(pRes.Entries, res.Entries)
		pRes.Pagination.PerPage += res.Pagination.PerPage
		pRes.Pagination.Total = res.Pagination.Total

		d := float64(res.Pagination.Pages) / float64(Pages)

		pRes.Pagination.Pages = int(math.Ceil(d))
		pRes.Pagination.Page = 1
	}

	if err := json.NewEncoder(w).Encode(pRes); err != nil {
		panic(err)
	}
}
