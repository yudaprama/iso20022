package model

// General information about the corporate action event.
type CorporateActionGeneralInformation24 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType5Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	UnderlyingSecurityIdentification *SecurityIdentification14 `xml:"UndrlygSctyId"`

	// Indicates that the additional business process relates to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat3Choice `xml:"AddtlBizPrcInd,omitempty"`
}

func (c *CorporateActionGeneralInformation24) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation24) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation24) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation24) AddEventType() *CorporateActionEventType5Choice {
	c.EventType = new(CorporateActionEventType5Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation24) AddUnderlyingSecurityIdentification() *SecurityIdentification14 {
	c.UnderlyingSecurityIdentification = new(SecurityIdentification14)
	return c.UnderlyingSecurityIdentification
}

func (c *CorporateActionGeneralInformation24) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat3Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat3Choice)
	return c.AdditionalBusinessProcessIndicator
}
