package model

// Additional specific modification criteria.
type AdditionalInformation13 struct {

	// Identification of the transaction as known by the account owner. Will be used in a unilateral split to provide the executing party with the account owner identification of each split transaction.
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Type of instrument involved in the transactions on which the modification request should apply.
	ClassificationType *ClassificationType32Choice `xml:"ClssfctnTp,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`

	// Identification of the financial instrument involved in the transactions on which the modification request should apply.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId,omitempty"`

	// Quantity of financial instrument concerned by the settlement condition modification request.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty,omitempty"`

	// Date/time when the request should take effect.
	EffectiveDate *DateAndDateTimeChoice `xml:"FctvDt,omitempty"`

	// Date/time when the request should cease to be in effect.
	ExpiryDate *DateAndDateTimeChoice `xml:"XpryDt,omitempty"`

	// Date/time of the release.
	CutOffDate *DateAndDateTimeChoice `xml:"CutOffDt,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification100 `xml:"Invstr,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	DeliveringParty1 *PartyIdentificationAndAccount117 `xml:"DlvrgPty1,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	ReceivingParty1 *PartyIdentificationAndAccount117 `xml:"RcvgPty1,omitempty"`
}

func (a *AdditionalInformation13) SetAccountOwnerTransactionIdentification(value string) {
	a.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalInformation13) AddClassificationType() *ClassificationType32Choice {
	a.ClassificationType = new(ClassificationType32Choice)
	return a.ClassificationType
}

func (a *AdditionalInformation13) AddSafekeepingAccount() *SecuritiesAccount19 {
	a.SafekeepingAccount = new(SecuritiesAccount19)
	return a.SafekeepingAccount
}

func (a *AdditionalInformation13) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return a.FinancialInstrumentIdentification
}

func (a *AdditionalInformation13) AddQuantity() *FinancialInstrumentQuantity1Choice {
	a.Quantity = new(FinancialInstrumentQuantity1Choice)
	return a.Quantity
}

func (a *AdditionalInformation13) AddEffectiveDate() *DateAndDateTimeChoice {
	a.EffectiveDate = new(DateAndDateTimeChoice)
	return a.EffectiveDate
}

func (a *AdditionalInformation13) AddExpiryDate() *DateAndDateTimeChoice {
	a.ExpiryDate = new(DateAndDateTimeChoice)
	return a.ExpiryDate
}

func (a *AdditionalInformation13) AddCutOffDate() *DateAndDateTimeChoice {
	a.CutOffDate = new(DateAndDateTimeChoice)
	return a.CutOffDate
}

func (a *AdditionalInformation13) AddInvestor() *PartyIdentification100 {
	a.Investor = new(PartyIdentification100)
	return a.Investor
}

func (a *AdditionalInformation13) AddDeliveringParty1() *PartyIdentificationAndAccount117 {
	a.DeliveringParty1 = new(PartyIdentificationAndAccount117)
	return a.DeliveringParty1
}

func (a *AdditionalInformation13) AddReceivingParty1() *PartyIdentificationAndAccount117 {
	a.ReceivingParty1 = new(PartyIdentificationAndAccount117)
	return a.ReceivingParty1
}
