package model

// Details of the intra-position movement.
type IntraPositionMovementDetails9 struct {

	// Identifications (account owner and/or account servicer) of the intra-position movement.
	Identification *References27Choice `xml:"Id,omitempty"`

	// Quantity of financial instrument effectively settled.
	SettledQuantity *FinancialInstrumentQuantity1Choice `xml:"SttldQty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity1Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity1Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Number identifying a Securities Sub balance Type (example restriction identification etc…).
	SecuritiesSubBalanceIdentification *GenericIdentification37 `xml:"SctiesSubBalId,omitempty"`

	// Balance to which the securities were moved.
	BalanceTo *SecuritiesBalanceType3Choice `xml:"BalTo"`

	// Date and time at which the securities were moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Date/time securities become available for sale (if securities become unavailable, this specifies the date/time at which they will become available again).
	AvailableDate *DateAndDateTimeChoice `xml:"AvlblDt,omitempty"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType16Choice `xml:"CorpActnEvtTp,omitempty"`

	// Value of the collateral available for the delivery settlement process at the account level.
	CollateralMonitorAmount *AmountAndDirection9 `xml:"CollMntrAmt,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *Max350Text `xml:"InstrPrcgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *IntraPositionMovementDetails9) AddIdentification() *References27Choice {
	i.Identification = new(References27Choice)
	return i.Identification
}

func (i *IntraPositionMovementDetails9) AddSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettledQuantity
}

func (i *IntraPositionMovementDetails9) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.PreviouslySettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.PreviouslySettledQuantity
}

func (i *IntraPositionMovementDetails9) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.RemainingToBeSettledQuantity
}

func (i *IntraPositionMovementDetails9) AddSecuritiesSubBalanceIdentification() *GenericIdentification37 {
	i.SecuritiesSubBalanceIdentification = new(GenericIdentification37)
	return i.SecuritiesSubBalanceIdentification
}

func (i *IntraPositionMovementDetails9) AddBalanceTo() *SecuritiesBalanceType3Choice {
	i.BalanceTo = new(SecuritiesBalanceType3Choice)
	return i.BalanceTo
}

func (i *IntraPositionMovementDetails9) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionMovementDetails9) AddAvailableDate() *DateAndDateTimeChoice {
	i.AvailableDate = new(DateAndDateTimeChoice)
	return i.AvailableDate
}

func (i *IntraPositionMovementDetails9) AddCorporateActionEventType() *CorporateActionEventType16Choice {
	i.CorporateActionEventType = new(CorporateActionEventType16Choice)
	return i.CorporateActionEventType
}

func (i *IntraPositionMovementDetails9) AddCollateralMonitorAmount() *AmountAndDirection9 {
	i.CollateralMonitorAmount = new(AmountAndDirection9)
	return i.CollateralMonitorAmount
}

func (i *IntraPositionMovementDetails9) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

func (i *IntraPositionMovementDetails9) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
