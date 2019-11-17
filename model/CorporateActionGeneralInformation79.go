package model

// General information about the corporate action event.
type CorporateActionGeneralInformation79 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType12Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Indicates that the additional business process relates to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat7Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat5Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Fractional quantity resulting from an event which will be paid with cash in lieu.
	FractionalQuantity *FinancialInstrumentQuantity1Choice `xml:"FrctnlQty,omitempty"`
}

func (c *CorporateActionGeneralInformation79) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation79) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation79) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation79) AddEventType() *CorporateActionEventType12Choice {
	c.EventType = new(CorporateActionEventType12Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation79) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionGeneralInformation79) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat7Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat7Choice)
	return c.AdditionalBusinessProcessIndicator
}

func (c *CorporateActionGeneralInformation79) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat5Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat5Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionGeneralInformation79) AddFractionalQuantity() *FinancialInstrumentQuantity1Choice {
	c.FractionalQuantity = new(FinancialInstrumentQuantity1Choice)
	return c.FractionalQuantity
}
