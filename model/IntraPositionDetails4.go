package model

// Details of the intra-position movement.
type IntraPositionDetails4 struct {

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification11 `xml:"FinInstrmId"`

	// Quantity of financial instrument effectively settled.
	SettledQuantity *FinancialInstrumentQuantity1Choice `xml:"SttldQty"`

	// Date and time at which the securities are to be moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesBalanceType2Choice `xml:"BalFr,omitempty"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesBalanceType2Choice `xml:"BalTo,omitempty"`
}

func (i *IntraPositionDetails4) SetPoolIdentification(value string) {
	i.PoolIdentification = (*Max35Text)(&value)
}

func (i *IntraPositionDetails4) AddAccountOwner() *PartyIdentification13Choice {
	i.AccountOwner = new(PartyIdentification13Choice)
	return i.AccountOwner
}

func (i *IntraPositionDetails4) AddSafekeepingAccount() *SecuritiesAccount13 {
	i.SafekeepingAccount = new(SecuritiesAccount13)
	return i.SafekeepingAccount
}

func (i *IntraPositionDetails4) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	i.FinancialInstrumentIdentification = new(SecurityIdentification11)
	return i.FinancialInstrumentIdentification
}

func (i *IntraPositionDetails4) AddSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettledQuantity
}

func (i *IntraPositionDetails4) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails4) AddBalanceFrom() *SecuritiesBalanceType2Choice {
	i.BalanceFrom = new(SecuritiesBalanceType2Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails4) AddBalanceTo() *SecuritiesBalanceType2Choice {
	i.BalanceTo = new(SecuritiesBalanceType2Choice)
	return i.BalanceTo
}
