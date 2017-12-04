package tools

import (
	"os"
)

func ReadFile(fileName string, dist *os.File)(error){
	dist, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer dist.Close()

	return nil
}

