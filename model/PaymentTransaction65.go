package model

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the return message applies.
type PaymentTransaction65 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the returned transaction.
	// Usage: The instructing party is the party sending the return message and not the party that sent the original instruction that is being returned.
	ReturnIdentification *Max35Text `xml:"RtrId,omitempty"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent in the returned transaction.
	ReturnedInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"RtrdIntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	//
	// Usage: the InterbankSettlementDate is the interbank settlement date of the return message, and not of the original instruction.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the settlement instruction.
	//
	//
	// Usage: the SettlementPriority is the settlement priority of the return message, and not of the original instruction.
	SettlementPriority *Priority3Code `xml:"SttlmPrty,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the returned transaction.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	ReturnedInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RtrdInstdAmt,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Amount of money asked or paid as compensation for the processing of the instruction.
	CompensationAmount *ActiveOrHistoricCurrencyAndAmount `xml:"CompstnAmt,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the return message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides information on the charges to be paid by the charge bearer(s) related to the processing of the return transaction.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the return message and not the party that sent the original instruction that is being returned.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the return message and not the party that received the original instruction that is being returned.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Provides detailed information on the return reason.
	ReturnReasonInformation []*PaymentReturnReason1 `xml:"RtrRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference22 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction65) SetReturnIdentification(value string) {
	p.ReturnIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction65) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction65) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction65) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction65) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction65) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction65) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction65) SetReturnedInterbankSettlementAmount(value, currency string) {
	p.ReturnedInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction65) SetInterbankSettlementDate(value string) {
	p.InterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction65) SetSettlementPriority(value string) {
	p.SettlementPriority = (*Priority3Code)(&value)
}

func (p *PaymentTransaction65) SetReturnedInstructedAmount(value, currency string) {
	p.ReturnedInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction65) SetExchangeRate(value string) {
	p.ExchangeRate = (*BaseOneRate)(&value)
}

func (p *PaymentTransaction65) SetCompensationAmount(value, currency string) {
	p.CompensationAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction65) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransaction65) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction65) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructingAgent
}

func (p *PaymentTransaction65) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructedAgent
}

func (p *PaymentTransaction65) AddReturnReasonInformation() *PaymentReturnReason1 {
	newValue := new(PaymentReturnReason1)
	p.ReturnReasonInformation = append(p.ReturnReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction65) AddOriginalTransactionReference() *OriginalTransactionReference22 {
	p.OriginalTransactionReference = new(OriginalTransactionReference22)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction65) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}
