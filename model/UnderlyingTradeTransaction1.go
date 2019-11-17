package model

// Information about a transaction.
type UnderlyingTradeTransaction1 struct {

	// Type of underlying transaction such as a tender, order, contract.
	Type *UnderlyingTradeTransactionType1Choice `xml:"Tp"`

	// Identification of the underlying transaction.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Date the underlying transaction was issued or awarded.
	TransactionDate *ISODate `xml:"TxDt,omitempty"`

	// Date the tender closes.
	TenderClosingDate *ISODate `xml:"TndrClsgDt,omitempty"`

	// Amount of the underlying transaction.
	TransactionAmount *ActiveCurrencyAndAmount `xml:"TxAmt,omitempty"`

	// Percentage of the underlying contract covered by the undertaking.
	ContractAmountPercentage *PercentageRate `xml:"CtrctAmtPctg,omitempty"`

	// Additional information related to the underlying transaction.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (u *UnderlyingTradeTransaction1) AddType() *UnderlyingTradeTransactionType1Choice {
	u.Type = new(UnderlyingTradeTransactionType1Choice)
	return u.Type
}

func (u *UnderlyingTradeTransaction1) SetIdentification(value string) {
	u.Identification = (*Max35Text)(&value)
}

func (u *UnderlyingTradeTransaction1) SetTransactionDate(value string) {
	u.TransactionDate = (*ISODate)(&value)
}

func (u *UnderlyingTradeTransaction1) SetTenderClosingDate(value string) {
	u.TenderClosingDate = (*ISODate)(&value)
}

func (u *UnderlyingTradeTransaction1) SetTransactionAmount(value, currency string) {
	u.TransactionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (u *UnderlyingTradeTransaction1) SetContractAmountPercentage(value string) {
	u.ContractAmountPercentage = (*PercentageRate)(&value)
}

func (u *UnderlyingTradeTransaction1) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*Max2000Text)(&value))
}
