package convert

import (
	"BiTaksi_Backend_Driver/models"
	"BiTaksi_Backend_Driver/tools/errors"
	"encoding/csv"
	"github.com/jszwec/csvutil"
	"io"
	"os"
	"path"
	"runtime"
)

func CsvToStruct() []models.Coordinat {
	_, filename, _, _ := runtime.Caller(0)
	f := path.Join(path.Dir(filename), "Coordinates.csv")
	csvFile, err := os.Open(f)
	errors.StandartErrorWithErrorLog(err, nil)

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	userHeader, _ := csvutil.Header(models.Coordinat{}, "csv")
	dec, _ := csvutil.NewDecoder(reader, userHeader...)

	var coordinats []models.Coordinat
	var count int = 0

	for {
		var u models.Coordinat
		if err := dec.Decode(&u); err == io.EOF {
			break
		}
		coordinats = append(coordinats, u)
		count++
	}
	return coordinats
}
