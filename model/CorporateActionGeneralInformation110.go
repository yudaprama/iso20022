package model

// General information about the corporate action event.
type CorporateActionGeneralInformation110 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType52Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId,omitempty"`
}

func (c *CorporateActionGeneralInformation110) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation110) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation110) AddEventType() *CorporateActionEventType52Choice {
	c.EventType = new(CorporateActionEventType52Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation110) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return c.FinancialInstrumentIdentification
}
