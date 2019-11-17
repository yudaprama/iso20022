package model

// Additional specific modification criteria.
type AdditionalInformation14 struct {

	// Identification of the transaction as known by the account owner. Will be used in a unilateral split to provide the executing party with the account owner identification of each split transaction.
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId,omitempty"`

	// Type of instrument involved in the transactions on which the modification request should apply.
	ClassificationType *ClassificationType33Choice `xml:"ClssfctnTp,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct,omitempty"`

	// Identification of the financial instrument involved in the transactions on which the modification request should apply.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId,omitempty"`

	// Quantity of financial instrument concerned by the settlement condition modification request.
	Quantity *FinancialInstrumentQuantity15Choice `xml:"Qty,omitempty"`

	// Date/time when the request should take effect.
	EffectiveDate *DateAndDateTimeChoice `xml:"FctvDt,omitempty"`

	// Date/time when the request should cease to be in effect.
	ExpiryDate *DateAndDateTimeChoice `xml:"XpryDt,omitempty"`

	// Date/time of the release.
	CutOffDate *DateAndDateTimeChoice `xml:"CutOffDt,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification111 `xml:"Invstr,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	DeliveringParty1 *PartyIdentificationAndAccount146 `xml:"DlvrgPty1,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	ReceivingParty1 *PartyIdentificationAndAccount146 `xml:"RcvgPty1,omitempty"`
}

func (a *AdditionalInformation14) SetAccountOwnerTransactionIdentification(value string) {
	a.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalInformation14) AddClassificationType() *ClassificationType33Choice {
	a.ClassificationType = new(ClassificationType33Choice)
	return a.ClassificationType
}

func (a *AdditionalInformation14) AddSafekeepingAccount() *SecuritiesAccount30 {
	a.SafekeepingAccount = new(SecuritiesAccount30)
	return a.SafekeepingAccount
}

func (a *AdditionalInformation14) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return a.FinancialInstrumentIdentification
}

func (a *AdditionalInformation14) AddQuantity() *FinancialInstrumentQuantity15Choice {
	a.Quantity = new(FinancialInstrumentQuantity15Choice)
	return a.Quantity
}

func (a *AdditionalInformation14) AddEffectiveDate() *DateAndDateTimeChoice {
	a.EffectiveDate = new(DateAndDateTimeChoice)
	return a.EffectiveDate
}

func (a *AdditionalInformation14) AddExpiryDate() *DateAndDateTimeChoice {
	a.ExpiryDate = new(DateAndDateTimeChoice)
	return a.ExpiryDate
}

func (a *AdditionalInformation14) AddCutOffDate() *DateAndDateTimeChoice {
	a.CutOffDate = new(DateAndDateTimeChoice)
	return a.CutOffDate
}

func (a *AdditionalInformation14) AddInvestor() *PartyIdentification111 {
	a.Investor = new(PartyIdentification111)
	return a.Investor
}

func (a *AdditionalInformation14) AddDeliveringParty1() *PartyIdentificationAndAccount146 {
	a.DeliveringParty1 = new(PartyIdentificationAndAccount146)
	return a.DeliveringParty1
}

func (a *AdditionalInformation14) AddReceivingParty1() *PartyIdentificationAndAccount146 {
	a.ReceivingParty1 = new(PartyIdentificationAndAccount146)
	return a.ReceivingParty1
}
