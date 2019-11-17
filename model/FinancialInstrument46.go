package model

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, for example, dividend option or valuation currency.
type FinancialInstrument46 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification23Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Financial Instrument Short Name (FISN) expressed in conformance with the ISO 18774 standard.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Specifies whether the financial instrument is transferred as an asset or as cash.
	TransferType *TransferType1Code `xml:"TrfTp"`

	// Specifies the quantity of assets to be transferred in units or in a percentage rate.
	Quantity *Quantity12Choice `xml:"Qty,omitempty"`

	// Average cost per share of a security, including all charges and commissions.
	AverageAcquisitionPrice *ActiveOrHistoricCurrencyAndAmount `xml:"AvrgAcqstnPric,omitempty"`

	// Identifies the currency to be used to transfer the holdings. Some transfer agents register holdings grouped by currency in addition to using the ISIN for multi-currency fund shares.
	TransferCurrency *ActiveOrHistoricCurrencyCode `xml:"TrfCcy,omitempty"`

	// Net asset on balance sheet - total portfolio value minus or plus the unrealised gain or loss.
	TotalBookValue *ActiveOrHistoricCurrencyAndAmount `xml:"TtlBookVal,omitempty"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account19 `xml:"TrfeeAcct,omitempty"`

	// Sub-accounts that are grouped in a master or omnibus account.
	SubAccountDetails *SubAccount5 `xml:"SubAcctDtls,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesReceivingSideDetails *ReceivingPartiesAndAccount14 `xml:"SttlmPtiesRcvgSdDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount125 `xml:"DlvrgAgtDtls,omitempty"`
}

func (f *FinancialInstrument46) AddIdentification() *SecurityIdentification23Choice {
	f.Identification = new(SecurityIdentification23Choice)
	return f.Identification
}

func (f *FinancialInstrument46) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument46) SetShortName(value string) {
	f.ShortName = (*Max35Text)(&value)
}

func (f *FinancialInstrument46) SetTransferType(value string) {
	f.TransferType = (*TransferType1Code)(&value)
}

func (f *FinancialInstrument46) AddQuantity() *Quantity12Choice {
	f.Quantity = new(Quantity12Choice)
	return f.Quantity
}

func (f *FinancialInstrument46) SetAverageAcquisitionPrice(value, currency string) {
	f.AverageAcquisitionPrice = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument46) SetTransferCurrency(value string) {
	f.TransferCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrument46) SetTotalBookValue(value, currency string) {
	f.TotalBookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument46) AddTransfereeAccount() *Account19 {
	f.TransfereeAccount = new(Account19)
	return f.TransfereeAccount
}

func (f *FinancialInstrument46) AddSubAccountDetails() *SubAccount5 {
	f.SubAccountDetails = new(SubAccount5)
	return f.SubAccountDetails
}

func (f *FinancialInstrument46) AddSettlementPartiesReceivingSideDetails() *ReceivingPartiesAndAccount14 {
	f.SettlementPartiesReceivingSideDetails = new(ReceivingPartiesAndAccount14)
	return f.SettlementPartiesReceivingSideDetails
}

func (f *FinancialInstrument46) AddDeliveringAgentDetails() *PartyIdentificationAndAccount125 {
	f.DeliveringAgentDetails = new(PartyIdentificationAndAccount125)
	return f.DeliveringAgentDetails
}
