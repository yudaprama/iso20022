package model

// Calculation of the net asset value for an investment fund/fund class.
type PriceValuation4 struct {

	// Unique technical identifier for an instance of a price valuation within a price report, as assigned by the issuer of the report.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Date and time of the price valuation for the investment fund/fund class.
	ValuationDateTime *DateAndDateTimeChoice `xml:"ValtnDtTm,omitempty"`

	// Date and time at which a price is applied, according to the terms stated in the prospectus. The NAV date is also known as the trade date. The NAV date becomes the trade date in an order.
	NAVDateTime *DateAndDateTimeChoice `xml:"NAVDtTm"`

	// Investment fund class for which the net asset value is calculated.
	FinancialInstrumentDetails *FinancialInstrument8 `xml:"FinInstrmDtls"`

	// Issuer of the fund.
	FundManagementCompany *PartyIdentification2Choice `xml:"FndMgmtCpny,omitempty"`

	// Value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	TotalNAV []*ActiveOrHistoricCurrencyAndAmount `xml:"TtlNAV,omitempty"`

	// Total number of investment fund class units that have been issued.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb,omitempty"`

	// Date and time of the next price valuation for the investment fund/fund class.
	NextValuationDateTime *DateAndDateTimeChoice `xml:"NxtValtnDtTm,omitempty"`

	// Date and time of the previous price valuation for the investment fund/fund class.
	PreviousValuationDateTime *DateAndDateTimeChoice `xml:"PrvsValtnDtTm,omitempty"`

	// Specifies how the valuation is done, based on the schedule stated in the prospectus.
	ValuationType *ValuationTiming1Code `xml:"ValtnTp"`

	// Frequency of the valuation.
	ValuationFrequency *EventFrequency1Code `xml:"ValtnFrqcy,omitempty"`

	// Indicates whether the valuation is an official valuation.
	OfficialValuationIndicator *YesNoIndicator `xml:"OffclValtnInd"`

	// Indicates whether the valuation of the investment fund class is suspended.
	SuspendedIndicator *YesNoIndicator `xml:"SspdInd"`

	// Amount of money for which goods or services are offered, sold, or bought.
	PriceDetails []*UnitPrice15 `xml:"PricDtls,omitempty"`

	// Information related to the price variations of an investment fund class.
	ValuationStatistics []*ValuationStatistics3 `xml:"ValtnSttstcs,omitempty"`

	// Factors that give indications about the performance of a fund.
	PerformanceDetails *PerformanceFactors1 `xml:"PrfrmncDtls,omitempty"`
}

func (p *PriceValuation4) SetIdentification(value string) {
	p.Identification = (*Max35Text)(&value)
}

func (p *PriceValuation4) AddValuationDateTime() *DateAndDateTimeChoice {
	p.ValuationDateTime = new(DateAndDateTimeChoice)
	return p.ValuationDateTime
}

func (p *PriceValuation4) AddNAVDateTime() *DateAndDateTimeChoice {
	p.NAVDateTime = new(DateAndDateTimeChoice)
	return p.NAVDateTime
}

func (p *PriceValuation4) AddFinancialInstrumentDetails() *FinancialInstrument8 {
	p.FinancialInstrumentDetails = new(FinancialInstrument8)
	return p.FinancialInstrumentDetails
}

func (p *PriceValuation4) AddFundManagementCompany() *PartyIdentification2Choice {
	p.FundManagementCompany = new(PartyIdentification2Choice)
	return p.FundManagementCompany
}

func (p *PriceValuation4) AddTotalNAV(value, currency string) {
	p.TotalNAV = append(p.TotalNAV, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (p *PriceValuation4) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	p.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return p.TotalUnitsNumber
}

func (p *PriceValuation4) AddNextValuationDateTime() *DateAndDateTimeChoice {
	p.NextValuationDateTime = new(DateAndDateTimeChoice)
	return p.NextValuationDateTime
}

func (p *PriceValuation4) AddPreviousValuationDateTime() *DateAndDateTimeChoice {
	p.PreviousValuationDateTime = new(DateAndDateTimeChoice)
	return p.PreviousValuationDateTime
}

func (p *PriceValuation4) SetValuationType(value string) {
	p.ValuationType = (*ValuationTiming1Code)(&value)
}

func (p *PriceValuation4) SetValuationFrequency(value string) {
	p.ValuationFrequency = (*EventFrequency1Code)(&value)
}

func (p *PriceValuation4) SetOfficialValuationIndicator(value string) {
	p.OfficialValuationIndicator = (*YesNoIndicator)(&value)
}

func (p *PriceValuation4) SetSuspendedIndicator(value string) {
	p.SuspendedIndicator = (*YesNoIndicator)(&value)
}

func (p *PriceValuation4) AddPriceDetails() *UnitPrice15 {
	newValue := new(UnitPrice15)
	p.PriceDetails = append(p.PriceDetails, newValue)
	return newValue
}

func (p *PriceValuation4) AddValuationStatistics() *ValuationStatistics3 {
	newValue := new(ValuationStatistics3)
	p.ValuationStatistics = append(p.ValuationStatistics, newValue)
	return newValue
}

func (p *PriceValuation4) AddPerformanceDetails() *PerformanceFactors1 {
	p.PerformanceDetails = new(PerformanceFactors1)
	return p.PerformanceDetails
}
