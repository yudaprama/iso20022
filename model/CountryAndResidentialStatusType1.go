package model

// Country and residential status.
type CountryAndResidentialStatusType1 struct {

	// Country to which the residential status is applicable.
	Country *CountryCode `xml:"Ctry"`

	// Residential status of an individual, for example, non-permanent resident.
	ResidentialStatus *ResidentialStatus1Code `xml:"ResdtlSts"`
}

func (c *CountryAndResidentialStatusType1) SetCountry(value string) {
	c.Country = (*CountryCode)(&value)
}

func (c *CountryAndResidentialStatusType1) SetResidentialStatus(value string) {
	c.ResidentialStatus = (*ResidentialStatus1Code)(&value)
}
