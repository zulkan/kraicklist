package usecase

import (
	"fmt"
	"github.com/hbollon/go-edlib"
	"github.com/schollz/closestmatch"
	"github.com/zulkan/kraicklist/domain"
	"github.com/zulkan/kraicklist/utils"
	"strings"
)

type searcherUseCase struct {
	recordRepo domain.RecordRepository
	closesMatch *closestmatch.ClosestMatch
}

func NewSearcherUseCase(repo domain.RecordRepository) domain.Searcher {
	bagSizes := []int{3, 4, 5, 6, 7, 8, 9, 10}

	return &searcherUseCase{
		recordRepo: repo,
		closesMatch: closestmatch.New(repo.FindAllTitle(), bagSizes),
	}
}

func (s searcherUseCase) Search(query string) ([]domain.Record, error) {
	closesMatchRes := s.getResultByClosesMatch(query)
	fmt.Println("matches using closesMatch:", len(closesMatchRes))

	edlibMatch := s.getResultByEdLib(query)
	fmt.Println("matches using edlib :", len(edlibMatch))

	resultAll := MergeResult(closesMatchRes, edlibMatch)

	result := make([]domain.Record, 0)
	for _, title := range resultAll {
		//fmt.Printf("adding %s, pos %v\n", title, dataPair[1])
		data, err := s.recordRepo.FindByTitle(title)
		if err != nil {
			fmt.Println("Error ", err)
			continue
		}
		result = append(result, *data)
	}

	return result, nil
}

func (s searcherUseCase) getResultByEdLib(query string) []string {
	res, _ :=
		edlib.FuzzySearchSetThreshold(query, s.recordRepo.FindAllTitle(), 50, 0.4, edlib.Levenshtein)

	return res
}

func (s searcherUseCase) getResultByClosesMatch(query string) []string {
	cm := s.closesMatch
	titleMatches := cm.ClosestN(query, s.recordRepo.Count())
	fmt.Println(cm.AccuracyMutatingWords())

	return titleMatches
}
// merge result and make it uniq, but maintain order
func MergeResult(resultArrList ...[]string) []string {
	joinedResult := make([]string, 0)

	existingData := make(map[string]bool)
	for _, resArr := range resultArrList {
		for _, res := range resArr {
			_, isExist := existingData[res]; if isExist || res == "" {
				continue
			}
			existingData[res] = true

			joinedResult = append(joinedResult, res)
		}
	}

	return joinedResult
}

func match(text, query string) bool {

	text = utils.TrimLowerString(text)
	query = utils.TrimLowerString(query)

	text = utils.ExtractTitle(text)

	lowerTextArr := strings.Split(text," ")

	//1 word query match
	for _, lowerTextSplitted := range lowerTextArr {
		if lowerTextSplitted == query {
			fmt.Printf("matched %s with %s\n", text, query)

			return true
		}
	}

	//multiple words query match
	queryTextArr := strings.Split(strings.ToLower(query)," ")
	len := len(queryTextArr)
	if len == 1 {
		return false
	}
	pos := 0
	for _, lowerTextSplitted := range lowerTextArr {
		if lowerTextSplitted == queryTextArr[pos] {
			pos++
		}
		if pos == len {
			fmt.Printf("matched %s with %s\n", text, query)

			return true
		}
	}


	return false
}
