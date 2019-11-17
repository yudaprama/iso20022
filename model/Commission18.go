package model

// Amount of money due to a party as compensation for a service.
type Commission18 struct {

	// Commission expressed as a percentage.
	Rate *PercentageRate `xml:"Rate"`

	// Additional information about the type of markup.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *Commission18) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *Commission18) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
