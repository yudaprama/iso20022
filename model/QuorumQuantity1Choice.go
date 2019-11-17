package model

// Choice of quorum quantity.
type QuorumQuantity1Choice struct {

	// Minimum quantity of securities required to hold a meeting.
	QuorumQuantity *Max35Text `xml:"QrmQty"`

	// Minimum quantity of securities, expressed as a percentage, required to hold a meeting.
	QuorumQuantityPercentage *PercentageRate `xml:"QrmQtyPctg"`
}

func (q *QuorumQuantity1Choice) SetQuorumQuantity(value string) {
	q.QuorumQuantity = (*Max35Text)(&value)
}

func (q *QuorumQuantity1Choice) SetQuorumQuantityPercentage(value string) {
	q.QuorumQuantityPercentage = (*PercentageRate)(&value)
}
