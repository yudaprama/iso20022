package model

// General information about the corporate action event.
type CorporateActionGeneralInformation40 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Specifies the type of narrative related to the message.
	NarrativeType *CorporateActionNarrative1Choice `xml:"NrrtvTp,omitempty"`
}

func (c *CorporateActionGeneralInformation40) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation40) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation40) AddNarrativeType() *CorporateActionNarrative1Choice {
	c.NarrativeType = new(CorporateActionNarrative1Choice)
	return c.NarrativeType
}
