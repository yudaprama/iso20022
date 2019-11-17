package model

// Additional specific modification criteria.
type AdditionalInformation9 struct {

	// Identification of the transaction as known by the account owner. Will be used in a unilateral split to provide the executing party with the account owner identification of each split transaction.
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Type of instrument involved in the transactions on which the modification request should apply.
	ClassificationType *ClassificationType1Choice `xml:"ClssfctnTp,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Identification of the financial instrument involved in the transactions on which the modification request should apply.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Quantity of financial instrument concerned by the settlement condition modification request.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty,omitempty"`

	// Date/time when the request should take effect.
	EffectiveDate *DateAndDateTimeChoice `xml:"FctvDt,omitempty"`

	// Date/time when the request should cease to be in effect.
	ExpiryDate *DateAndDateTimeChoice `xml:"XpryDt,omitempty"`

	// Date/time of the release.
	CutOffDate *DateAndDateTimeChoice `xml:"CutOffDt,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification43Choice `xml:"Invstr,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	DeliveringParty1 *PartyIdentificationAndAccount43 `xml:"DlvrgPty1,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	ReceivingParty1 *PartyIdentificationAndAccount43 `xml:"RcvgPty1,omitempty"`

	// Provides details on the processing status of the transaction.
	ProcessingStatus *ProcessingStatus39Choice `xml:"PrcgSts,omitempty"`
}

func (a *AdditionalInformation9) SetAccountOwnerTransactionIdentification(value string) {
	a.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalInformation9) AddClassificationType() *ClassificationType1Choice {
	a.ClassificationType = new(ClassificationType1Choice)
	return a.ClassificationType
}

func (a *AdditionalInformation9) AddSafekeepingAccount() *SecuritiesAccount13 {
	a.SafekeepingAccount = new(SecuritiesAccount13)
	return a.SafekeepingAccount
}

func (a *AdditionalInformation9) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return a.FinancialInstrumentIdentification
}

func (a *AdditionalInformation9) AddQuantity() *FinancialInstrumentQuantity1Choice {
	a.Quantity = new(FinancialInstrumentQuantity1Choice)
	return a.Quantity
}

func (a *AdditionalInformation9) AddEffectiveDate() *DateAndDateTimeChoice {
	a.EffectiveDate = new(DateAndDateTimeChoice)
	return a.EffectiveDate
}

func (a *AdditionalInformation9) AddExpiryDate() *DateAndDateTimeChoice {
	a.ExpiryDate = new(DateAndDateTimeChoice)
	return a.ExpiryDate
}

func (a *AdditionalInformation9) AddCutOffDate() *DateAndDateTimeChoice {
	a.CutOffDate = new(DateAndDateTimeChoice)
	return a.CutOffDate
}

func (a *AdditionalInformation9) AddInvestor() *PartyIdentification43Choice {
	a.Investor = new(PartyIdentification43Choice)
	return a.Investor
}

func (a *AdditionalInformation9) AddDeliveringParty1() *PartyIdentificationAndAccount43 {
	a.DeliveringParty1 = new(PartyIdentificationAndAccount43)
	return a.DeliveringParty1
}

func (a *AdditionalInformation9) AddReceivingParty1() *PartyIdentificationAndAccount43 {
	a.ReceivingParty1 = new(PartyIdentificationAndAccount43)
	return a.ReceivingParty1
}

func (a *AdditionalInformation9) AddProcessingStatus() *ProcessingStatus39Choice {
	a.ProcessingStatus = new(ProcessingStatus39Choice)
	return a.ProcessingStatus
}
