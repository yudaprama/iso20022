package model

// Specifies the amount of a premium on a currency option together with its calculation method.
type PremiumQuote1Choice struct {

	// Premium calculation is based on a percentage of the call amount.
	PercentageOfCallAmount *PercentageRate `xml:"PctgOfCallAmt"`

	// Premium calculation is based on a percentage of the put amount.
	PercentageOfPutAmount *PercentageRate `xml:"PctgOfPutAmt"`

	// Premium calculation is based on points of the call amount.
	PointsOfCallAmount *BaseOneRate `xml:"PtsOfCallAmt"`

	// Premium calculation is based on points of the put amount.
	PointsOfPutAmount *BaseOneRate `xml:"PtsOfPutAmt"`
}

func (p *PremiumQuote1Choice) SetPercentageOfCallAmount(value string) {
	p.PercentageOfCallAmount = (*PercentageRate)(&value)
}

func (p *PremiumQuote1Choice) SetPercentageOfPutAmount(value string) {
	p.PercentageOfPutAmount = (*PercentageRate)(&value)
}

func (p *PremiumQuote1Choice) SetPointsOfCallAmount(value string) {
	p.PointsOfCallAmount = (*BaseOneRate)(&value)
}

func (p *PremiumQuote1Choice) SetPointsOfPutAmount(value string) {
	p.PointsOfPutAmount = (*BaseOneRate)(&value)
}
