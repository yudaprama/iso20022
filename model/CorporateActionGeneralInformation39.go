package model

// General information about the corporate action event.
type CorporateActionGeneralInformation39 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType8Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Indicates that the additional business process relates to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat7Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat6Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`
}

func (c *CorporateActionGeneralInformation39) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation39) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation39) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation39) AddEventType() *CorporateActionEventType8Choice {
	c.EventType = new(CorporateActionEventType8Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation39) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionGeneralInformation39) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat7Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat7Choice)
	return c.AdditionalBusinessProcessIndicator
}

func (c *CorporateActionGeneralInformation39) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat6Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat6Choice)
	return c.IntermediateSecuritiesDistributionType
}
