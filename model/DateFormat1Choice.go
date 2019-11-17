package model

// Choice between the formats to express a date.
type DateFormat1Choice struct {

	// Date expressed as a calendar date.
	Date *ISODate `xml:"Dt"`

	// Date expressed as a code.
	Code *SettlementDate1Code `xml:"Cd"`

	// Date expressed as a calendar date and time.
	DateTime *ISODateTime `xml:"DtTm"`
}

func (d *DateFormat1Choice) SetDate(value string) {
	d.Date = (*ISODate)(&value)
}

func (d *DateFormat1Choice) SetCode(value string) {
	d.Code = (*SettlementDate1Code)(&value)
}

func (d *DateFormat1Choice) SetDateTime(value string) {
	d.DateTime = (*ISODateTime)(&value)
}
