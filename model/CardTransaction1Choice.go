package model

// Choice between a acquirer globalised card transaction or an individual card transaction.
type CardTransaction1Choice struct {

	// Card transaction details, based on card transaction aggregated data performed by the card acquirer.
	Aggregated *CardAggregated1 `xml:"Aggtd"`

	// Card transaction details for the individual transaction, as recorded at the POI (point of interaction).
	Individual *CardIndividualTransaction1 `xml:"Indv"`
}

func (c *CardTransaction1Choice) AddAggregated() *CardAggregated1 {
	c.Aggregated = new(CardAggregated1)
	return c.Aggregated
}

func (c *CardTransaction1Choice) AddIndividual() *CardIndividualTransaction1 {
	c.Individual = new(CardIndividualTransaction1)
	return c.Individual
}
