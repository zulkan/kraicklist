package usecase

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zulkan/kraicklist/domain"
	fileRecordRepo "github.com/zulkan/kraicklist/kraicklist/repository/file"
	"github.com/zulkan/kraicklist/mocks"

	"testing"
)

var recordData = []domain.Record{
	{
		ID:        1,
		Title:     "iphone 1",
		Content:   "selling an iphone",
		ThumbURL:  "",
		Tags:      nil,
		UpdatedAt: 0,
		ImageURLs: nil,
	},
	{
		ID:        2,
		Title:     "iphone X",
		Content:   "i have iphone x",
		ThumbURL:  "",
		Tags:      nil,
		UpdatedAt: 0,
		ImageURLs: nil,
	},
	{
		ID:        3,
		Title:     "iphone 8",
		Content:   "not used anymore, good condition",
		ThumbURL:  "",
		Tags:      nil,
		UpdatedAt: 0,
		ImageURLs: nil,
	},
	{
		ID:        100,
		Title:     "random element",
		Content:   "i am batman",
		ThumbURL:  "",
		Tags:      nil,
		UpdatedAt: 0,
		ImageURLs: nil,
	},
	{
		ID:        200,
		Title:     "hehe",
		Content:   "happy eid",
		ThumbURL:  "",
		Tags:      nil,
		UpdatedAt: 0,
		ImageURLs: nil,
	},
}

func TestSearcherUseCase_Search(t *testing.T) {
	repoMock := mocks.RecordRepository{}
	repoMock.On("FindAllTitle").Return(fileRecordRepo.ConvertToTitleArray(recordData))
	repoMock.On("Count").Return(len(recordData))
	repoMock.On("FindByTitle", mock.AnythingOfType("string")).Return(
		func(title string) *domain.Record {
			if data, isExist := fileRecordRepo.ConvertToTitleRecordMap(recordData)[title]; isExist {
				return &data
			}
			return nil
		},
		func(title string) error {
			if _, isExist := fileRecordRepo.ConvertToTitleRecordMap(recordData)[title]; !isExist {
				return fmt.Errorf("record '%s' not found", title)
			}
			return nil
		},
	)

	searcher := NewSearcherUseCase(&repoMock)
	res, err := searcher.Search("iphone")

	if assert.Nil(t, err) {
		assert.Equal(t, 3, len(res)) //have 3 match
	}
}

func TestMergeResult_removeDuplicateAndMaintainOrder(t *testing.T) {
	title1 := []string{"abc", "def"}
	title2 := []string{"def", "xyz"}

	result := MergeResult(title1, title2)

	assert.Equal(t, []string{"abc", "def", "xyz"}, result)
}

func TestMergeResult_removeEmptyString(t *testing.T) {
	title1 := []string{"abc", "def"}
	title2 := []string{"", "xyz"}

	result := MergeResult(title1, title2)

	assert.Equal(t, []string{"abc", "def", "xyz"}, result)
}