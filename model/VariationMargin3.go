package model

// Margin required to cover the risk because of the price fluctuations occurred on the unsettled exposures towards central counterparty.
type VariationMargin3 struct {

	// Provides details about the security identification.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Margin required to cover the risk because of the price fluctuations occurred on the unsettled exposures towards the central counterparty.
	TotalVariationMargin []*TotalVariationMargin1 `xml:"TtlVartnMrgn"`

	// Net unrealised profit or loss on the value of the netted, gross and failing positions.
	TotalMarkToMarket *Amount2 `xml:"TtlMrkToMkt"`

	// Unrealised net loss calculated at the participant portfolio level.
	MarkToMarketNetted []*Amount2 `xml:"MrkToMktNetd,omitempty"`

	// Unrealised net loss calculated in that market/boundary.
	MarkToMarketGross []*Amount2 `xml:"MrkToMktGrss,omitempty"`

	// Sum of the unrealised loss without taking profit into consideration.
	MarkToMarketFails []*Amount2 `xml:"MrkToMktFls,omitempty"`

	// Haircut applied to the absolute value of the participants net positions. Calculation depends on a participants credit rating.
	FailsHaircut *Amount2 `xml:"FlsHrcut,omitempty"`
}

func (v *VariationMargin3) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	v.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return v.FinancialInstrumentIdentification
}

func (v *VariationMargin3) AddTotalVariationMargin() *TotalVariationMargin1 {
	newValue := new(TotalVariationMargin1)
	v.TotalVariationMargin = append(v.TotalVariationMargin, newValue)
	return newValue
}

func (v *VariationMargin3) AddTotalMarkToMarket() *Amount2 {
	v.TotalMarkToMarket = new(Amount2)
	return v.TotalMarkToMarket
}

func (v *VariationMargin3) AddMarkToMarketNetted() *Amount2 {
	newValue := new(Amount2)
	v.MarkToMarketNetted = append(v.MarkToMarketNetted, newValue)
	return newValue
}

func (v *VariationMargin3) AddMarkToMarketGross() *Amount2 {
	newValue := new(Amount2)
	v.MarkToMarketGross = append(v.MarkToMarketGross, newValue)
	return newValue
}

func (v *VariationMargin3) AddMarkToMarketFails() *Amount2 {
	newValue := new(Amount2)
	v.MarkToMarketFails = append(v.MarkToMarketFails, newValue)
	return newValue
}

func (v *VariationMargin3) AddFailsHaircut() *Amount2 {
	v.FailsHaircut = new(Amount2)
	return v.FailsHaircut
}
