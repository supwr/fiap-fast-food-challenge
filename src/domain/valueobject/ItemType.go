package valueobject

import (
	"database/sql/driver"
	"fmt"
	"slices"
	"strings"
)

const (
	Beverage   = "BEVERAGE"
	Food       = "FOOD"
	Dessert    = "DESSERT"
	Ingredient = "INGREDIENT"
	Sides      = "SIDES"
)

type ItemType struct {
	value string
}

func NewItemType(i string) (*ItemType, error) {
	itemType := &ItemType{
		value: i,
	}

	if err := itemType.validate(); err != nil {
		return nil, err
	}

	return itemType, nil
}

func (i *ItemType) validate() error {
	itemTypes := []string{Beverage, Food, Dessert, Ingredient, Sides}

	if !slices.Contains(itemTypes, i.value) {
		return fmt.Errorf("invalid item type: %s", i.value)
	}

	return nil
}

func (d *ItemType) String() string {
	return strings.ToUpper(d.value)
}

func (d *ItemType) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *ItemType) UnmarshalText(v []byte) error {
	var err error
	d, err = NewItemType(string(v))
	return err
}

func (d *ItemType) Scan(value interface{}) error {
	*d = ItemType{value: strings.ToUpper(fmt.Sprint(value))}
	return nil
}

func (d ItemType) Value() (driver.Value, error) {
	if len(d.String()) == 0 {
		return nil, nil
	}

	return d.String(), nil
}
