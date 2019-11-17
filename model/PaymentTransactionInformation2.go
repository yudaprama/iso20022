package model

// Reference and status information concerning the original transactions, included in the original instruction, to which the return message applies.
type PaymentTransactionInformation2 struct {

	// Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the returned transaction.
	ReturnIdentification *Max35Text `xml:"RtrId,omitempty"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Original unique instruction identification as assigned by an instructing party for an instructed party to unambiguously identify the original instruction.
	//
	// Usage: the original instruction identification is the original point to point reference used between the instructing party and the instructed party to refer to the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Original unique identification assigned by the initiating party to unambiguously identify the original transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Original identification of a transaction, as assigned by the first instructing agent and passed on, unchanged, throughout the entire interbank chain.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Original amount of money as moved between the instructing agent and the instructed agent in the original transaction.
	OriginalInterbankSettlementAmount *CurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Returned amount of money moved between the instructing agent and the instructed agent in the return transaction.
	ReturnedInterbankSettlementAmount *CurrencyAndAmount `xml:"RtrdIntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	//
	// Usage: the InterbankSettlementDate is the interbank settlement date of the return message, and not of the original instruction.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the returned transaction.
	ReturnedInstructedAmount *CurrencyAndAmount `xml:"RtrdInstdAmt,omitempty"`

	// The factor used for conversion of an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Amount of money asked or paid as compensation for the processing of the instruction.
	CompensationAmount *CurrencyAndAmount `xml:"CompstnAmt,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the return message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Transaction charges to be paid by the charge bearer for the return transaction.
	ChargesInformation []*ChargesInformation1 `xml:"ChrgsInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdAgt,omitempty"`

	// Detailed information on the return reason.
	ReturnReasonInformation []*ReturnReasonInformation1 `xml:"RtrRsnInf,omitempty"`

	// Set of key elements of the original transaction being referred to.
	OriginalTransactionReference *OriginalTransactionReference1 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation2) SetReturnIdentification(value string) {
	p.ReturnIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation2) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransactionInformation2) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation2) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation2) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation2) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation2) SetReturnedInterbankSettlementAmount(value, currency string) {
	p.ReturnedInterbankSettlementAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation2) SetInterbankSettlementDate(value string) {
	p.InterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransactionInformation2) SetReturnedInstructedAmount(value, currency string) {
	p.ReturnedInstructedAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation2) SetExchangeRate(value string) {
	p.ExchangeRate = (*BaseOneRate)(&value)
}

func (p *PaymentTransactionInformation2) SetCompensationAmount(value, currency string) {
	p.CompensationAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation2) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransactionInformation2) AddChargesInformation() *ChargesInformation1 {
	newValue := new(ChargesInformation1)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation2) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return p.InstructingAgent
}

func (p *PaymentTransactionInformation2) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification3 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification3)
	return p.InstructedAgent
}

func (p *PaymentTransactionInformation2) AddReturnReasonInformation() *ReturnReasonInformation1 {
	newValue := new(ReturnReasonInformation1)
	p.ReturnReasonInformation = append(p.ReturnReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation2) AddOriginalTransactionReference() *OriginalTransactionReference1 {
	p.OriginalTransactionReference = new(OriginalTransactionReference1)
	return p.OriginalTransactionReference
}
