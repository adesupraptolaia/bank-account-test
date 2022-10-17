package helper

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/adesupraptolaia/bank-account-test/model"
)

func convertFloat64ToString(input float64) string {
	return fmt.Sprintf("%.0f", input)
}

func convertIntToString(input int) string {
	return fmt.Sprintf("%d", input)
}

func convertStringToFloat64(input string) float64 {
	result, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatalf("error when 'convertStringToFloat64' with input: '%s' and error: %s", input, err.Error())
	}

	return result
}

func convertStringToInt(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf("error when 'convertStringToInt' with input: '%s' and error: %s", input, err.Error())
	}

	return result
}

func ReadCSVFile(fileName string) []model.AfterEOD {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("error when open file '%s' with error : %s", fileName, err.Error())
	}

	// remember to close the file at the end of the program
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'

	records := make([]model.AfterEOD, 200)

	for i := 0; i <= 200; i++ {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error when read file on line %d with error: %s", i, err.Error())
		}
		if i == 0 {
			continue
		}
		records[i-1] = model.AfterEOD{
			ID:              record[0],
			Name:            record[1],
			Age:             record[2],
			Balance:         convertStringToFloat64(record[3]),
			PreviousBalance: convertStringToFloat64(record[4]),
			AverageBalance:  convertStringToFloat64(record[5]),
			FreeTransfer:    convertStringToInt(record[6]),
		}
	}

	return records
}

func WriteToCSV(fileName string, records []model.AfterEOD) {
	result := make([][]string, len(records)+1)

	// define header
	result[0] = []string{"id", "Nama", "Age", "Balanced", "No 2b Thread-No", "No 3 Thread-No", "Previous Balanced", "Average Balanced", "No 1 Thread-No", "Free Transfer", "No 2a Thread-No"}

	for i := 0; i < len(records); i++ {
		record := records[i]

		result[i+1] = []string{
			record.ID,
			record.Name,
			record.Age,
			convertFloat64ToString(record.Balance),
			record.ThreadNo2b,
			record.ThreadNo3,
			convertFloat64ToString(record.PreviousBalance),
			convertFloat64ToString(record.AverageBalance),
			record.ThreadNo1,
			convertIntToString(record.FreeTransfer),
			record.ThreadNo2a,
		}
	}

	os.Remove(fileName)

	csvFile, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("error when create file '%s' with error: %s", fileName, err.Error())
	}

	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	csvwriter.Comma = ';'

	csvwriter.WriteAll(result)
}
