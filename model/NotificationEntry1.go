package model

// Specifies the elements of an entry in the report.
type NotificationEntry1 struct {

	// Amount of money in the cash entry.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Specifies if an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the entry is the result of a reversal operation.
	//
	// Usage : this element should only be present if the entry is the result of a reversal operation.
	// If the CreditDebitIndicator is CRDT and ReversalIndicator is Yes, the original operation was a debit entry.
	// If the CreditDebitIndicator is DBIT and ReversalIndicator is Yes, the original operation was a credit entry.
	ReversalIndicator *TrueFalseIndicator `xml:"RvslInd,omitempty"`

	// Status of an entry on the books of the account servicer.
	Status *EntryStatus4Code `xml:"Sts"`

	// Date and time when an entry is posted to an account on the account servicer's books.
	//
	// Usage : Booking date is only present if Status is booked.
	BookingDate *DateAndDateTimeChoice `xml:"BookgDt,omitempty"`

	// Date and time assets become available to the account owner (in a credit entry), or cease to be available to the account owner (in a debit entry).
	//
	// Usage : When entry status is pending , and value date is present, value date refers to an expected/requested value date.
	// For entries which are subject to availability/float (and for which availability information is present), value date must  not be used, as the availability component identifies the number of availability days.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Account servicing institution's reference for the entry.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Set of elements used to indicate when the booked amount of money will become available, ie can be accessed and start generating interest.
	//
	// Usage : this type of info is eg used in US, and is linked to particular instruments, such as cheques.
	// Example : When a cheque is deposited, it will be booked on the deposit day, but the funds will only be accessible as of the indicated availability day (according to national banking regulations).
	Availability []*CashBalanceAvailability1 `xml:"Avlbty,omitempty"`

	// Set of elements to fully identify the type of underlying transaction resulting in the entry.
	BankTransactionCode *BankTransactionCodeStructure1 `xml:"BkTxCd"`

	// Indicates whether the transaction is exempt from commission.
	CommissionWaiverIndicator *YesNoIndicator `xml:"ComssnWvrInd,omitempty"`

	// Indicates whether the underlying transaction details are provided through a separate message, eg in case of aggregate bookings.
	AdditionalInformationIndicator *MessageIdentification2 `xml:"AddtlInfInd,omitempty"`

	// Set of elements providing details on batched transactions.
	//
	// Usage : this element can be repeated in case more than one batch is included in the entry, eg, in lockbox scenarios, to specify the ID and number of transactions included in each of the batches.
	Batch []*BatchInformation1 `xml:"Btch,omitempty"`

	// Set of elements providing information on the original amount.
	//
	// Usage : This component (on entry level) should be used when a total original batch or aggregate amount has to be provided. (If required, the individual original amounts can be included in the same component on transaction details level).
	AmountDetails *AmountAndCurrencyExchange2 `xml:"AmtDtls,omitempty"`

	// Provides information on the charges included in the entry amount.
	//
	// Usage : this component is used on entry level in case of batch or aggregate bookings.
	Charges []*ChargesInformation3 `xml:"Chrgs,omitempty"`

	// Set of elements providing details on the interest amount included in the entry amount.
	//
	// Usage : This component is used on entry level in case of batch or aggregate bookings.
	Interest []*TransactionInterest1 `xml:"Intrst,omitempty"`

	// Set of elements providing information on the underlying transaction (s).
	TransactionDetails []*EntryTransaction1 `xml:"TxDtls,omitempty"`

	// Further details on the entry details.
	AdditionalEntryInformation *Max500Text `xml:"AddtlNtryInf,omitempty"`
}

func (n *NotificationEntry1) SetAmount(value, currency string) {
	n.Amount = NewCurrencyAndAmount(value, currency)
}

func (n *NotificationEntry1) SetCreditDebitIndicator(value string) {
	n.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (n *NotificationEntry1) SetReversalIndicator(value string) {
	n.ReversalIndicator = (*TrueFalseIndicator)(&value)
}

func (n *NotificationEntry1) SetStatus(value string) {
	n.Status = (*EntryStatus4Code)(&value)
}

func (n *NotificationEntry1) AddBookingDate() *DateAndDateTimeChoice {
	n.BookingDate = new(DateAndDateTimeChoice)
	return n.BookingDate
}

func (n *NotificationEntry1) AddValueDate() *DateAndDateTimeChoice {
	n.ValueDate = new(DateAndDateTimeChoice)
	return n.ValueDate
}

func (n *NotificationEntry1) SetAccountServicerReference(value string) {
	n.AccountServicerReference = (*Max35Text)(&value)
}

func (n *NotificationEntry1) AddAvailability() *CashBalanceAvailability1 {
	newValue := new(CashBalanceAvailability1)
	n.Availability = append(n.Availability, newValue)
	return newValue
}

func (n *NotificationEntry1) AddBankTransactionCode() *BankTransactionCodeStructure1 {
	n.BankTransactionCode = new(BankTransactionCodeStructure1)
	return n.BankTransactionCode
}

func (n *NotificationEntry1) SetCommissionWaiverIndicator(value string) {
	n.CommissionWaiverIndicator = (*YesNoIndicator)(&value)
}

func (n *NotificationEntry1) AddAdditionalInformationIndicator() *MessageIdentification2 {
	n.AdditionalInformationIndicator = new(MessageIdentification2)
	return n.AdditionalInformationIndicator
}

func (n *NotificationEntry1) AddBatch() *BatchInformation1 {
	newValue := new(BatchInformation1)
	n.Batch = append(n.Batch, newValue)
	return newValue
}

func (n *NotificationEntry1) AddAmountDetails() *AmountAndCurrencyExchange2 {
	n.AmountDetails = new(AmountAndCurrencyExchange2)
	return n.AmountDetails
}

func (n *NotificationEntry1) AddCharges() *ChargesInformation3 {
	newValue := new(ChargesInformation3)
	n.Charges = append(n.Charges, newValue)
	return newValue
}

func (n *NotificationEntry1) AddInterest() *TransactionInterest1 {
	newValue := new(TransactionInterest1)
	n.Interest = append(n.Interest, newValue)
	return newValue
}

func (n *NotificationEntry1) AddTransactionDetails() *EntryTransaction1 {
	newValue := new(EntryTransaction1)
	n.TransactionDetails = append(n.TransactionDetails, newValue)
	return newValue
}

func (n *NotificationEntry1) SetAdditionalEntryInformation(value string) {
	n.AdditionalEntryInformation = (*Max500Text)(&value)
}
