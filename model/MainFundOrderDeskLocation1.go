package model

// U-003-2009 Addition and Modification of choice component. S-009-2009 They are not valid business options in the redemption processing context. S-015-2009 Add new data elements to indicate time zone.
type MainFundOrderDeskLocation1 struct {

	// Country in which it is authorised to commercialise the fund.
	Country *CountryCode `xml:"Ctry"`

	// Offset of the reporting time before or after 00:00 hour UTC.
	TimeZoneOffSet *UTCOffset1 `xml:"TmZoneOffSet"`
}

func (m *MainFundOrderDeskLocation1) SetCountry(value string) {
	m.Country = (*CountryCode)(&value)
}

func (m *MainFundOrderDeskLocation1) AddTimeZoneOffSet() *UTCOffset1 {
	m.TimeZoneOffSet = new(UTCOffset1)
	return m.TimeZoneOffSet
}
