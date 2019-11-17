package model

// Choice between a acquirer globalised card transaction or an individual card transaction.
type CardTransaction2Choice struct {

	// Card transaction details, based on card transaction aggregated data performed by the card acquirer.
	Aggregated *CardAggregated1 `xml:"Aggtd"`

	// Card transaction details for the individual transaction, as recorded at the POI (point of interaction).
	Individual *CardIndividualTransaction2 `xml:"Indv"`
}

func (c *CardTransaction2Choice) AddAggregated() *CardAggregated1 {
	c.Aggregated = new(CardAggregated1)
	return c.Aggregated
}

func (c *CardTransaction2Choice) AddIndividual() *CardIndividualTransaction2 {
	c.Individual = new(CardIndividualTransaction2)
	return c.Individual
}
