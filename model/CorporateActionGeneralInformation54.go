package model

// General information about the corporate action event.
type CorporateActionGeneralInformation54 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingTypeChoice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType14Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Indicates whether the message is related to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat2Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes43 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation54) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation54) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation54) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation54) AddEventProcessingType() *CorporateActionEventProcessingTypeChoice {
	c.EventProcessingType = new(CorporateActionEventProcessingTypeChoice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation54) AddEventType() *CorporateActionEventType14Choice {
	c.EventType = new(CorporateActionEventType14Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation54) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation54) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat2Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat2Choice)
	return c.AdditionalBusinessProcessIndicator
}

func (c *CorporateActionGeneralInformation54) AddUnderlyingSecurity() *FinancialInstrumentAttributes43 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes43)
	return c.UnderlyingSecurity
}
