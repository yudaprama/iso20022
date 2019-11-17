package model

// Context of the card transaction.
type CardTransactionContext4 struct {

	// Data used to assign specific condition such as liability shift or preferential interchange fees.
	SpecialConditions []*CardTransactionCondition1 `xml:"SpclConds,omitempty"`
}

func (c *CardTransactionContext4) AddSpecialConditions() *CardTransactionCondition1 {
	newValue := new(CardTransactionCondition1)
	c.SpecialConditions = append(c.SpecialConditions, newValue)
	return newValue
}
