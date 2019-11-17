package model

// Calculation of the net asset value for an investment fund/fund class.
type PriceValuation2 struct {

	// Unique technical identifier for an instance of a price valuation within a price report, as assigned by the issuer of the report.
	Identification *Max35Text `xml:"Id"`

	// Date and time of the price valuation for the investment fund/fund class.
	ValuationDateTime *DateAndDateTimeChoice `xml:"ValtnDtTm,omitempty"`

	// Date and time at which a price is applied, according to the terms stated in the prospectus.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Investment fund class for which the net asset value is calculated.
	FinancialInstrumentDetails *FinancialInstrument5 `xml:"FinInstrmDtls"`

	// Value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	TotalNAV []*ActiveOrHistoricCurrencyAndAmount `xml:"TtlNAV,omitempty"`

	// Total number of investment fund class units that have been issued.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb,omitempty"`

	// Date and time of the next price valuation for the investment fund/fund class.
	NextValuationDateTime *DateAndDateTimeChoice `xml:"NxtValtnDtTm,omitempty"`

	// Date and time of the previous price valuation for the investment fund/fund class.
	PreviousValuationDateTime *DateAndDateTimeChoice `xml:"PrvsValtnDtTm,omitempty"`

	// Specifies how the valuation is done, based on the schedule stated in the prospectus.
	ValuationCycle *ValuationTiming1Code `xml:"ValtnCycl"`

	// Indicates whether the valuation of the investment fund class is suspended.
	SuspendedIndicator *YesNoIndicator `xml:"SspdInd"`

	// Amount of money for which goods or services are offered, sold, or bought.
	PriceDetails []*UnitPrice6 `xml:"PricDtls,omitempty"`

	// Information related to the price variations of an investment fund class.
	ValuationStatistics []*ValuationStatistics2 `xml:"ValtnSttstcs,omitempty"`
}

func (p *PriceValuation2) SetIdentification(value string) {
	p.Identification = (*Max35Text)(&value)
}

func (p *PriceValuation2) AddValuationDateTime() *DateAndDateTimeChoice {
	p.ValuationDateTime = new(DateAndDateTimeChoice)
	return p.ValuationDateTime
}

func (p *PriceValuation2) AddTradeDateTime() *DateAndDateTimeChoice {
	p.TradeDateTime = new(DateAndDateTimeChoice)
	return p.TradeDateTime
}

func (p *PriceValuation2) AddFinancialInstrumentDetails() *FinancialInstrument5 {
	p.FinancialInstrumentDetails = new(FinancialInstrument5)
	return p.FinancialInstrumentDetails
}

func (p *PriceValuation2) AddTotalNAV(value, currency string) {
	p.TotalNAV = append(p.TotalNAV, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (p *PriceValuation2) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	p.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return p.TotalUnitsNumber
}

func (p *PriceValuation2) AddNextValuationDateTime() *DateAndDateTimeChoice {
	p.NextValuationDateTime = new(DateAndDateTimeChoice)
	return p.NextValuationDateTime
}

func (p *PriceValuation2) AddPreviousValuationDateTime() *DateAndDateTimeChoice {
	p.PreviousValuationDateTime = new(DateAndDateTimeChoice)
	return p.PreviousValuationDateTime
}

func (p *PriceValuation2) SetValuationCycle(value string) {
	p.ValuationCycle = (*ValuationTiming1Code)(&value)
}

func (p *PriceValuation2) SetSuspendedIndicator(value string) {
	p.SuspendedIndicator = (*YesNoIndicator)(&value)
}

func (p *PriceValuation2) AddPriceDetails() *UnitPrice6 {
	newValue := new(UnitPrice6)
	p.PriceDetails = append(p.PriceDetails, newValue)
	return newValue
}

func (p *PriceValuation2) AddValuationStatistics() *ValuationStatistics2 {
	newValue := new(ValuationStatistics2)
	p.ValuationStatistics = append(p.ValuationStatistics, newValue)
	return newValue
}
