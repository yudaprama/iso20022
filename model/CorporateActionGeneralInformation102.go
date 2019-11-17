package model

// General information about the corporate action event.
type CorporateActionGeneralInformation102 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Specifies the type of narrative related to the message.
	NarrativeType *CorporateActionNarrative4Choice `xml:"NrrtvTp,omitempty"`
}

func (c *CorporateActionGeneralInformation102) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation102) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation102) AddNarrativeType() *CorporateActionNarrative4Choice {
	c.NarrativeType = new(CorporateActionNarrative4Choice)
	return c.NarrativeType
}
