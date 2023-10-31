package valueobject

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

const (
	cpfLength  = 11
	cnpjLength = 14
)

type Document struct {
	value string
}

func NewDocument(d string) (*Document, error) {
	document := &Document{
		value: d,
	}

	if err := document.validate(); err != nil {
		return nil, err
	}

	return document, nil
}

func (d *Document) validate() error {
	if len(d.value) != cpfLength && len(d.value) != cnpjLength {
		return errors.New("document needs to be a valid CPF or CNPJ")
	}

	return nil
}

func (d *Document) String() string {
	return d.value
}

func (d *Document) MarshalText() ([]byte, error) {
	return []byte(d.value), nil
}

func (d *Document) UnmarshalText(v []byte) error {
	var err error
	d, err = NewDocument(string(v))
	return err
}

func (d *Document) Scan(value interface{}) error {
	*d = Document{value: fmt.Sprint(value)}
	return nil
}

func (d *Document) Value() (driver.Value, error) {
	if len(d.value) == 0 {
		return nil, nil
	}

	return d.value, nil
}
