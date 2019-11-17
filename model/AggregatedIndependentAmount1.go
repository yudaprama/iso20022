package model

// Independent amount could be defined at a trade level or portfolio level.  It is assumed that their treatment will be based on the exposure convention that is whether netted together or treated on a gross basis.
type AggregatedIndependentAmount1 struct {

	// Total independent amount defined in the confirmations of individual trades.
	Trade *IndependentAmount1 `xml:"Trad,omitempty"`

	// Portfolio level independent amount that reflects portfolio change over a short time period using statistical techniques such as volatility and risk factor correlations.
	ValueAtRisk *IndependentAmount1 `xml:"ValAtRsk,omitempty"`

	// Portfolio level independent amount related to parties net open position. Net open position means the total of the net long FX and the net options in respect of each currency where: net long FX for any currency shall be the net amount (if any) of that currency which the party “A” is long as against party “B” in respect of all FX transactions.
	NetOpenPosition *IndependentAmount1 `xml:"NetOpnPos,omitempty"`

	// Any other amount that should be considered to calculate the independent amount.
	OtherAmount []*IndependentAmount2 `xml:"OthrAmt,omitempty"`
}

func (a *AggregatedIndependentAmount1) AddTrade() *IndependentAmount1 {
	a.Trade = new(IndependentAmount1)
	return a.Trade
}

func (a *AggregatedIndependentAmount1) AddValueAtRisk() *IndependentAmount1 {
	a.ValueAtRisk = new(IndependentAmount1)
	return a.ValueAtRisk
}

func (a *AggregatedIndependentAmount1) AddNetOpenPosition() *IndependentAmount1 {
	a.NetOpenPosition = new(IndependentAmount1)
	return a.NetOpenPosition
}

func (a *AggregatedIndependentAmount1) AddOtherAmount() *IndependentAmount2 {
	newValue := new(IndependentAmount2)
	a.OtherAmount = append(a.OtherAmount, newValue)
	return newValue
}
