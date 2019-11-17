package model

// Provides information on an event that happened in a system.
type Event1 struct {

	// Proprietary code used to specify an event that occurred in a system.
	EventCode *Max4AlphaNumericText `xml:"EvtCd"`

	// Describes the parameters of an event which occurred in a system.
	EventParameter []*Max35Text `xml:"EvtParam,omitempty"`

	// Free text used to describe an event which occurred in a system.
	EventDescription *Max350Text `xml:"EvtDesc,omitempty"`

	// Date and time at which the event occurred.
	EventTime *ISODateTime `xml:"EvtTm,omitempty"`
}

func (e *Event1) SetEventCode(value string) {
	e.EventCode = (*Max4AlphaNumericText)(&value)
}

func (e *Event1) AddEventParameter(value string) {
	e.EventParameter = append(e.EventParameter, (*Max35Text)(&value))
}

func (e *Event1) SetEventDescription(value string) {
	e.EventDescription = (*Max350Text)(&value)
}

func (e *Event1) SetEventTime(value string) {
	e.EventTime = (*ISODateTime)(&value)
}
