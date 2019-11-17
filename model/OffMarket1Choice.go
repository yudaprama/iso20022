package model

// Choice for specifying how the trade was executed off-market.
type OffMarket1Choice struct {

	// Indicates that the trade was executed off-exchange.
	OffMarketIndicator *OffMarket1Code `xml:"OffMktInd"`

	// Provides the BIC code of the systematic internaliser.
	SystematicInternaliser *AnyBICIdentifier `xml:"SystmtcIntlr"`
}

func (o *OffMarket1Choice) SetOffMarketIndicator(value string) {
	o.OffMarketIndicator = (*OffMarket1Code)(&value)
}

func (o *OffMarket1Choice) SetSystematicInternaliser(value string) {
	o.SystematicInternaliser = (*AnyBICIdentifier)(&value)
}
