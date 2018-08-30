package write

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

// write writes data to a csv.
func write(header []string, prefix string, cmd string, data [][]string) error {
	t := time.Now()
	now := fmt.Sprintf("%d-%02d-%02d_%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	fileTitle := fmt.Sprintf("%v_%v_%v.csv", prefix, cmd, now)
	file, err := os.Create(fileTitle)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if !removeHeaderF {
		err = writer.Write(header)
		if err != nil {
			return err
		}
	}

	for _, value := range data {
		err = writer.Write(value)
		if err != nil {
			return err
		}
	}
	fmt.Println(fileTitle)
	return nil
}
