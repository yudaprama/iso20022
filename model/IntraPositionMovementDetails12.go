package model

// Details of the intra-position movement.
type IntraPositionMovementDetails12 struct {

	// Identifications (account owner and/or account servicer) of the intra-position movement.
	Identification *References51Choice `xml:"Id,omitempty"`

	// Quantity of financial instrument effectively settled.
	SettledQuantity *FinancialInstrumentQuantity15Choice `xml:"SttldQty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity15Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity15Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Number identifying a securities sub-balance type (example restriction identification etc…).
	SecuritiesSubBalanceIdentification *GenericIdentification39 `xml:"SctiesSubBalId,omitempty"`

	// Balance to which the securities were moved.
	BalanceTo *SecuritiesBalanceType8Choice `xml:"BalTo"`

	// Date and time at which the securities were moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Date/time securities become available for sale (if securities become unavailable, this specifies the date/time at which they will become available again).
	AvailableDate *DateAndDateTimeChoice `xml:"AvlblDt,omitempty"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType46Choice `xml:"CorpActnEvtTp,omitempty"`

	// Value of the collateral available for the delivery settlement process at the account level.
	CollateralMonitorAmount *AmountAndDirection55 `xml:"CollMntrAmt,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"InstrPrcgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *IntraPositionMovementDetails12) AddIdentification() *References51Choice {
	i.Identification = new(References51Choice)
	return i.Identification
}

func (i *IntraPositionMovementDetails12) AddSettledQuantity() *FinancialInstrumentQuantity15Choice {
	i.SettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return i.SettledQuantity
}

func (i *IntraPositionMovementDetails12) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity15Choice {
	i.PreviouslySettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return i.PreviouslySettledQuantity
}

func (i *IntraPositionMovementDetails12) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity15Choice {
	i.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return i.RemainingToBeSettledQuantity
}

func (i *IntraPositionMovementDetails12) AddSecuritiesSubBalanceIdentification() *GenericIdentification39 {
	i.SecuritiesSubBalanceIdentification = new(GenericIdentification39)
	return i.SecuritiesSubBalanceIdentification
}

func (i *IntraPositionMovementDetails12) AddBalanceTo() *SecuritiesBalanceType8Choice {
	i.BalanceTo = new(SecuritiesBalanceType8Choice)
	return i.BalanceTo
}

func (i *IntraPositionMovementDetails12) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionMovementDetails12) AddAvailableDate() *DateAndDateTimeChoice {
	i.AvailableDate = new(DateAndDateTimeChoice)
	return i.AvailableDate
}

func (i *IntraPositionMovementDetails12) AddCorporateActionEventType() *CorporateActionEventType46Choice {
	i.CorporateActionEventType = new(CorporateActionEventType46Choice)
	return i.CorporateActionEventType
}

func (i *IntraPositionMovementDetails12) AddCollateralMonitorAmount() *AmountAndDirection55 {
	i.CollateralMonitorAmount = new(AmountAndDirection55)
	return i.CollateralMonitorAmount
}

func (i *IntraPositionMovementDetails12) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

func (i *IntraPositionMovementDetails12) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
