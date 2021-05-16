package file

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/zulkan/kraicklist/domain"
	"github.com/zulkan/kraicklist/utils"
	"os"
)

type fileRecordRepository struct {
	records []domain.Record
	titleList []string
	titleRecordMap map[string]domain.Record
}

func NewFileRecordRepository(filepath string) (domain.RecordRepository, error) {
	// open file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("unable to open source file due: %v", err)
	}
	defer file.Close()
	// read as gzip
	reader, err := gzip.NewReader(file)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize gzip reader due: %v", err)
	}
	// read the reader using scanner to contstruct records
	records := make([]domain.Record, 0)

	cs := bufio.NewScanner(reader)
	for cs.Scan() {
		var r domain.Record
		err = json.Unmarshal(cs.Bytes(), &r)
		if err != nil {
			continue
		}
		//fmt.Printf("Adding %s to records\n", extractedTitle)
		records = append(records, r)

	}
	fmt.Printf("Added %d records\n", len(records))

	return &fileRecordRepository{
		records:        records,
		titleList:      ConvertToTitleArray(records),
		titleRecordMap: ConvertToTitleRecordMap(records),
	}, nil
}
// extracted to function so can be reused and tested
func ConvertToTitleArray(records []domain.Record) []string {
	result := make([]string, 0)
	for _, record := range records {
		extractedTitle := utils.ExtractTitle(utils.TrimLowerString(record.Title))

		result = append(result, extractedTitle)
	}
	return result
}

// extracted to function so can be reused and tested
func ConvertToTitleRecordMap(records []domain.Record) map[string]domain.Record {
	result := make(map[string]domain.Record)
	for _, record := range records {
		extractedTitle := utils.ExtractTitle(utils.TrimLowerString(record.Title))

		result[extractedTitle] = record
	}
	return result
}

func (f *fileRecordRepository) FindAll() ([]domain.Record, error) {
	return f.records, nil
}

func (f *fileRecordRepository) FindAllTitle() ([]string) {
	return f.titleList
}

func (f *fileRecordRepository) FindByTitle(query string) (*domain.Record, error) {
	if data, isExist := f.titleRecordMap[query]; isExist {
		return &data, nil
	}

	return nil, fmt.Errorf("record '%s' not found", query)
}

func (f *fileRecordRepository) Count() int {
	return len(f.titleList)
}