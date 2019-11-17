package model

// Return provided by a financial instrument.
type YieldCalculation2 struct {

	// Result of the yield calculation.
	Value *PercentageRate `xml:"Val"`

	// Specifies the type of calculation.
	CalculationType *CalculationType1Code `xml:"ClctnTp"`

	// Price to which the yield has been calculated.
	RedemptionPrice *Price4 `xml:"RedPric,omitempty"`

	// Date/time on which the calculation is based, for example, valuation on October 1 (price date) based on price of September 19 ( value date).
	ValueDate *ISODate `xml:"ValDt,omitempty"`

	// Period on which the calculation is based.
	ValuePeriod *DateTimePeriodChoice `xml:"ValPrd,omitempty"`

	// Included as needed to clarify yield irregularities associated with date, e.g. when it falls on a non-business day.
	CalculationDate *ISODate `xml:"ClctnDt,omitempty"`
}

func (y *YieldCalculation2) SetValue(value string) {
	y.Value = (*PercentageRate)(&value)
}

func (y *YieldCalculation2) SetCalculationType(value string) {
	y.CalculationType = (*CalculationType1Code)(&value)
}

func (y *YieldCalculation2) AddRedemptionPrice() *Price4 {
	y.RedemptionPrice = new(Price4)
	return y.RedemptionPrice
}

func (y *YieldCalculation2) SetValueDate(value string) {
	y.ValueDate = (*ISODate)(&value)
}

func (y *YieldCalculation2) AddValuePeriod() *DateTimePeriodChoice {
	y.ValuePeriod = new(DateTimePeriodChoice)
	return y.ValuePeriod
}

func (y *YieldCalculation2) SetCalculationDate(value string) {
	y.CalculationDate = (*ISODate)(&value)
}
