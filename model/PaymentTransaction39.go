package model

// Provides reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction39 struct {

	// Unique and unambiguous identifier of a cancellation request status, as assigned by the assigner.
	//
	// Usage: The cancellation status identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationStatusIdentification *Max35Text `xml:"CxlStsId,omitempty"`

	// Identifies the resolved case.
	ResolvedCase *Case3 `xml:"RslvdCase,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Specifies the status of the transaction cancellation request.
	TransactionCancellationStatus *CancellationIndividualStatus1Code `xml:"TxCxlSts,omitempty"`

	// Provides detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReason2 `xml:"CxlStsRsnInf,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Date at which the initiating party originally requested the clearing agent to process the payment.
	OriginalRequestedExecutionDate *ISODate `xml:"OrgnlReqdExctnDt,omitempty"`

	// Date at which the creditor originally requested the collection of the amount of money from the debtor.
	OriginalRequestedCollectionDate *ISODate `xml:"OrgnlReqdColltnDt,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference16 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransaction39) SetCancellationStatusIdentification(value string) {
	p.CancellationStatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction39) AddResolvedCase() *Case3 {
	p.ResolvedCase = new(Case3)
	return p.ResolvedCase
}

func (p *PaymentTransaction39) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction39) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction39) SetTransactionCancellationStatus(value string) {
	p.TransactionCancellationStatus = (*CancellationIndividualStatus1Code)(&value)
}

func (p *PaymentTransaction39) AddCancellationStatusReasonInformation() *CancellationStatusReason2 {
	newValue := new(CancellationStatusReason2)
	p.CancellationStatusReasonInformation = append(p.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction39) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction39) SetOriginalRequestedExecutionDate(value string) {
	p.OriginalRequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentTransaction39) SetOriginalRequestedCollectionDate(value string) {
	p.OriginalRequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentTransaction39) AddOriginalTransactionReference() *OriginalTransactionReference16 {
	p.OriginalTransactionReference = new(OriginalTransactionReference16)
	return p.OriginalTransactionReference
}
