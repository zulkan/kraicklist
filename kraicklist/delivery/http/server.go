package http

import (
	"fmt"
	"github.com/zulkan/kraicklist/domain"
	"net/http"
)

type server struct {
	searcherService domain.Searcher
}

func NewHttpServer(port int, searcherService domain.Searcher) error {

	svr := &server{searcherService: searcherService}

	http.Handle("/", http.FileServer(http.Dir("./kraicklist/delivery/static")))
	http.HandleFunc("/search", svr.handleSearch)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	return err
}
