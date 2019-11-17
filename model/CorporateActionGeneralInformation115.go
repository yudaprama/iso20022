package model

// General information about the corporate action event.
type CorporateActionGeneralInformation115 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType58Choice `xml:"EvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes84 `xml:"UndrlygScty,omitempty"`
}

func (c *CorporateActionGeneralInformation115) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation115) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation115) AddEventType() *CorporateActionEventType58Choice {
	c.EventType = new(CorporateActionEventType58Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation115) AddUnderlyingSecurity() *FinancialInstrumentAttributes84 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes84)
	return c.UnderlyingSecurity
}
