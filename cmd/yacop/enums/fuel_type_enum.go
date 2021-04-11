package enums

import (
	"fmt"
	"io"
	"strconv"
)

type FuelType string

const (
	FuelTypePetrol       FuelType = "PETROL"
	FuelTypeDiesel       FuelType = "DIESEL"
	FuelTypeElectricity  FuelType = "ELECTRICITY"
	FuelTypeNaturalGas   FuelType = "NATURAL_GAS"
	FuelTypeHybridPetrol FuelType = "HYBRID_PETROL"
	FuelTypeHybridDiesel FuelType = "HYBRID_DIESEL"
	FuelTypeHydrogen     FuelType = "HYDROGEN"
)

func (e FuelType) IsValid() bool {
	switch e {
	case FuelTypePetrol,
		FuelTypeDiesel,
		FuelTypeElectricity,
		FuelTypeNaturalGas,
		FuelTypeHybridPetrol,
		FuelTypeHybridDiesel,
		FuelTypeHydrogen:
		return true
	}
	return false
}

func (e FuelType) String() string {
	return string(e)
}

func (e *FuelType) Unmarshal(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FuelType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FuelType", str)
	}
	return nil
}

func (e FuelType) Marshal(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
