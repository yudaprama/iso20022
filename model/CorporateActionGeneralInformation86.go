package model

// General information about the corporate action event.
type CorporateActionGeneralInformation86 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType33Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary3Choice `xml:"MndtryVlntryEvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`
}

func (c *CorporateActionGeneralInformation86) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation86) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation86) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation86) AddEventType() *CorporateActionEventType33Choice {
	c.EventType = new(CorporateActionEventType33Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation86) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary3Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary3Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation86) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return c.FinancialInstrumentIdentification
}
