package model

// Information requested against money laundering for a transfer transaction.
type AntiMoneyLaundering1 struct {

	// Name of the sender.
	SenderName *Max70Text `xml:"SndrNm,omitempty"`

	// Address of the sender.
	SenderAddress *PostalAddress18 `xml:"SndrAdr,omitempty"`

	// National identifier number of the sender.
	SenderNationalIdentifier *Max35Text `xml:"SndrNtlIdr,omitempty"`

	// Country of the national identifier (ISO 3166-1 alpha-2 or alpha-3).
	NationalIdentifierCountry *Min2Max3AlphaText `xml:"NtlIdrCtry,omitempty"`

	// Passport number of the sender.
	SenderPassportNumber *Max35Text `xml:"SndrPsptNb,omitempty"`

	// Country issuing the passport (ISO 3166-1 alpha-2 or alpha-3).
	PassportIssuingCountry *Min2Max3AlphaText `xml:"PsptIssgCtry,omitempty"`

	// Tax identifier of the sender.
	SenderTaxIdentifier *Max35Text `xml:"SndrTaxIdr,omitempty"`

	// Country of the tax (ISO 3166-1 alpha-2 or alpha-3).
	TaxCountry *Min2Max3AlphaText `xml:"TaxCtry,omitempty"`

	// Customer identifier of the sender.
	SenderCustomerIdentifier *Max35Text `xml:"SndrCstmrIdr,omitempty"`

	// Date and place of birth of the sender.
	SenderDateAndPlaceOfBirth *DateAndPlaceOfBirth `xml:"SndrDtAndPlcOfBirth,omitempty"`

	// Name of the receiver.
	ReceiverName *Max70Text `xml:"RcvrNm,omitempty"`

	// Unique transaction reference number for sender and the receiver.
	TransactionReference *Max35Text `xml:"TxRef,omitempty"`
}

func (a *AntiMoneyLaundering1) SetSenderName(value string) {
	a.SenderName = (*Max70Text)(&value)
}

func (a *AntiMoneyLaundering1) AddSenderAddress() *PostalAddress18 {
	a.SenderAddress = new(PostalAddress18)
	return a.SenderAddress
}

func (a *AntiMoneyLaundering1) SetSenderNationalIdentifier(value string) {
	a.SenderNationalIdentifier = (*Max35Text)(&value)
}

func (a *AntiMoneyLaundering1) SetNationalIdentifierCountry(value string) {
	a.NationalIdentifierCountry = (*Min2Max3AlphaText)(&value)
}

func (a *AntiMoneyLaundering1) SetSenderPassportNumber(value string) {
	a.SenderPassportNumber = (*Max35Text)(&value)
}

func (a *AntiMoneyLaundering1) SetPassportIssuingCountry(value string) {
	a.PassportIssuingCountry = (*Min2Max3AlphaText)(&value)
}

func (a *AntiMoneyLaundering1) SetSenderTaxIdentifier(value string) {
	a.SenderTaxIdentifier = (*Max35Text)(&value)
}

func (a *AntiMoneyLaundering1) SetTaxCountry(value string) {
	a.TaxCountry = (*Min2Max3AlphaText)(&value)
}

func (a *AntiMoneyLaundering1) SetSenderCustomerIdentifier(value string) {
	a.SenderCustomerIdentifier = (*Max35Text)(&value)
}

func (a *AntiMoneyLaundering1) AddSenderDateAndPlaceOfBirth() *DateAndPlaceOfBirth {
	a.SenderDateAndPlaceOfBirth = new(DateAndPlaceOfBirth)
	return a.SenderDateAndPlaceOfBirth
}

func (a *AntiMoneyLaundering1) SetReceiverName(value string) {
	a.ReceiverName = (*Max70Text)(&value)
}

func (a *AntiMoneyLaundering1) SetTransactionReference(value string) {
	a.TransactionReference = (*Max35Text)(&value)
}
