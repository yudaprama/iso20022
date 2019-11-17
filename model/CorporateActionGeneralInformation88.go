package model

// General information about the corporate action event.
type CorporateActionGeneralInformation88 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType32Choice `xml:"EvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes65 `xml:"UndrlygScty,omitempty"`
}

func (c *CorporateActionGeneralInformation88) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation88) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation88) AddEventType() *CorporateActionEventType32Choice {
	c.EventType = new(CorporateActionEventType32Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation88) AddUnderlyingSecurity() *FinancialInstrumentAttributes65 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes65)
	return c.UnderlyingSecurity
}
