package model

// General information about the corporate action event.
type CorporateActionGeneralInformation113 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *RestrictedFINXMax16Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingType3Choice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType57Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary4Choice `xml:"MndtryVlntryEvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes85 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation113) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation113) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation113) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation113) AddEventProcessingType() *CorporateActionEventProcessingType3Choice {
	c.EventProcessingType = new(CorporateActionEventProcessingType3Choice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation113) AddEventType() *CorporateActionEventType57Choice {
	c.EventType = new(CorporateActionEventType57Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation113) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary4Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary4Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation113) AddUnderlyingSecurity() *FinancialInstrumentAttributes85 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes85)
	return c.UnderlyingSecurity
}
