package tools

import (
	"encoding/csv"
	"golang.org/x/text/transform"
	"os"
	"golang.org/x/text/encoding"
	"io"
	"github.com/pkg/errors"
)

func ReadCSV(file *os.File, decoder *encoding.Decoder) ([]string, error){
	reader := csv.NewReader(transform.NewReader(file, decoder))

	var contents []string

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else {
			return nil, errors.New("fail to read tools file.")
		}

		for i, v := range line {
			if i > 0 {
				contents = append(contents, v)
			}
		}

	}
	return contents, nil
}


