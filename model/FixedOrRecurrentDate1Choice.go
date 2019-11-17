package model

// Choice between a fixed date and a recurring date.
type FixedOrRecurrentDate1Choice struct {

	// Date on which the variation is triggered.
	FixedDate *ISODate `xml:"FxdDt"`

	// Details related to recurrent dates on which the variation is triggered.
	RecurrentDate *DateInformation1 `xml:"RcrntDt"`
}

func (f *FixedOrRecurrentDate1Choice) SetFixedDate(value string) {
	f.FixedDate = (*ISODate)(&value)
}

func (f *FixedOrRecurrentDate1Choice) AddRecurrentDate() *DateInformation1 {
	f.RecurrentDate = new(DateInformation1)
	return f.RecurrentDate
}
