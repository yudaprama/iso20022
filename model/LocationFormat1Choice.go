package model

// Choice between the formats to indicate the location of the meeting.
type LocationFormat1Choice struct {

	// Specifies physical address of the meeting
	Address *PostalAddress1 `xml:"Adr"`

	// Indicates that the location is unknown.
	LocationCode *PlaceType1Code `xml:"LctnCd"`
}

func (l *LocationFormat1Choice) AddAddress() *PostalAddress1 {
	l.Address = new(PostalAddress1)
	return l.Address
}

func (l *LocationFormat1Choice) SetLocationCode(value string) {
	l.LocationCode = (*PlaceType1Code)(&value)
}
