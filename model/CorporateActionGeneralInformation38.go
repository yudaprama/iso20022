package model

// General information about the corporate action event.
type CorporateActionGeneralInformation38 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingTypeChoice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType7Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Indicates whether the message is related to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat2Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes33 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation38) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation38) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation38) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation38) AddEventProcessingType() *CorporateActionEventProcessingTypeChoice {
	c.EventProcessingType = new(CorporateActionEventProcessingTypeChoice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation38) AddEventType() *CorporateActionEventType7Choice {
	c.EventType = new(CorporateActionEventType7Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation38) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation38) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat2Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat2Choice)
	return c.AdditionalBusinessProcessIndicator
}

func (c *CorporateActionGeneralInformation38) AddUnderlyingSecurity() *FinancialInstrumentAttributes33 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes33)
	return c.UnderlyingSecurity
}
