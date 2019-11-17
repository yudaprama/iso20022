package model

// Additional specific modification criteria.
type AdditionalInformation3 struct {

	// Identification of the transaction as known by the account owner. Will be used in a unilateral split to provide the executing party with the account owner identification of each split transaction.
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Type of instrument involved in the transactions on which the modification request should apply.
	ClassificationType *ClassificationType1Choice `xml:"ClssfctnTp,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Identification of the financial instrument involved in the transactions on which the modification request should apply.
	FinancialInstrumentIdentification *SecurityIdentification11 `xml:"FinInstrmId,omitempty"`

	// Quantity of financial instrument concerned by the settlement condition modification request.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty,omitempty"`

	// Date/time when the request should take effect.
	EffectiveDate *DateAndDateTimeChoice `xml:"FctvDt,omitempty"`

	// Date/time when the request should cease to be in effect.
	ExpiryDate *DateAndDateTimeChoice `xml:"XpryDt,omitempty"`

	// Date/time of the release.
	CutOffDate *DateAndDateTimeChoice `xml:"CutOffDt,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification10Choice `xml:"Invstr,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	DeliveringParty1 *PartyIdentificationAndAccount16 `xml:"DlvrgPty1,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	ReceivingParty1 *PartyIdentificationAndAccount16 `xml:"RcvgPty1,omitempty"`
}

func (a *AdditionalInformation3) SetAccountOwnerTransactionIdentification(value string) {
	a.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalInformation3) AddClassificationType() *ClassificationType1Choice {
	a.ClassificationType = new(ClassificationType1Choice)
	return a.ClassificationType
}

func (a *AdditionalInformation3) AddSafekeepingAccount() *SecuritiesAccount13 {
	a.SafekeepingAccount = new(SecuritiesAccount13)
	return a.SafekeepingAccount
}

func (a *AdditionalInformation3) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification11)
	return a.FinancialInstrumentIdentification
}

func (a *AdditionalInformation3) AddQuantity() *FinancialInstrumentQuantity1Choice {
	a.Quantity = new(FinancialInstrumentQuantity1Choice)
	return a.Quantity
}

func (a *AdditionalInformation3) AddEffectiveDate() *DateAndDateTimeChoice {
	a.EffectiveDate = new(DateAndDateTimeChoice)
	return a.EffectiveDate
}

func (a *AdditionalInformation3) AddExpiryDate() *DateAndDateTimeChoice {
	a.ExpiryDate = new(DateAndDateTimeChoice)
	return a.ExpiryDate
}

func (a *AdditionalInformation3) AddCutOffDate() *DateAndDateTimeChoice {
	a.CutOffDate = new(DateAndDateTimeChoice)
	return a.CutOffDate
}

func (a *AdditionalInformation3) AddInvestor() *PartyIdentification10Choice {
	a.Investor = new(PartyIdentification10Choice)
	return a.Investor
}

func (a *AdditionalInformation3) AddDeliveringParty1() *PartyIdentificationAndAccount16 {
	a.DeliveringParty1 = new(PartyIdentificationAndAccount16)
	return a.DeliveringParty1
}

func (a *AdditionalInformation3) AddReceivingParty1() *PartyIdentificationAndAccount16 {
	a.ReceivingParty1 = new(PartyIdentificationAndAccount16)
	return a.ReceivingParty1
}
