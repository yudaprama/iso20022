package model

// Fund Processing Passsport (FPP) is a fully harmonised document with all key operational information that fund promoters should provide on their investment funds in order to facilitate their trading.
type FundProcessingPassport1 struct {

	// Date of last revision.
	UpdatedDate *UpdatedDate `xml:"UpdtdDt"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	SecurityIdentification *SecurityIdentification1 `xml:"SctyId"`

	// Principal entity appointed by the fund, to which orders should be submitted. Usually located in the country of domicile.
	MainFundOrderDesk *ContactAttributes1 `xml:"MainFndOrdrDsk"`

	// Company that is responsible for the management and operation of the fund, eg, determines the investment strategy, appoints
	// the service providers, and makes major decisions for the fund. It is usually responsible for the distribution and marketing
	// of the fund. For self-managed funds, this wlll often be a separate promoter or sponsor of the fund.
	FundManagementCompany *ContactAttributes1 `xml:"FndMgmtCpny"`

	// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
	FundDetails *FinancialInstrument20 `xml:"FndDtls"`

	// Processing characteristics linked to the instrument, ie, not to  the market.
	ValuationDealingCharacteristics *ValuationDealingProcessingCharacteristics2 `xml:"ValtnDealgChrtcs"`

	// Processing characteristics linked to the instrument, ie, not to  the market.
	InvestmentRestrictions *InvestmentRestrictions2 `xml:"InvstmtRstrctns"`

	// Processing characteristics linked to the instrument, ie, not to  the market.
	SubscriptionProcessingCharacteristics *ProcessingCharacteristics2 `xml:"SbcptPrcgChrtcs"`

	// Processing characteristics linked to the instrument, ie, not to  the market.
	RedemptionProcessingCharacteristics *ProcessingCharacteristics3 `xml:"RedPrcgChrtcs"`

	// Account to or from which a cash entry is made.
	SettlementDetails []*CashAccount22 `xml:"SttlmDtls"`

	// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
	LocalMarketAnnex []*LocalMarketAnnex2 `xml:"LclMktAnx,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (f *FundProcessingPassport1) AddUpdatedDate() *UpdatedDate {
	f.UpdatedDate = new(UpdatedDate)
	return f.UpdatedDate
}

func (f *FundProcessingPassport1) AddSecurityIdentification() *SecurityIdentification1 {
	f.SecurityIdentification = new(SecurityIdentification1)
	return f.SecurityIdentification
}

func (f *FundProcessingPassport1) AddMainFundOrderDesk() *ContactAttributes1 {
	f.MainFundOrderDesk = new(ContactAttributes1)
	return f.MainFundOrderDesk
}

func (f *FundProcessingPassport1) AddFundManagementCompany() *ContactAttributes1 {
	f.FundManagementCompany = new(ContactAttributes1)
	return f.FundManagementCompany
}

func (f *FundProcessingPassport1) AddFundDetails() *FinancialInstrument20 {
	f.FundDetails = new(FinancialInstrument20)
	return f.FundDetails
}

func (f *FundProcessingPassport1) AddValuationDealingCharacteristics() *ValuationDealingProcessingCharacteristics2 {
	f.ValuationDealingCharacteristics = new(ValuationDealingProcessingCharacteristics2)
	return f.ValuationDealingCharacteristics
}

func (f *FundProcessingPassport1) AddInvestmentRestrictions() *InvestmentRestrictions2 {
	f.InvestmentRestrictions = new(InvestmentRestrictions2)
	return f.InvestmentRestrictions
}

func (f *FundProcessingPassport1) AddSubscriptionProcessingCharacteristics() *ProcessingCharacteristics2 {
	f.SubscriptionProcessingCharacteristics = new(ProcessingCharacteristics2)
	return f.SubscriptionProcessingCharacteristics
}

func (f *FundProcessingPassport1) AddRedemptionProcessingCharacteristics() *ProcessingCharacteristics3 {
	f.RedemptionProcessingCharacteristics = new(ProcessingCharacteristics3)
	return f.RedemptionProcessingCharacteristics
}

func (f *FundProcessingPassport1) AddSettlementDetails() *CashAccount22 {
	newValue := new(CashAccount22)
	f.SettlementDetails = append(f.SettlementDetails, newValue)
	return newValue
}

func (f *FundProcessingPassport1) AddLocalMarketAnnex() *LocalMarketAnnex2 {
	newValue := new(LocalMarketAnnex2)
	f.LocalMarketAnnex = append(f.LocalMarketAnnex, newValue)
	return newValue
}

func (f *FundProcessingPassport1) AddExtension() *Extension1 {
	newValue := new(Extension1)
	f.Extension = append(f.Extension, newValue)
	return newValue
}
