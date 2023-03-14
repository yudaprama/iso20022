package model

// Provides further details on an entry in the report.
type ReportEntry8 struct {

	// Unique reference for the entry.
	EntryReference *Max35Text `xml:"NtryRef,omitempty"`

	// Amount of money in the cash entry.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether the entry is a credit or a debit entry.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether or not the entry is the result of a reversal.
	// Usage: This element should only be present if the entry is the result of a reversal.
	// If the CreditDebitIndicator is CRDT and ReversalIndicator is Yes, the original operation was a debit entry.
	// If the CreditDebitIndicator is DBIT and ReversalIndicator is Yes, the original operation was a credit entry.
	ReversalIndicator *TrueFalseIndicator `xml:"RvslInd,omitempty"`

	// Status of an entry on the books of the account servicer.
	Status *EntryStatus1Choice `xml:"Sts"`

	// Date and time when an entry is posted to an account on the account servicer's books.
	//
	// Usage: Booking date is the expected booking date, unless the status is booked, in which case it is the actual booking date.
	BookingDate *DateAndDateTimeChoice `xml:"BookgDt,omitempty"`

	// Date and time at which assets become available to the account owner in case of a credit entry, or cease to be available to the account owner in case of a debit entry.
	// Usage: If entry status is pending and value date is present, then the value date refers to an expected/requested value date.
	// For entries subject to availability/float and for which availability information is provided, the value date must not be used. In this case the availability component identifies the number of availability days.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Unique reference as assigned by the account servicing institution to unambiguously identify the entry.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Indicates when the booked amount of money will become available, that is can be accessed and starts generating interest.
	//
	// Usage: This type of information is used in the US and is linked to particular instruments such as cheques.
	// Example: When a cheque is deposited, it will be booked on the deposit day, but the amount of money will only be accessible as of the indicated availability day (according to national banking regulations).
	Availability []*CashAvailability1 `xml:"Avlbty,omitempty"`

	// Set of elements used to fully identify the type of underlying transaction resulting in an entry.
	BankTransactionCode *BankTransactionCodeStructure4 `xml:"BkTxCd"`

	// Indicates whether the transaction is exempt from commission.
	CommissionWaiverIndicator *YesNoIndicator `xml:"ComssnWvrInd,omitempty"`

	// Indicates whether the underlying transaction details are provided through a separate message, as in the case of aggregate bookings.
	AdditionalInformationIndicator *MessageIdentification2 `xml:"AddtlInfInd,omitempty"`

	// Provides information on the original amount.
	//
	// Usage: This component (on entry level) should be used when a total original batch or aggregate amount has to be provided. If required, the individual original amounts can be included in the same component on transaction details level.
	AmountDetails *AmountAndCurrencyExchange3 `xml:"AmtDtls,omitempty"`

	// Provides information on the charges, pre-advised or included in the entry amount .
	//
	// Usage: This component is used on entry level in case of batch or aggregate bookings.
	//
	Charges *Charges4 `xml:"Chrgs,omitempty"`

	// Channel used to technically input the instruction related to the entry.
	TechnicalInputChannel *TechnicalInputChannel1Choice `xml:"TechInptChanl,omitempty"`

	// Provides details of the interest amount included in the entry amount.
	//
	// Usage: This component is used on entry level in the case of batch or aggregate bookings.
	Interest *TransactionInterest3 `xml:"Intrst,omitempty"`

	// Provides details of the card transaction included in the entry amount, when globalised by the account servicer .
	CardTransaction *CardEntry2 `xml:"CardTx,omitempty"`

	// Provides details on the entry.
	EntryDetails []*EntryDetails7 `xml:"NtryDtls,omitempty"`

	// Further details of the entry.
	AdditionalEntryInformation *Max500Text `xml:"AddtlNtryInf,omitempty"`
}

func (r *ReportEntry8) SetEntryReference(value string) {
	r.EntryReference = (*Max35Text)(&value)
}

func (r *ReportEntry8) SetAmount(value, currency string) {
	r.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *ReportEntry8) SetCreditDebitIndicator(value string) {
	r.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (r *ReportEntry8) SetReversalIndicator(value string) {
	r.ReversalIndicator = (*TrueFalseIndicator)(&value)
}

func (r *ReportEntry8) SetStatus(cd, prtry string) {
	r.Status.Cd = (*ExternalEntryStatus1Code)(&cd)
	r.Status.Prtry = (*Max35Text)(&prtry)
}

func (r *ReportEntry8) AddBookingDate() *DateAndDateTimeChoice {
	r.BookingDate = new(DateAndDateTimeChoice)
	return r.BookingDate
}

func (r *ReportEntry8) AddValueDate() *DateAndDateTimeChoice {
	r.ValueDate = new(DateAndDateTimeChoice)
	return r.ValueDate
}

func (r *ReportEntry8) SetAccountServicerReference(value string) {
	r.AccountServicerReference = (*Max35Text)(&value)
}

func (r *ReportEntry8) AddAvailability() *CashAvailability1 {
	newValue := new(CashAvailability1)
	r.Availability = append(r.Availability, newValue)
	return newValue
}

func (r *ReportEntry8) AddBankTransactionCode() *BankTransactionCodeStructure4 {
	r.BankTransactionCode = new(BankTransactionCodeStructure4)
	return r.BankTransactionCode
}

func (r *ReportEntry8) SetCommissionWaiverIndicator(value string) {
	r.CommissionWaiverIndicator = (*YesNoIndicator)(&value)
}

func (r *ReportEntry8) AddAdditionalInformationIndicator() *MessageIdentification2 {
	r.AdditionalInformationIndicator = new(MessageIdentification2)
	return r.AdditionalInformationIndicator
}

func (r *ReportEntry8) AddAmountDetails() *AmountAndCurrencyExchange3 {
	r.AmountDetails = new(AmountAndCurrencyExchange3)
	return r.AmountDetails
}

func (r *ReportEntry8) AddCharges() *Charges4 {
	r.Charges = new(Charges4)
	return r.Charges
}

func (r *ReportEntry8) AddTechnicalInputChannel() *TechnicalInputChannel1Choice {
	r.TechnicalInputChannel = new(TechnicalInputChannel1Choice)
	return r.TechnicalInputChannel
}

func (r *ReportEntry8) AddInterest() *TransactionInterest3 {
	r.Interest = new(TransactionInterest3)
	return r.Interest
}

func (r *ReportEntry8) AddCardTransaction() *CardEntry2 {
	r.CardTransaction = new(CardEntry2)
	return r.CardTransaction
}

func (r *ReportEntry8) AddEntryDetails() *EntryDetails7 {
	newValue := new(EntryDetails7)
	r.EntryDetails = append(r.EntryDetails, newValue)
	return newValue
}

func (r *ReportEntry8) SetAdditionalEntryInformation(value string) {
	r.AdditionalEntryInformation = (*Max500Text)(&value)
}
