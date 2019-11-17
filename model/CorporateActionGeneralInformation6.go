package model

// General information about the corporate action event.
type CorporateActionGeneralInformation6 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType3Choice `xml:"EvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes6 `xml:"UndrlygScty,omitempty"`
}

func (c *CorporateActionGeneralInformation6) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation6) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation6) AddEventType() *CorporateActionEventType3Choice {
	c.EventType = new(CorporateActionEventType3Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation6) AddUnderlyingSecurity() *FinancialInstrumentAttributes6 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes6)
	return c.UnderlyingSecurity
}
