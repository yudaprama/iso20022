package model

// Set of elements used to provide reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction74 struct {

	// Unique identification, as assigned by the assigner, to unambiguously identify a cancellation request.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationIdentification *Max35Text `xml:"CxlId,omitempty"`

	// Set of elements to uniquely and unambiguously identify an exception or an investigation workflow.
	Case *Case3 `xml:"Case,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	//
	// Usage: This amount has to be transported unchanged through the transaction chain.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Date at which the initiating party originally requested the clearing agent to process the payment.
	OriginalRequestedExecutionDate *DateAndDateTimeChoice `xml:"OrgnlReqdExctnDt,omitempty"`

	// Date at which the creditor originally requested the collection of the amount of money from the debtor.
	OriginalRequestedCollectionDate *ISODate `xml:"OrgnlReqdColltnDt,omitempty"`

	// Provides detailed information on the cancellation reason.
	CancellationReasonInformation []*PaymentCancellationReason3 `xml:"CxlRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference24 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction74) SetCancellationIdentification(value string) {
	p.CancellationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction74) AddCase() *Case3 {
	p.Case = new(Case3)
	return p.Case
}

func (p *PaymentTransaction74) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction74) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction74) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction74) AddOriginalRequestedExecutionDate() *DateAndDateTimeChoice {
	p.OriginalRequestedExecutionDate = new(DateAndDateTimeChoice)
	return p.OriginalRequestedExecutionDate
}

func (p *PaymentTransaction74) SetOriginalRequestedCollectionDate(value string) {
	p.OriginalRequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentTransaction74) AddCancellationReasonInformation() *PaymentCancellationReason3 {
	newValue := new(PaymentCancellationReason3)
	p.CancellationReasonInformation = append(p.CancellationReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction74) AddOriginalTransactionReference() *OriginalTransactionReference24 {
	p.OriginalTransactionReference = new(OriginalTransactionReference24)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction74) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}
