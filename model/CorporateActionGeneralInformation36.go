package model

// General information about the corporate action event.
type CorporateActionGeneralInformation36 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType7Choice `xml:"EvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes32 `xml:"UndrlygScty,omitempty"`
}

func (c *CorporateActionGeneralInformation36) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation36) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation36) AddEventType() *CorporateActionEventType7Choice {
	c.EventType = new(CorporateActionEventType7Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation36) AddUnderlyingSecurity() *FinancialInstrumentAttributes32 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes32)
	return c.UnderlyingSecurity
}
