package main

import (
	"net/http"
	"time"

	"github.com/dragorific/makeuoft_wildfirepredictor/libraries/apilib"
	"github.com/dragorific/makeuoft_wildfirepredictor/libraries/elasticsearch"
	"github.com/dragorific/makeuoft_wildfirepredictor/setup"
	"github.com/gorilla/mux"
	"github.com/olivere/elastic"
)

func main() {
	//gets state file
	s := setup.GetMainState("api engine")

	//creates new router for api
	router := mux.NewRouter()
	router.StrictSlash(true)

	//sets up api subrouter
	api := router.PathPrefix("/api/").Subrouter()

	//Lets user know if route is working
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("healthy"))
	})
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 not found"))
	})

	setUpRoutes(s, router, api)

	server := &http.Server{
		Addr:         ":6060",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// start HTTP server
	s.Log.Info("http endpoint now active on :6060")
	err := server.ListenAndServe()
	if err != nil {
		s.Log.Fatal(err)
	}
}

func setUpRoutes(s *setup.State, router *mux.Router, api *mux.Router) {

	api.HandleFunc("/get-markers", func(w http.ResponseWriter, r *http.Request) {
		s.Log.Info("new request on /get-markers")
		if r.Method != "GET" {
			s.Log.Error("/get-markers did not receive a get request")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad request"))
			return
		}

		if !elasticsearch.ExistsByID(s, "markers", "markers") {
			client, ctx := s.Elastic, s.Ctx
			data := `{"markers":[["US","37.0902","-95.7129"],["Canada","55.585901","-105.750596"]]}`
			_, err := client.Index().Index("markers").Id("markers").BodyJson(data).Do(ctx)
			if err != nil {
				s.Log.Error("error indexing markers document to markers ", err)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(data))
			return
		}

		docJSON, err := elasticsearch.GetDocumentByID(s, "markers", "markers")
		if err != nil {
			s.Log.Error("Unable to get data from markers index", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad request"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(docJSON))
		return
	})
	api.HandleFunc("/getData-{name}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			s.Log.Error("/get-markers did not receive a get request")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad request"))
			return
		}
		name := mux.Vars(r)
		markerName := name["name"]
		client, ctx := s.Elastic, s.Ctx
		result, err := client.Search().Index(markerName).SortBy(elastic.NewFieldSort("timestamp").Asc()).From(0).Size(30).Do(ctx)
		if err != nil {
			s.Log.Error("Unable to get data from getData index", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad request"))
			return
		}
		var arr [30][]byte
		for i, data := range result.Hits.Hits {
			arr[i] = data.Source
		}
		returnedData := apilib.ParseSensorData(arr)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(returnedData))
		return
	})

}
