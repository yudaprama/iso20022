package model

// Provides information about event of a corporate action.
type EventInformation3 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Provides the reference of the linked official corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType7Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Provides information about the identification of the last notification.
	LastNotificationIdentification *NotificationIdentification1 `xml:"LastNtfctnId,omitempty"`
}

func (e *EventInformation3) SetCorporateActionEventIdentification(value string) {
	e.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (e *EventInformation3) SetOfficialCorporateActionEventIdentification(value string) {
	e.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (e *EventInformation3) AddEventType() *CorporateActionEventType7Choice {
	e.EventType = new(CorporateActionEventType7Choice)
	return e.EventType
}

func (e *EventInformation3) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	e.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return e.MandatoryVoluntaryEventType
}

func (e *EventInformation3) AddLastNotificationIdentification() *NotificationIdentification1 {
	e.LastNotificationIdentification = new(NotificationIdentification1)
	return e.LastNotificationIdentification
}
