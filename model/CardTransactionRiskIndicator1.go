package model

// Indicates to the issuer the level of risk of the transaction.
type CardTransactionRiskIndicator1 struct {

	// Reason to indicate a certain level of risk for the transaction.
	Reason []*CardTransactionRiskReason1Code `xml:"Rsn,omitempty"`

	// Level of risk, from 1 to 99.
	Level *Number `xml:"Lvl"`

	// Recommended action for the issuer.
	RecommendedAction []*ActionType4Code `xml:"RcmmnddActn,omitempty"`
}

func (c *CardTransactionRiskIndicator1) AddReason(value string) {
	c.Reason = append(c.Reason, (*CardTransactionRiskReason1Code)(&value))
}

func (c *CardTransactionRiskIndicator1) SetLevel(value string) {
	c.Level = (*Number)(&value)
}

func (c *CardTransactionRiskIndicator1) AddRecommendedAction(value string) {
	c.RecommendedAction = append(c.RecommendedAction, (*ActionType4Code)(&value))
}
