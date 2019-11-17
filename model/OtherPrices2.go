package model

// Specifies the type of price and information about the price.
type OtherPrices2 struct {

	// Specifies the maximum price.
	Maximum *Price4 `xml:"Max,omitempty"`

	// Specifies the transaction price.
	Transaction *Price4 `xml:"Tx,omitempty"`

	// Market price including or excluding the broker's commission.
	MarketBrokerCommission *Price4 `xml:"MktBrkrComssn,omitempty"`

	// In case of an order to buy, the price that the broker paid on the market plus the broker's commission.
	MarkedUp *Price4 `xml:"MrkdUp,omitempty"`

	// In case of an order to sell, the price the broker receives in the market minus the broker's commission.
	MarkedDown *Price4 `xml:"MrkdDwn,omitempty"`

	// Price is net to the disclosed client.
	NetDisclosed *Price4 `xml:"NetDscld,omitempty"`

	// Price is net to the client undisclosed (used in the UK market).
	NetUndisclosed *Price4 `xml:"NetUdscld,omitempty"`

	// Price is notional gross (used in the UK market).
	NotionalGross *Price4 `xml:"NtnlGrss,omitempty"`

	// Price is weighted average price of the benchmark prices at the time of each partial fill.
	BenchmarkWeightedAverage *Price4 `xml:"BchmkWghtdAvrg,omitempty"`

	// Price is weighted average price of all market executions during the completion of the order.
	AllMarketsWeightedAverage *Price4 `xml:"AllMktsWghtdAvrg,omitempty"`

	// Price is a benchmark price relating to the current partial fills (eg, last trade tick from market).
	Benchmark *Price4 `xml:"Bchmk,omitempty"`

	// Type of price that is not defined explicitly.
	OtherPrice *Price4 `xml:"OthrPric,omitempty"`

	// Price of securities representing a particular market or a portion of it.
	IndexPrice *Price4 `xml:"IndxPric,omitempty"`

	// Price used to differentiate from price on a confirmation of a marked up or marked down principal trade.
	ReportedPrice *Price4 `xml:"RptdPric,omitempty"`

	// Price of reference of the concerned financial instrument.
	ReferencePrice *PriceInformation11 `xml:"RefPric,omitempty"`
}

func (o *OtherPrices2) AddMaximum() *Price4 {
	o.Maximum = new(Price4)
	return o.Maximum
}

func (o *OtherPrices2) AddTransaction() *Price4 {
	o.Transaction = new(Price4)
	return o.Transaction
}

func (o *OtherPrices2) AddMarketBrokerCommission() *Price4 {
	o.MarketBrokerCommission = new(Price4)
	return o.MarketBrokerCommission
}

func (o *OtherPrices2) AddMarkedUp() *Price4 {
	o.MarkedUp = new(Price4)
	return o.MarkedUp
}

func (o *OtherPrices2) AddMarkedDown() *Price4 {
	o.MarkedDown = new(Price4)
	return o.MarkedDown
}

func (o *OtherPrices2) AddNetDisclosed() *Price4 {
	o.NetDisclosed = new(Price4)
	return o.NetDisclosed
}

func (o *OtherPrices2) AddNetUndisclosed() *Price4 {
	o.NetUndisclosed = new(Price4)
	return o.NetUndisclosed
}

func (o *OtherPrices2) AddNotionalGross() *Price4 {
	o.NotionalGross = new(Price4)
	return o.NotionalGross
}

func (o *OtherPrices2) AddBenchmarkWeightedAverage() *Price4 {
	o.BenchmarkWeightedAverage = new(Price4)
	return o.BenchmarkWeightedAverage
}

func (o *OtherPrices2) AddAllMarketsWeightedAverage() *Price4 {
	o.AllMarketsWeightedAverage = new(Price4)
	return o.AllMarketsWeightedAverage
}

func (o *OtherPrices2) AddBenchmark() *Price4 {
	o.Benchmark = new(Price4)
	return o.Benchmark
}

func (o *OtherPrices2) AddOtherPrice() *Price4 {
	o.OtherPrice = new(Price4)
	return o.OtherPrice
}

func (o *OtherPrices2) AddIndexPrice() *Price4 {
	o.IndexPrice = new(Price4)
	return o.IndexPrice
}

func (o *OtherPrices2) AddReportedPrice() *Price4 {
	o.ReportedPrice = new(Price4)
	return o.ReportedPrice
}

func (o *OtherPrices2) AddReferencePrice() *PriceInformation11 {
	o.ReferencePrice = new(PriceInformation11)
	return o.ReferencePrice
}
