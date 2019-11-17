package model

// Details of the intra-position movement.
type IntraPositionDetails43 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *FinancialInstrumentQuantity15Choice `xml:"SttldQty"`

	// Number identifying a Securities Sub balance Type (example restriction identification etc…).
	SecuritiesSubBalanceIdentification *GenericIdentification39 `xml:"SctiesSubBalId,omitempty"`

	// Value of the collateral available for the delivery settlement process at the account level.
	CollateralMonitorAmount *AmountAndDirection55 `xml:"CollMntrAmt,omitempty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity15Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity15Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Date and time at which the securities were moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Date/time securities become available for sale (if securities become unavailable, this specifies the date/time at which they will become available again).
	AvailableDate *DateAndDateTimeChoice `xml:"AvlblDt,omitempty"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType69Choice `xml:"CorpActnEvtTp,omitempty"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesSubBalanceTypeAndQuantityBreakdown4 `xml:"BalFr"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesSubBalanceTypeAndQuantityBreakdown4 `xml:"BalTo"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (i *IntraPositionDetails43) AddSettledQuantity() *FinancialInstrumentQuantity15Choice {
	i.SettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return i.SettledQuantity
}

func (i *IntraPositionDetails43) AddSecuritiesSubBalanceIdentification() *GenericIdentification39 {
	i.SecuritiesSubBalanceIdentification = new(GenericIdentification39)
	return i.SecuritiesSubBalanceIdentification
}

func (i *IntraPositionDetails43) AddCollateralMonitorAmount() *AmountAndDirection55 {
	i.CollateralMonitorAmount = new(AmountAndDirection55)
	return i.CollateralMonitorAmount
}

func (i *IntraPositionDetails43) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity15Choice {
	i.PreviouslySettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return i.PreviouslySettledQuantity
}

func (i *IntraPositionDetails43) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity15Choice {
	i.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return i.RemainingToBeSettledQuantity
}

func (i *IntraPositionDetails43) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails43) AddAvailableDate() *DateAndDateTimeChoice {
	i.AvailableDate = new(DateAndDateTimeChoice)
	return i.AvailableDate
}

func (i *IntraPositionDetails43) AddCorporateActionEventType() *CorporateActionEventType69Choice {
	i.CorporateActionEventType = new(CorporateActionEventType69Choice)
	return i.CorporateActionEventType
}

func (i *IntraPositionDetails43) AddBalanceFrom() *SecuritiesSubBalanceTypeAndQuantityBreakdown4 {
	i.BalanceFrom = new(SecuritiesSubBalanceTypeAndQuantityBreakdown4)
	return i.BalanceFrom
}

func (i *IntraPositionDetails43) AddBalanceTo() *SecuritiesSubBalanceTypeAndQuantityBreakdown4 {
	i.BalanceTo = new(SecuritiesSubBalanceTypeAndQuantityBreakdown4)
	return i.BalanceTo
}

func (i *IntraPositionDetails43) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}
