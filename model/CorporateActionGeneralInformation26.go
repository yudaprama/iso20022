package model

// General information about the corporate action event.
type CorporateActionGeneralInformation26 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType3Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	UnderlyingSecurityIdentification *SecurityIdentification14 `xml:"UndrlygSctyId,omitempty"`
}

func (c *CorporateActionGeneralInformation26) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation26) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation26) AddEventType() *CorporateActionEventType3Choice {
	c.EventType = new(CorporateActionEventType3Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation26) AddUnderlyingSecurityIdentification() *SecurityIdentification14 {
	c.UnderlyingSecurityIdentification = new(SecurityIdentification14)
	return c.UnderlyingSecurityIdentification
}
