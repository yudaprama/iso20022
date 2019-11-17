package model

// Specifies the high level purpose of the instruction based on a set of pre-defined categories.
// Usage: This is used by the initiating party to provide information concerning the processing of the payment. It is likely to trigger special processing by any of the agents involved in the payment chain.
type CategoryPurpose1Choice struct {

	// Category purpose, as published in an external category purpose code list.
	Code *ExternalCategoryPurpose1Code `xml:"Cd"`

	// Category purpose, in a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (c *CategoryPurpose1Choice) SetCode(value string) {
	c.Code = (*ExternalCategoryPurpose1Code)(&value)
}

func (c *CategoryPurpose1Choice) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}
