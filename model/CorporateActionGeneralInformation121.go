package model

// General information about the corporate action event.
type CorporateActionGeneralInformation121 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType58Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification21 `xml:"FinInstrmId,omitempty"`
}

func (c *CorporateActionGeneralInformation121) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation121) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation121) AddEventType() *CorporateActionEventType58Choice {
	c.EventType = new(CorporateActionEventType58Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation121) AddFinancialInstrumentIdentification() *SecurityIdentification21 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification21)
	return c.FinancialInstrumentIdentification
}
