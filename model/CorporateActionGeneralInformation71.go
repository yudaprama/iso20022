package model

// General information about the corporate action event.
type CorporateActionGeneralInformation71 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType11Choice `xml:"EvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes50 `xml:"UndrlygScty,omitempty"`
}

func (c *CorporateActionGeneralInformation71) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation71) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation71) AddEventType() *CorporateActionEventType11Choice {
	c.EventType = new(CorporateActionEventType11Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation71) AddUnderlyingSecurity() *FinancialInstrumentAttributes50 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes50)
	return c.UnderlyingSecurity
}
