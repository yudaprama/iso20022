package model

// Choice of the source (place) of the price quotation
type PriceSourceFormatChoice struct {

	// Source of price quotation is the market, expressed as a Market Identifier Code (MIC).
	LocalMarketPlace *MICIdentifier `xml:"LclMktPlc"`

	// Source of a price quotation when it is not the local market.
	NonLocalMarketPlace *PriceSource `xml:"NonLclMktPlc"`

	// Source of a price quotation expressed with a propriety identification scheme.
	PlaceAsDSS *GenericIdentification5 `xml:"PlcAsDSS"`
}

func (p *PriceSourceFormatChoice) SetLocalMarketPlace(value string) {
	p.LocalMarketPlace = (*MICIdentifier)(&value)
}

func (p *PriceSourceFormatChoice) AddNonLocalMarketPlace() *PriceSource {
	p.NonLocalMarketPlace = new(PriceSource)
	return p.NonLocalMarketPlace
}

func (p *PriceSourceFormatChoice) AddPlaceAsDSS() *GenericIdentification5 {
	p.PlaceAsDSS = new(GenericIdentification5)
	return p.PlaceAsDSS
}
