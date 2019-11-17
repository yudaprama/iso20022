package model

// Additional specific modification criteria.
type AdditionalInformation12 struct {

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

	// Provides details on the processing status of the transaction.
	ProcessingStatus *ProcessingStatus59Choice `xml:"PrcgSts,omitempty"`
}

func (a *AdditionalInformation12) SetAccountOwnerTransactionIdentification(value string) {
	a.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalInformation12) AddClassificationType() *ClassificationType33Choice {
	a.ClassificationType = new(ClassificationType33Choice)
	return a.ClassificationType
}

func (a *AdditionalInformation12) AddSafekeepingAccount() *SecuritiesAccount30 {
	a.SafekeepingAccount = new(SecuritiesAccount30)
	return a.SafekeepingAccount
}

func (a *AdditionalInformation12) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return a.FinancialInstrumentIdentification
}

func (a *AdditionalInformation12) AddQuantity() *FinancialInstrumentQuantity15Choice {
	a.Quantity = new(FinancialInstrumentQuantity15Choice)
	return a.Quantity
}

func (a *AdditionalInformation12) AddEffectiveDate() *DateAndDateTimeChoice {
	a.EffectiveDate = new(DateAndDateTimeChoice)
	return a.EffectiveDate
}

func (a *AdditionalInformation12) AddExpiryDate() *DateAndDateTimeChoice {
	a.ExpiryDate = new(DateAndDateTimeChoice)
	return a.ExpiryDate
}

func (a *AdditionalInformation12) AddCutOffDate() *DateAndDateTimeChoice {
	a.CutOffDate = new(DateAndDateTimeChoice)
	return a.CutOffDate
}

func (a *AdditionalInformation12) AddInvestor() *PartyIdentification111 {
	a.Investor = new(PartyIdentification111)
	return a.Investor
}

func (a *AdditionalInformation12) AddDeliveringParty1() *PartyIdentificationAndAccount146 {
	a.DeliveringParty1 = new(PartyIdentificationAndAccount146)
	return a.DeliveringParty1
}

func (a *AdditionalInformation12) AddReceivingParty1() *PartyIdentificationAndAccount146 {
	a.ReceivingParty1 = new(PartyIdentificationAndAccount146)
	return a.ReceivingParty1
}

func (a *AdditionalInformation12) AddProcessingStatus() *ProcessingStatus59Choice {
	a.ProcessingStatus = new(ProcessingStatus59Choice)
	return a.ProcessingStatus
}
