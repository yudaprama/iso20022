package model

// General information about the corporate action event.
type CorporateActionGeneralInformation114 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *RestrictedFINXMax16Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType58Choice `xml:"EvtTp"`
}

func (c *CorporateActionGeneralInformation114) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation114) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation114) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation114) AddEventType() *CorporateActionEventType58Choice {
	c.EventType = new(CorporateActionEventType58Choice)
	return c.EventType
}
