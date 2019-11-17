package model

// Specifies the physical parameters of a shareholders meeting. Several dates and places can be defined for a meeting
type Meeting2 struct {

	// Date and time at which the meeting will take place.
	DateAndTime *DateFormat2Choice `xml:"DtAndTm"`

	// Indicates the status of a meeting date.
	DateStatus *MeetingDateStatus1Code `xml:"DtSts,omitempty"`

	// Specifies whether a minimum number of security representation is required to hold a meeting.
	QuorumRequired *YesNoIndicator `xml:"QrmReqrd"`

	// Specifies location where meeting will take place.
	Location []*LocationFormat1Choice `xml:"Lctn"`

	// Minimum quantity of securities required to hold a meeting.
	QuorumQuantity *Max35Text `xml:"QrmQty,omitempty"`

	// Minimum quantity of securities, expressed as a percentage, required to hold a meeting.
	QuorumQuantityPercentage *PercentageRate `xml:"QrmQtyPctg,omitempty"`
}

func (m *Meeting2) AddDateAndTime() *DateFormat2Choice {
	m.DateAndTime = new(DateFormat2Choice)
	return m.DateAndTime
}

func (m *Meeting2) SetDateStatus(value string) {
	m.DateStatus = (*MeetingDateStatus1Code)(&value)
}

func (m *Meeting2) SetQuorumRequired(value string) {
	m.QuorumRequired = (*YesNoIndicator)(&value)
}

func (m *Meeting2) AddLocation() *LocationFormat1Choice {
	newValue := new(LocationFormat1Choice)
	m.Location = append(m.Location, newValue)
	return newValue
}

func (m *Meeting2) SetQuorumQuantity(value string) {
	m.QuorumQuantity = (*Max35Text)(&value)
}

func (m *Meeting2) SetQuorumQuantityPercentage(value string) {
	m.QuorumQuantityPercentage = (*PercentageRate)(&value)
}
