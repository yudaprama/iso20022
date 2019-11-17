package model

// Specifies a country by its code or its name.
type CountryCodeAndName1 struct {

	// Country is specified by its code.
	Code *CountryCode `xml:"Cd,omitempty"`

	// Country is specified by its name.
	Name *Max35Text `xml:"Nm,omitempty"`
}

func (c *CountryCodeAndName1) SetCode(value string) {
	c.Code = (*CountryCode)(&value)
}

func (c *CountryCodeAndName1) SetName(value string) {
	c.Name = (*Max35Text)(&value)
}
