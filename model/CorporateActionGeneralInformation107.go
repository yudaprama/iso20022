package model

// General information about the corporate action event.
type CorporateActionGeneralInformation107 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType52Choice `xml:"EvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes81 `xml:"UndrlygScty,omitempty"`
}

func (c *CorporateActionGeneralInformation107) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation107) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation107) AddEventType() *CorporateActionEventType52Choice {
	c.EventType = new(CorporateActionEventType52Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation107) AddUnderlyingSecurity() *FinancialInstrumentAttributes81 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes81)
	return c.UnderlyingSecurity
}
