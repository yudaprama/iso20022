package model

// Set of elements used to provide reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransactionInformation30 struct {

	// Unique identification, as assigned by the assigner, to unambiguously identify a cancellation request.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationIdentification *Max35Text `xml:"CxlId,omitempty"`

	// Set of elements to uniquely and unambiguously identify an exception or an investigation workflow.
	Case *Case2 `xml:"Case,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Date at which the initiating party originally requested the clearing agent to process the payment.
	OriginalRequestedExecutionDate *ISODate `xml:"OrgnlReqdExctnDt,omitempty"`

	// Date at which the creditor originally requested the collection of the amount of money from the debtor.
	OriginalRequestedCollectionDate *ISODate `xml:"OrgnlReqdColltnDt,omitempty"`

	// Set of elements used to provide detailed information on the cancellation reason.
	CancellationReasonInformation []*CancellationReasonInformation3 `xml:"CxlRsnInf,omitempty"`

	// Set of key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference13 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation30) SetCancellationIdentification(value string) {
	p.CancellationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation30) AddCase() *Case2 {
	p.Case = new(Case2)
	return p.Case
}

func (p *PaymentTransactionInformation30) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation30) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation30) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation30) SetOriginalRequestedExecutionDate(value string) {
	p.OriginalRequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentTransactionInformation30) SetOriginalRequestedCollectionDate(value string) {
	p.OriginalRequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentTransactionInformation30) AddCancellationReasonInformation() *CancellationReasonInformation3 {
	newValue := new(CancellationReasonInformation3)
	p.CancellationReasonInformation = append(p.CancellationReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation30) AddOriginalTransactionReference() *OriginalTransactionReference13 {
	p.OriginalTransactionReference = new(OriginalTransactionReference13)
	return p.OriginalTransactionReference
}
