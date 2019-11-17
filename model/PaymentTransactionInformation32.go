package model

// Set of elements used to provide reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransactionInformation32 struct {

	// Unique and unambiguous identifier of a cancellation request status, as assigned by the assigner.
	//
	// Usage: The cancellation status identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationStatusIdentification *Max35Text `xml:"CxlStsId,omitempty"`

	// Identifies the resolved case.
	ResolvedCase *Case2 `xml:"RslvdCase,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Specifies the status of the transaction cancellation request.
	TransactionCancellationStatus *CancellationIndividualStatus1Code `xml:"TxCxlSts,omitempty"`

	// Set of elements used to provide detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReasonInformation1 `xml:"CxlStsRsnInf,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Date at which the initiating party originally requested the clearing agent to process the payment.
	OriginalRequestedExecutionDate *ISODate `xml:"OrgnlReqdExctnDt,omitempty"`

	// Date at which the creditor originally requested the collection of the amount of money from the debtor.
	OriginalRequestedCollectionDate *ISODate `xml:"OrgnlReqdColltnDt,omitempty"`

	// Set of key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference13 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation32) SetCancellationStatusIdentification(value string) {
	p.CancellationStatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation32) AddResolvedCase() *Case2 {
	p.ResolvedCase = new(Case2)
	return p.ResolvedCase
}

func (p *PaymentTransactionInformation32) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation32) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation32) SetTransactionCancellationStatus(value string) {
	p.TransactionCancellationStatus = (*CancellationIndividualStatus1Code)(&value)
}

func (p *PaymentTransactionInformation32) AddCancellationStatusReasonInformation() *CancellationStatusReasonInformation1 {
	newValue := new(CancellationStatusReasonInformation1)
	p.CancellationStatusReasonInformation = append(p.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation32) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation32) SetOriginalRequestedExecutionDate(value string) {
	p.OriginalRequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentTransactionInformation32) SetOriginalRequestedCollectionDate(value string) {
	p.OriginalRequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentTransactionInformation32) AddOriginalTransactionReference() *OriginalTransactionReference13 {
	p.OriginalTransactionReference = new(OriginalTransactionReference13)
	return p.OriginalTransactionReference
}
