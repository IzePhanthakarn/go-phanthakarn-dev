package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// CSV csv model
type CSV struct {
	path string
}

// New new csv model
func New(path string) *CSV {
	return &CSV{
		path: path,
	}
}

// Read read
func (c *CSV) Read(f func(row []string)) error {
	ff, err := os.Open(c.path)
	if err != nil {
		return err
	}
	r := csv.NewReader(ff)
	_, _ = r.Read()
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		f(row)
	}
	return nil
}

// ReadWithError csv with handle error
func (c *CSV) ReadWithError(f func(row []string) error) error {
	ff, err := os.Open(c.path)
	if err != nil {
		return err
	}
	r := csv.NewReader(ff)
	_, _ = r.Read()
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		err = f(row)
		if err != nil {
			return err
		}
	}
	return nil
}

// ReadWithHeader read with header
func (c *CSV) ReadWithHeader(header []string, f func(row []string)) error {
	ff, err := os.Open(c.path)
	if err != nil {
		return err
	}
	r := csv.NewReader(ff)
	row, err := r.Read()
	if err != nil {
		return fmt.Errorf("invalid csv format")
	}
	if len(row) != len(header) {
		return fmt.Errorf("invalid header")
	}
	for index := 0; index < len(header); index++ {
		c := row[index]
		if index == 0 {
			b := []byte(c)
			if b[0] == 0xef || b[1] == 0xbb || b[2] == 0xbf {
				c = string(b[3:])
			}
		}
		if header[index] != c {
			return fmt.Errorf("invalid header")
		}
	}
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		f(row)
	}
	return nil
}

// ReadAll read all data
func (c *CSV) ReadAll() ([][]string, error) {
	ff, err := os.Open(c.path)
	if err != nil {
		return nil, err
	}
	return csv.NewReader(ff).ReadAll()
}
