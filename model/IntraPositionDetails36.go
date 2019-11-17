package model

// Details of the intra-position movement.
type IntraPositionDetails36 struct {

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric5Choice `xml:"Prty,omitempty"`

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity15Choice `xml:"SttlmQty"`

	// Number identifying a Securities Sub balance Type (e.g. restriction identification etc…).
	SecuritiesSubBalanceIdentification *GenericIdentification39 `xml:"SctiesSubBalId,omitempty"`

	// Date and time at which the securities are to be moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesSubBalanceTypeAndQuantityBreakdown4 `xml:"BalFr"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesSubBalanceTypeAndQuantityBreakdown4 `xml:"BalTo"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (i *IntraPositionDetails36) AddPriority() *PriorityNumeric5Choice {
	i.Priority = new(PriorityNumeric5Choice)
	return i.Priority
}

func (i *IntraPositionDetails36) AddSettlementQuantity() *FinancialInstrumentQuantity15Choice {
	i.SettlementQuantity = new(FinancialInstrumentQuantity15Choice)
	return i.SettlementQuantity
}

func (i *IntraPositionDetails36) AddSecuritiesSubBalanceIdentification() *GenericIdentification39 {
	i.SecuritiesSubBalanceIdentification = new(GenericIdentification39)
	return i.SecuritiesSubBalanceIdentification
}

func (i *IntraPositionDetails36) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails36) AddBalanceFrom() *SecuritiesSubBalanceTypeAndQuantityBreakdown4 {
	i.BalanceFrom = new(SecuritiesSubBalanceTypeAndQuantityBreakdown4)
	return i.BalanceFrom
}

func (i *IntraPositionDetails36) AddBalanceTo() *SecuritiesSubBalanceTypeAndQuantityBreakdown4 {
	i.BalanceTo = new(SecuritiesSubBalanceTypeAndQuantityBreakdown4)
	return i.BalanceTo
}

func (i *IntraPositionDetails36) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}
