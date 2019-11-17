package model

// General information about the corporate action event.
type CorporateActionGeneralInformation123 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *RestrictedFINXMax16Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType61Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary4Choice `xml:"MndtryVlntryEvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`
}

func (c *CorporateActionGeneralInformation123) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation123) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation123) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation123) AddEventType() *CorporateActionEventType61Choice {
	c.EventType = new(CorporateActionEventType61Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation123) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary4Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary4Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation123) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return c.FinancialInstrumentIdentification
}
