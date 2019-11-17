package model

// Information about an individual person.
type CitizenshipInformation2 struct {

	// Country where a person was born or is legally accepted as belonging to the country.
	Nationality *NationalityCode `xml:"Ntlty"`

	// Indicates whether the person is a legal minor. This may depend on the nationality, the domicile country or the transaction in which the person is involved.
	MinorIndicator *YesNoIndicator `xml:"MnrInd"`
}

func (c *CitizenshipInformation2) SetNationality(value string) {
	c.Nationality = (*NationalityCode)(&value)
}

func (c *CitizenshipInformation2) SetMinorIndicator(value string) {
	c.MinorIndicator = (*YesNoIndicator)(&value)
}
