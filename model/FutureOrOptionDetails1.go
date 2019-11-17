package model

// Parameters for contracts which obligate the buyer to receive and the seller to deliver in the future the assets specified at an agreed price or contracts which grant to the holder either the privilege to purchase or the privilege to sell the assets specified at a predetermined price or formula at or within a time in the future.
type FutureOrOptionDetails1 struct {

	// Specifies the type of the contract for futures and options.
	FutureAndOptionContractType *FutureAndOptionContractType1Code `xml:"FutrAndOptnCtrctTp,omitempty"`

	// Last date/time by which the option for physical delivery may still be exercised.
	LastDeliveryDate *ISODateTime `xml:"LastDlvryDt,omitempty"`

	// Used to indicate the size of the underlying commodity on which the contract is based (e.g., 2500 lbs of lean cattle, 1000 barrels of crude oil, 1000 bushels of corn, etc.)
	UnitOfMeasure *UnitOfMeasure1Code `xml:"UnitOfMeasr,omitempty"`

	// Date on which future contracts settle.
	FutureDate *ISODateTime `xml:"FutrDt,omitempty"`

	// Specifies the minimum ratio or multiply factor used to convert from contracts to shares.
	MinimumSize *ActiveCurrencyAndAmount `xml:"MinSz,omitempty"`

	// Date/time, as announced by the issuer, at which the securities will be issued.
	AnnouncementDate *ISODateTime `xml:"AnncmntDt,omitempty"`

	// Specifies the deliverability of a security.
	Appearance *Appearance1Code `xml:"Apprnc,omitempty"`

	// Indicates whether the interest is separable from the principal.
	StrippableIndicator *YesNoIndicator `xml:"StrpblInd,omitempty"`

	// Indicates the maximum number of listed option contracts on a single security which can be held by an investor or group of investors acting jointly.
	PositionLimit *Number `xml:"PosLmt,omitempty"`

	// Position limit in the near-term contract for a given exchange-traded product.
	NearTermPositionLimit *Number `xml:"NearTermPosLmt,omitempty"`

	// Minimum price increase for a given exchange-traded Instrument
	MinimumTradingPricingIncrement *Number `xml:"MinTradgPricgIncrmt,omitempty"`

	// Reason for which money is raised through the issuance of a security.
	Purpose *Max256Text `xml:"Purp,omitempty"`

	// Specifies when the contract (i.e. MBS/TBA) will settle.
	ContractSettlementMonth *ISOYearMonth `xml:"CtrctSttlmMnth,omitempty"`

	// Date on which new securities begin trading.
	FirstDealingDate *DateAndDateTime1Choice `xml:"FrstDealgDt,omitempty"`

	// Ratio applied to convert the related security.
	Ratio []*UnderlyingRatio1 `xml:"Ratio,omitempty"`

	// Rating(s) of the security.
	Rating []*Rating1 `xml:"Ratg,omitempty"`

	// Initial issue price of a financial instrument.
	IssuePrice *Price4 `xml:"IssePric,omitempty"`

	// Rights to exercise the privilege to purchase or to sell the assets specified at a predetermined price or formula at or within a time in the future.
	OptionRights *OptionRight1Choice `xml:"OptnRghts,omitempty"`

	// Indicates whether or not this is the last transaction.
	LastTransaction *YesNoIndicator `xml:"LastTx,omitempty"`

	// Specifies that there will be one price and one transaction when two contracts are carried out simultaneously, one to buy and the other one to sell with two different expiration dates.
	SpreadTransaction *YesNoIndicator `xml:"SprdTx,omitempty"`
}

func (f *FutureOrOptionDetails1) SetFutureAndOptionContractType(value string) {
	f.FutureAndOptionContractType = (*FutureAndOptionContractType1Code)(&value)
}

func (f *FutureOrOptionDetails1) SetLastDeliveryDate(value string) {
	f.LastDeliveryDate = (*ISODateTime)(&value)
}

func (f *FutureOrOptionDetails1) SetUnitOfMeasure(value string) {
	f.UnitOfMeasure = (*UnitOfMeasure1Code)(&value)
}

func (f *FutureOrOptionDetails1) SetFutureDate(value string) {
	f.FutureDate = (*ISODateTime)(&value)
}

func (f *FutureOrOptionDetails1) SetMinimumSize(value, currency string) {
	f.MinimumSize = NewActiveCurrencyAndAmount(value, currency)
}

func (f *FutureOrOptionDetails1) SetAnnouncementDate(value string) {
	f.AnnouncementDate = (*ISODateTime)(&value)
}

func (f *FutureOrOptionDetails1) SetAppearance(value string) {
	f.Appearance = (*Appearance1Code)(&value)
}

func (f *FutureOrOptionDetails1) SetStrippableIndicator(value string) {
	f.StrippableIndicator = (*YesNoIndicator)(&value)
}

func (f *FutureOrOptionDetails1) SetPositionLimit(value string) {
	f.PositionLimit = (*Number)(&value)
}

func (f *FutureOrOptionDetails1) SetNearTermPositionLimit(value string) {
	f.NearTermPositionLimit = (*Number)(&value)
}

func (f *FutureOrOptionDetails1) SetMinimumTradingPricingIncrement(value string) {
	f.MinimumTradingPricingIncrement = (*Number)(&value)
}

func (f *FutureOrOptionDetails1) SetPurpose(value string) {
	f.Purpose = (*Max256Text)(&value)
}

func (f *FutureOrOptionDetails1) SetContractSettlementMonth(value string) {
	f.ContractSettlementMonth = (*ISOYearMonth)(&value)
}

func (f *FutureOrOptionDetails1) AddFirstDealingDate() *DateAndDateTime1Choice {
	f.FirstDealingDate = new(DateAndDateTime1Choice)
	return f.FirstDealingDate
}

func (f *FutureOrOptionDetails1) AddRatio() *UnderlyingRatio1 {
	newValue := new(UnderlyingRatio1)
	f.Ratio = append(f.Ratio, newValue)
	return newValue
}

func (f *FutureOrOptionDetails1) AddRating() *Rating1 {
	newValue := new(Rating1)
	f.Rating = append(f.Rating, newValue)
	return newValue
}

func (f *FutureOrOptionDetails1) AddIssuePrice() *Price4 {
	f.IssuePrice = new(Price4)
	return f.IssuePrice
}

func (f *FutureOrOptionDetails1) AddOptionRights() *OptionRight1Choice {
	f.OptionRights = new(OptionRight1Choice)
	return f.OptionRights
}

func (f *FutureOrOptionDetails1) SetLastTransaction(value string) {
	f.LastTransaction = (*YesNoIndicator)(&value)
}

func (f *FutureOrOptionDetails1) SetSpreadTransaction(value string) {
	f.SpreadTransaction = (*YesNoIndicator)(&value)
}
