package model

// General information about the corporate action event.
type CorporateActionGeneralInformation53 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType14Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`
}

func (c *CorporateActionGeneralInformation53) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation53) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation53) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation53) AddEventType() *CorporateActionEventType14Choice {
	c.EventType = new(CorporateActionEventType14Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation53) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation53) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return c.FinancialInstrumentIdentification
}
