package model

// Specifies details related to the attendance card.
type AttendanceCard1 struct {

	// Information to be indicated on the attendance card.
	AttendanceCardLabelling *Max105Text `xml:"AttndncCardLbllg,omitempty"`

	// Specifies where the attendance card must be delivered.
	DeliveryMethod *DeliveryPlace1Code `xml:"DlvryMtd"`

	// Name and address of a party.
	OtherAddress *NameAndAddress9 `xml:"OthrAdr,omitempty"`
}

func (a *AttendanceCard1) SetAttendanceCardLabelling(value string) {
	a.AttendanceCardLabelling = (*Max105Text)(&value)
}

func (a *AttendanceCard1) SetDeliveryMethod(value string) {
	a.DeliveryMethod = (*DeliveryPlace1Code)(&value)
}

func (a *AttendanceCard1) AddOtherAddress() *NameAndAddress9 {
	a.OtherAddress = new(NameAndAddress9)
	return a.OtherAddress
}
