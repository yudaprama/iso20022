package model

// General information about the corporate action event.
type CorporateActionGeneralInformation97 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType36Choice `xml:"EvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes69 `xml:"UndrlygScty,omitempty"`
}

func (c *CorporateActionGeneralInformation97) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation97) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation97) AddEventType() *CorporateActionEventType36Choice {
	c.EventType = new(CorporateActionEventType36Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation97) AddUnderlyingSecurity() *FinancialInstrumentAttributes69 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes69)
	return c.UnderlyingSecurity
}
