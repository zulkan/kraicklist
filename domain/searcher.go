package domain

import (
)

type Searcher interface {
	Search(query string) ([]Record, error)
}
