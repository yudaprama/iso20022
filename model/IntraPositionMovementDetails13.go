package model

// Details of the intra-position movement.
type IntraPositionMovementDetails13 struct {

	// Identifications (account owner and/or account servicer) of the intra-position movement.
	Identification *References42Choice `xml:"Id,omitempty"`

	// Quantity of financial instrument effectively settled.
	SettledQuantity *FinancialInstrumentQuantity1Choice `xml:"SttldQty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity1Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity1Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Number identifying a securities sub-balance type (example restriction identification etc…).
	SecuritiesSubBalanceIdentification *GenericIdentification37 `xml:"SctiesSubBalId,omitempty"`

	// Balance to which the securities were moved.
	BalanceTo *SecuritiesBalanceType6Choice `xml:"BalTo"`

	// Date and time at which the securities were moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Date/time securities become available for sale (if securities become unavailable, this specifies the date/time at which they will become available again).
	AvailableDate *DateAndDateTimeChoice `xml:"AvlblDt,omitempty"`

	// Time stamp on when the transaction is acknowledged.
	AcknowledgedStatusTimeStamp *ISODateTime `xml:"AckdStsTmStmp,omitempty"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType56Choice `xml:"CorpActnEvtTp,omitempty"`

	// Value of the collateral available for the delivery settlement process at the account level.
	CollateralMonitorAmount *AmountAndDirection44 `xml:"CollMntrAmt,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *Max350Text `xml:"InstrPrcgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *IntraPositionMovementDetails13) AddIdentification() *References42Choice {
	i.Identification = new(References42Choice)
	return i.Identification
}

func (i *IntraPositionMovementDetails13) AddSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettledQuantity
}

func (i *IntraPositionMovementDetails13) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.PreviouslySettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.PreviouslySettledQuantity
}

func (i *IntraPositionMovementDetails13) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.RemainingToBeSettledQuantity
}

func (i *IntraPositionMovementDetails13) AddSecuritiesSubBalanceIdentification() *GenericIdentification37 {
	i.SecuritiesSubBalanceIdentification = new(GenericIdentification37)
	return i.SecuritiesSubBalanceIdentification
}

func (i *IntraPositionMovementDetails13) AddBalanceTo() *SecuritiesBalanceType6Choice {
	i.BalanceTo = new(SecuritiesBalanceType6Choice)
	return i.BalanceTo
}

func (i *IntraPositionMovementDetails13) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionMovementDetails13) AddAvailableDate() *DateAndDateTimeChoice {
	i.AvailableDate = new(DateAndDateTimeChoice)
	return i.AvailableDate
}

func (i *IntraPositionMovementDetails13) SetAcknowledgedStatusTimeStamp(value string) {
	i.AcknowledgedStatusTimeStamp = (*ISODateTime)(&value)
}

func (i *IntraPositionMovementDetails13) AddCorporateActionEventType() *CorporateActionEventType56Choice {
	i.CorporateActionEventType = new(CorporateActionEventType56Choice)
	return i.CorporateActionEventType
}

func (i *IntraPositionMovementDetails13) AddCollateralMonitorAmount() *AmountAndDirection44 {
	i.CollateralMonitorAmount = new(AmountAndDirection44)
	return i.CollateralMonitorAmount
}

func (i *IntraPositionMovementDetails13) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

func (i *IntraPositionMovementDetails13) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
