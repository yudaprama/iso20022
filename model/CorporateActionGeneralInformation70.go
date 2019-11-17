package model

// General information about the corporate action event.
type CorporateActionGeneralInformation70 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingTypeChoice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType13Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes48 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation70) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation70) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation70) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation70) AddEventProcessingType() *CorporateActionEventProcessingTypeChoice {
	c.EventProcessingType = new(CorporateActionEventProcessingTypeChoice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation70) AddEventType() *CorporateActionEventType13Choice {
	c.EventType = new(CorporateActionEventType13Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation70) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation70) AddUnderlyingSecurity() *FinancialInstrumentAttributes48 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes48)
	return c.UnderlyingSecurity
}
