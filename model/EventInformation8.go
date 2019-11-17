package model

// Provides information about event of a corporate action.
type EventInformation8 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Provides the reference of the linked official corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType36Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary4Choice `xml:"MndtryVlntryEvtTp"`

	// Provides information about the identification of the last notification.
	LastNotificationIdentification *NotificationIdentification4 `xml:"LastNtfctnId,omitempty"`
}

func (e *EventInformation8) SetCorporateActionEventIdentification(value string) {
	e.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (e *EventInformation8) SetOfficialCorporateActionEventIdentification(value string) {
	e.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (e *EventInformation8) AddEventType() *CorporateActionEventType36Choice {
	e.EventType = new(CorporateActionEventType36Choice)
	return e.EventType
}

func (e *EventInformation8) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary4Choice {
	e.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary4Choice)
	return e.MandatoryVoluntaryEventType
}

func (e *EventInformation8) AddLastNotificationIdentification() *NotificationIdentification4 {
	e.LastNotificationIdentification = new(NotificationIdentification4)
	return e.LastNotificationIdentification
}
