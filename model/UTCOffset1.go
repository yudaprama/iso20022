package model

// Specifies the parameters to calculate the local reporting time.
type UTCOffset1 struct {

	// Indicates whether the offset is before or after 00:00 hour UTC.
	Sign *PlusOrMinusIndicator `xml:"Sgn"`

	// Offset of the reporting time, in hours,  before or after 00:00 hour UTC.
	NumberOfHours *ISOTime `xml:"NbOfHrs"`
}

func (u *UTCOffset1) SetSign(value string) {
	u.Sign = (*PlusOrMinusIndicator)(&value)
}

func (u *UTCOffset1) SetNumberOfHours(value string) {
	u.NumberOfHours = (*ISOTime)(&value)
}
