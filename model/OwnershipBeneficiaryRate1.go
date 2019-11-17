package model

// Percentage of ownership or of beneficial ownership of the shares/units in the account.
type OwnershipBeneficiaryRate1 struct {

	// Ownership or beneficial ownership expressed as a percentage.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Ownership or beneficial ownership expressed as a fraction or another form.
	Fraction *Max35Text `xml:"Frctn,omitempty"`
}

func (o *OwnershipBeneficiaryRate1) SetRate(value string) {
	o.Rate = (*PercentageRate)(&value)
}

func (o *OwnershipBeneficiaryRate1) SetFraction(value string) {
	o.Fraction = (*Max35Text)(&value)
}
