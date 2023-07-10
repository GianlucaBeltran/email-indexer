package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Inbox struct {
	Search_type string   `json:"search_type"`
	Sort_fields []string `json:"sort_fields"`
	From        int      `json:"from"`
	Max_results int      `json:"max_results"`
	Source      []string `json:"_source"`
}

func getInbox(w http.ResponseWriter, r *http.Request) {
	auth := "admin:Complexpass#123"

	currentPage, err := strconv.Atoi(chi.URLParam(r, "currentPage"))
	if err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("error fetching inbox: %s", err)))
		return
	}

	currentMaxPerPage, err := strconv.Atoi(chi.URLParam(r, "currentMaxPerPage"))

	if err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("error fetching inbox: %s", err)))
		return
	}

	i := Inbox{
		Search_type: "alldocuments",
		Sort_fields: []string{"-Date"},
		From:        currentPage * currentMaxPerPage,
		Max_results: currentMaxPerPage,
		Source:      []string{"From", "To", "Date"},
	}

	ij, err := json.Marshal(i)

	fmt.Println(string(ij))

	if err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("error fetching inbox: %s", err)))
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:4080/api/enron_mail/_search", bytes.NewReader(ij))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(auth)))

	if err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("error fetching inbox: %s", err)))
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("error fetching inbox: %s", err)))
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	w.Write([]byte(body))
}

type Email struct {
	Search_type string                 `json:"search_type"`
	Query       map[string]interface{} `json:"query"`
	Source      []string               `json:"_source"`
}

func getEmail(w http.ResponseWriter, r *http.Request) {
	auth := "admin:Complexpass#123"

	emailID := chi.URLParam(r, "emailID")

	e := Email{
		Search_type: "term",
		Query: map[string]interface{}{
			"term":  emailID,
			"field": "_id",
		},
		Source: []string{},
	}

	ej, _ := json.Marshal(e)
	fmt.Println(string(ej))
	req, err := http.NewRequest("POST", "http://localhost:4080/api/enron_mail/_search", bytes.NewReader(ej))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(auth)))

	if err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("error fetching email: %s", err)))
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("error fetching email: %s", err)))
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	w.Write([]byte(body))

}

type Search struct {
	Search_type string                 `json:"search_type"`
	Query       map[string]interface{} `json:"query"`
	From        int                    `json:"from"`
	Max_results int                    `json:"max_results"`
	Source      []string               `json:"_source"`
	Highlight   map[string]interface{} `json:"highlight"`
}

func getSearch(w http.ResponseWriter, r *http.Request) {
	auth := "admin:Complexpass#123"

	searchTerm := chi.URLParam(r, "searchTerm")
	from, _ := strconv.Atoi(chi.URLParam(r, "from"))
	maxResults, _ := strconv.Atoi(chi.URLParam(r, "maxResults"))

	s := Search{
		Search_type: "querystring",
		Query: map[string]interface{}{
			"term": searchTerm,
		},
		From:        from,
		Max_results: maxResults,
		Source:      []string{},
		Highlight: map[string]interface{}{
			"pre_tags":  []string{"<mark class='bg-yellow-300'>"},
			"post_tags": []string{"</mark>"},
			"fields": map[string]interface{}{
				"From": map[string]interface{}{
					"pre_tags":  []string{},
					"post_tags": []string{},
				},
				"To": map[string]interface{}{
					"pre_tags":  []string{},
					"post_tags": []string{},
				},
				"Date": map[string]interface{}{
					"pre_tags":  []string{},
					"post_tags": []string{},
				},
			},
		},
	}

	sj, _ := json.Marshal(s)
	fmt.Println(string(sj))
	req, err := http.NewRequest("POST", "http://localhost:4080/api/enron_mail/_search", bytes.NewReader(sj))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(auth)))

	if err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("error fetching email: %s", err)))
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("error fetching email: %s", err)))
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	w.Write([]byte(body))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Get("/inbox/{currentPage}-{currentMaxPerPage}", getInbox)
	r.Get("/email/{emailID}", getEmail)
	r.Get("/search/{searchTerm}/{from}-{maxResults}", getSearch)

	http.ListenAndServe(":3000", r)
}
