package model

// Details of the intra-position movement.
type IntraPositionDetails19 struct {

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Number identifying a lot constituting the financial instrument.
	LotNumber *GenericIdentification37 `xml:"LotNb,omitempty"`

	// Date and time at which the securities are to be moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesBalanceType2Choice `xml:"BalFr,omitempty"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesBalanceType2Choice `xml:"BalTo,omitempty"`
}

func (i *IntraPositionDetails19) SetPoolIdentification(value string) {
	i.PoolIdentification = (*Max35Text)(&value)
}

func (i *IntraPositionDetails19) AddAccountOwner() *PartyIdentification36Choice {
	i.AccountOwner = new(PartyIdentification36Choice)
	return i.AccountOwner
}

func (i *IntraPositionDetails19) AddSafekeepingAccount() *SecuritiesAccount13 {
	i.SafekeepingAccount = new(SecuritiesAccount13)
	return i.SafekeepingAccount
}

func (i *IntraPositionDetails19) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	i.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return i.FinancialInstrumentIdentification
}

func (i *IntraPositionDetails19) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettlementQuantity
}

func (i *IntraPositionDetails19) AddLotNumber() *GenericIdentification37 {
	i.LotNumber = new(GenericIdentification37)
	return i.LotNumber
}

func (i *IntraPositionDetails19) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails19) AddBalanceFrom() *SecuritiesBalanceType2Choice {
	i.BalanceFrom = new(SecuritiesBalanceType2Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails19) AddBalanceTo() *SecuritiesBalanceType2Choice {
	i.BalanceTo = new(SecuritiesBalanceType2Choice)
	return i.BalanceTo
}
