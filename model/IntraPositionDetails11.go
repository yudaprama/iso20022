package model

// Details of the intra-position movement.
type IntraPositionDetails11 struct {

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Number identifying a lot constituting the financial instrument.
	LotNumber *Number2Choice `xml:"LotNb,omitempty"`

	// Date and time at which the securities are to be moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesBalanceType2Choice `xml:"BalFr"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesBalanceType2Choice `xml:"BalTo"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *Max350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (i *IntraPositionDetails11) AddPriority() *PriorityNumeric1Choice {
	i.Priority = new(PriorityNumeric1Choice)
	return i.Priority
}

func (i *IntraPositionDetails11) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettlementQuantity
}

func (i *IntraPositionDetails11) AddLotNumber() *Number2Choice {
	i.LotNumber = new(Number2Choice)
	return i.LotNumber
}

func (i *IntraPositionDetails11) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails11) AddBalanceFrom() *SecuritiesBalanceType2Choice {
	i.BalanceFrom = new(SecuritiesBalanceType2Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails11) AddBalanceTo() *SecuritiesBalanceType2Choice {
	i.BalanceTo = new(SecuritiesBalanceType2Choice)
	return i.BalanceTo
}

func (i *IntraPositionDetails11) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}
