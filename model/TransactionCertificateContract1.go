package model

// Certificate and contract reference of a transaction.
type TransactionCertificateContract1 struct {

	// Reference of the contract provided as through the date and identification of the contract or through the registered contract identification.
	ContractReference *ContractRegistrationReference1Choice `xml:"CtrctRef,omitempty"`

	// Provides the amount of the transaction in the currency of the registered contract.
	TransactionAmountInContractCurrency *ActiveCurrencyAndAmount `xml:"TxAmtInCtrctCcy,omitempty"`

	// Expected shipment date as per registered contract.
	ExpectedShipmentDate *ISODate `xml:"XpctdShipmntDt,omitempty"`

	// Expected advance payment (or prepayment) return date in case counterparty will not deliver the goods/services.
	ExpectedAdvancePaymentReturnDate *ISODate `xml:"XpctdAdvncPmtRtrDt,omitempty"`

	// Further details on the transaction certificate contract.
	AdditionalInformation *Max1025Text `xml:"AddtlInf,omitempty"`
}

func (t *TransactionCertificateContract1) AddContractReference() *ContractRegistrationReference1Choice {
	t.ContractReference = new(ContractRegistrationReference1Choice)
	return t.ContractReference
}

func (t *TransactionCertificateContract1) SetTransactionAmountInContractCurrency(value, currency string) {
	t.TransactionAmountInContractCurrency = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TransactionCertificateContract1) SetExpectedShipmentDate(value string) {
	t.ExpectedShipmentDate = (*ISODate)(&value)
}

func (t *TransactionCertificateContract1) SetExpectedAdvancePaymentReturnDate(value string) {
	t.ExpectedAdvancePaymentReturnDate = (*ISODate)(&value)
}

func (t *TransactionCertificateContract1) SetAdditionalInformation(value string) {
	t.AdditionalInformation = (*Max1025Text)(&value)
}
