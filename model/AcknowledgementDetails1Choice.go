package model

// Reference of the PayInSchedule or the PayInCall being acknowledged. This is the Message Identification element from the Report Data sequence of the Pay In Schedule message or the Pay In Call message.
//
type AcknowledgementDetails1Choice struct {

	// Reference to the pay in schedule that is being acknowledged.
	PayInScheduleReference *Max35Text `xml:"PayInSchdlRef"`

	// Reference to the pay in call that is being acknowledged.
	PayInCallReference *Max35Text `xml:"PayInCallRef"`
}

func (a *AcknowledgementDetails1Choice) SetPayInScheduleReference(value string) {
	a.PayInScheduleReference = (*Max35Text)(&value)
}

func (a *AcknowledgementDetails1Choice) SetPayInCallReference(value string) {
	a.PayInCallReference = (*Max35Text)(&value)
}
