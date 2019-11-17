package model

// Provides information about event of a corporate action.
type EventInformation1 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Provides the reference of the linked official corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType3Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Provides information about the identification of the last notification.
	LastNotificationIdentification *NotificationIdentification1 `xml:"LastNtfctnId,omitempty"`
}

func (e *EventInformation1) SetCorporateActionEventIdentification(value string) {
	e.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (e *EventInformation1) SetOfficialCorporateActionEventIdentification(value string) {
	e.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (e *EventInformation1) AddEventType() *CorporateActionEventType3Choice {
	e.EventType = new(CorporateActionEventType3Choice)
	return e.EventType
}

func (e *EventInformation1) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	e.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return e.MandatoryVoluntaryEventType
}

func (e *EventInformation1) AddLastNotificationIdentification() *NotificationIdentification1 {
	e.LastNotificationIdentification = new(NotificationIdentification1)
	return e.LastNotificationIdentification
}
