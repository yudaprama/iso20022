package model

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type CitizenshipInformation struct {

	// Specifies the country where a person was born or is legally accepted as belonging to the country.
	Nationality *NationalityCode `xml:"Ntlty"`

	// Indicates whether the person is a legal minor. It may depend on the nationality, the domicile country or the transaction in which the person is involved.
	MinorIndicator *YesNoIndicator `xml:"MnrInd"`
}

func (c *CitizenshipInformation) SetNationality(value string) {
	c.Nationality = (*NationalityCode)(&value)
}

func (c *CitizenshipInformation) SetMinorIndicator(value string) {
	c.MinorIndicator = (*YesNoIndicator)(&value)
}
