package model

// Provides information about event of a corporate action.
type EventInformation9 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Provides the reference of the linked official corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType52Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary3Choice `xml:"MndtryVlntryEvtTp"`

	// Provides information about the identification of the last notification.
	LastNotificationIdentification *NotificationIdentification3 `xml:"LastNtfctnId,omitempty"`
}

func (e *EventInformation9) SetCorporateActionEventIdentification(value string) {
	e.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (e *EventInformation9) SetOfficialCorporateActionEventIdentification(value string) {
	e.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (e *EventInformation9) AddEventType() *CorporateActionEventType52Choice {
	e.EventType = new(CorporateActionEventType52Choice)
	return e.EventType
}

func (e *EventInformation9) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary3Choice {
	e.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary3Choice)
	return e.MandatoryVoluntaryEventType
}

func (e *EventInformation9) AddLastNotificationIdentification() *NotificationIdentification3 {
	e.LastNotificationIdentification = new(NotificationIdentification3)
	return e.LastNotificationIdentification
}
