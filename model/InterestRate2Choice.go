package model

// Choice between a fixed rate and a floating rate.
type InterestRate2Choice struct {

	// Indicates that the rate is fixed.
	Fixed *PercentageRate `xml:"Fxd"`

	// Provides details about the variable rate.
	Floating *FloatingInterestRate4 `xml:"Fltg"`
}

func (i *InterestRate2Choice) SetFixed(value string) {
	i.Fixed = (*PercentageRate)(&value)
}

func (i *InterestRate2Choice) AddFloating() *FloatingInterestRate4 {
	i.Floating = new(FloatingInterestRate4)
	return i.Floating
}
