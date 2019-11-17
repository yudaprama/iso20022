package model

// Specifies the billing rate.
type BillingRate1 struct {

	// Defines the type of rate or factor.
	Identification *BillingRateIdentification1Choice `xml:"Id"`

	// Value of the rate or factor identified in the rate identification.
	Value *PercentageRate `xml:"Val"`

	// Number of days in the statement period.
	//
	// Usage: Used along with DaysInYear for time dependent per annum rate value.
	DaysInPeriod *Number `xml:"DaysInPrd,omitempty"`

	// Number of days in the year.
	//
	// Usage: Used along with DaysInPeriod for time dependent per annum rate value.
	DaysInYear *Number `xml:"DaysInYr,omitempty"`
}

func (b *BillingRate1) AddIdentification() *BillingRateIdentification1Choice {
	b.Identification = new(BillingRateIdentification1Choice)
	return b.Identification
}

func (b *BillingRate1) SetValue(value string) {
	b.Value = (*PercentageRate)(&value)
}

func (b *BillingRate1) SetDaysInPeriod(value string) {
	b.DaysInPeriod = (*Number)(&value)
}

func (b *BillingRate1) SetDaysInYear(value string) {
	b.DaysInYear = (*Number)(&value)
}
