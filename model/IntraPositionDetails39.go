package model

// Details of the intra-position movement.
type IntraPositionDetails39 struct {

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification92Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Number identifying a lot constituting the financial instrument.
	LotNumber *GenericIdentification37 `xml:"LotNb,omitempty"`

	// Date and time at which the securities are to be moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Time stamp on when the transaction is acknowledged.
	AcknowledgedStatusTimeStamp *ISODateTime `xml:"AckdStsTmStmp,omitempty"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesBalanceType7Choice `xml:"BalFr,omitempty"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesBalanceType7Choice `xml:"BalTo,omitempty"`
}

func (i *IntraPositionDetails39) SetPoolIdentification(value string) {
	i.PoolIdentification = (*Max35Text)(&value)
}

func (i *IntraPositionDetails39) AddAccountOwner() *PartyIdentification92Choice {
	i.AccountOwner = new(PartyIdentification92Choice)
	return i.AccountOwner
}

func (i *IntraPositionDetails39) AddSafekeepingAccount() *SecuritiesAccount19 {
	i.SafekeepingAccount = new(SecuritiesAccount19)
	return i.SafekeepingAccount
}

func (i *IntraPositionDetails39) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	i.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return i.FinancialInstrumentIdentification
}

func (i *IntraPositionDetails39) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettlementQuantity
}

func (i *IntraPositionDetails39) AddLotNumber() *GenericIdentification37 {
	i.LotNumber = new(GenericIdentification37)
	return i.LotNumber
}

func (i *IntraPositionDetails39) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails39) SetAcknowledgedStatusTimeStamp(value string) {
	i.AcknowledgedStatusTimeStamp = (*ISODateTime)(&value)
}

func (i *IntraPositionDetails39) AddBalanceFrom() *SecuritiesBalanceType7Choice {
	i.BalanceFrom = new(SecuritiesBalanceType7Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails39) AddBalanceTo() *SecuritiesBalanceType7Choice {
	i.BalanceTo = new(SecuritiesBalanceType7Choice)
	return i.BalanceTo
}
