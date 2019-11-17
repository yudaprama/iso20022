package model

// Charge is expressed as a code or a bilaterally agreed code.
type ChargeTypeChoice struct {

	// Type of service for which a charge is asked or paid.
	Code *ChargeType1Code `xml:"Cd"`

	// Type of charge is a bilaterally agreed code.
	ProprietaryCode *Max4AlphaNumericText `xml:"PrtryCd"`
}

func (c *ChargeTypeChoice) SetCode(value string) {
	c.Code = (*ChargeType1Code)(&value)
}

func (c *ChargeTypeChoice) SetProprietaryCode(value string) {
	c.ProprietaryCode = (*Max4AlphaNumericText)(&value)
}
