package http

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (svr server) handleSearch(w http.ResponseWriter, r *http.Request) {
		// fetch query string from query params
		q := r.URL.Query().Get("q")
		if len(q) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing search query in query params"))
			return
		}
		// search relevant records
		records, err := svr.searcherService.Search(q)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		// output success response
		buf := new(bytes.Buffer)
		encoder := json.NewEncoder(buf)
		encoder.Encode(records)
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf.Bytes())

}
