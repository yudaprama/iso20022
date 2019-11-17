package model

// Set of elements used to provide information on the return or reversal expected by the party that initiated the initial payment instruction after a cancellation or modification request.
type ResolutionInformation1 struct {

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the clearing channel to be used to process the payment instruction.
	ClearingChannel *ClearingChannel2Code `xml:"ClrChanl,omitempty"`
}

func (r *ResolutionInformation1) SetInterbankSettlementAmount(value, currency string) {
	r.InterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *ResolutionInformation1) SetInterbankSettlementDate(value string) {
	r.InterbankSettlementDate = (*ISODate)(&value)
}

func (r *ResolutionInformation1) SetClearingChannel(value string) {
	r.ClearingChannel = (*ClearingChannel2Code)(&value)
}
