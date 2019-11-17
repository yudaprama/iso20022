package model

// Details of the intra-position movement.
type IntraPositionMovementDetails14 struct {

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

	// Time stamp on when the transaction is acknowledged.
	AcknowledgedStatusTimeStamp *ISODateTime `xml:"AckdStsTmStmp,omitempty"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType69Choice `xml:"CorpActnEvtTp,omitempty"`

	// Value of the collateral available for the delivery settlement process at the account level.
	CollateralMonitorAmount *AmountAndDirection55 `xml:"CollMntrAmt,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"InstrPrcgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *IntraPositionMovementDetails14) AddIdentification() *References51Choice {
	i.Identification = new(References51Choice)
	return i.Identification
}

func (i *IntraPositionMovementDetails14) AddSettledQuantity() *FinancialInstrumentQuantity15Choice {
	i.SettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return i.SettledQuantity
}

func (i *IntraPositionMovementDetails14) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity15Choice {
	i.PreviouslySettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return i.PreviouslySettledQuantity
}

func (i *IntraPositionMovementDetails14) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity15Choice {
	i.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return i.RemainingToBeSettledQuantity
}

func (i *IntraPositionMovementDetails14) AddSecuritiesSubBalanceIdentification() *GenericIdentification39 {
	i.SecuritiesSubBalanceIdentification = new(GenericIdentification39)
	return i.SecuritiesSubBalanceIdentification
}

func (i *IntraPositionMovementDetails14) AddBalanceTo() *SecuritiesBalanceType8Choice {
	i.BalanceTo = new(SecuritiesBalanceType8Choice)
	return i.BalanceTo
}

func (i *IntraPositionMovementDetails14) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionMovementDetails14) AddAvailableDate() *DateAndDateTimeChoice {
	i.AvailableDate = new(DateAndDateTimeChoice)
	return i.AvailableDate
}

func (i *IntraPositionMovementDetails14) SetAcknowledgedStatusTimeStamp(value string) {
	i.AcknowledgedStatusTimeStamp = (*ISODateTime)(&value)
}

func (i *IntraPositionMovementDetails14) AddCorporateActionEventType() *CorporateActionEventType69Choice {
	i.CorporateActionEventType = new(CorporateActionEventType69Choice)
	return i.CorporateActionEventType
}

func (i *IntraPositionMovementDetails14) AddCollateralMonitorAmount() *AmountAndDirection55 {
	i.CollateralMonitorAmount = new(AmountAndDirection55)
	return i.CollateralMonitorAmount
}

func (i *IntraPositionMovementDetails14) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

func (i *IntraPositionMovementDetails14) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
