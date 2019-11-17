package model

// General information about the corporate action event.
type CorporateActionGeneralInformation105 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingType2Choice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType51Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary3Choice `xml:"MndtryVlntryEvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes79 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation105) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation105) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation105) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation105) AddEventProcessingType() *CorporateActionEventProcessingType2Choice {
	c.EventProcessingType = new(CorporateActionEventProcessingType2Choice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation105) AddEventType() *CorporateActionEventType51Choice {
	c.EventType = new(CorporateActionEventType51Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation105) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary3Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary3Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation105) AddUnderlyingSecurity() *FinancialInstrumentAttributes79 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes79)
	return c.UnderlyingSecurity
}
