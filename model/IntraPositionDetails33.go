package model

// Details of the intra-position movement.
type IntraPositionDetails33 struct {

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric4Choice `xml:"Prty,omitempty"`

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Number identifying a Securities Sub balance Type (e.g. restriction identification etc…).
	SecuritiesSubBalanceIdentification *GenericIdentification37 `xml:"SctiesSubBalId,omitempty"`

	// Date and time at which the securities are to be moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesSubBalanceTypeAndQuantityBreakdown3 `xml:"BalFr"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesSubBalanceTypeAndQuantityBreakdown3 `xml:"BalTo"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *Max350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (i *IntraPositionDetails33) AddPriority() *PriorityNumeric4Choice {
	i.Priority = new(PriorityNumeric4Choice)
	return i.Priority
}

func (i *IntraPositionDetails33) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettlementQuantity
}

func (i *IntraPositionDetails33) AddSecuritiesSubBalanceIdentification() *GenericIdentification37 {
	i.SecuritiesSubBalanceIdentification = new(GenericIdentification37)
	return i.SecuritiesSubBalanceIdentification
}

func (i *IntraPositionDetails33) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails33) AddBalanceFrom() *SecuritiesSubBalanceTypeAndQuantityBreakdown3 {
	i.BalanceFrom = new(SecuritiesSubBalanceTypeAndQuantityBreakdown3)
	return i.BalanceFrom
}

func (i *IntraPositionDetails33) AddBalanceTo() *SecuritiesSubBalanceTypeAndQuantityBreakdown3 {
	i.BalanceTo = new(SecuritiesSubBalanceTypeAndQuantityBreakdown3)
	return i.BalanceTo
}

func (i *IntraPositionDetails33) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}
