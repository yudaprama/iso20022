package model

// General information about the corporate action event.
type CorporateActionGeneralInformation55 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType11Choice `xml:"EvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes32 `xml:"UndrlygScty,omitempty"`
}

func (c *CorporateActionGeneralInformation55) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation55) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation55) AddEventType() *CorporateActionEventType11Choice {
	c.EventType = new(CorporateActionEventType11Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation55) AddUnderlyingSecurity() *FinancialInstrumentAttributes32 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes32)
	return c.UnderlyingSecurity
}
