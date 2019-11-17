package model

// Identifies the transaction with a trade reference and provides its status. If the status is rejected, a reason for this status must be given.
type TradeTransactionStatusAndReason1 struct {

	// Provides the identification of the RegulatoryTransactionReport document that was previously sent by the reporting institution.
	RelatedReference *Max35Text `xml:"RltdRef"`

	// Unique identification assigned to a trade once it is received or matched by an executing system.
	TradeReference *Max70Text `xml:"TradRef"`

	// Indicates the status of a trade transaction.
	Status *Status2Code `xml:"Sts"`

	// Indicates that the report is rejected and provides a reason why.
	Rejected []*RejectedStatusReason9Choice `xml:"Rjctd"`
}

func (t *TradeTransactionStatusAndReason1) SetRelatedReference(value string) {
	t.RelatedReference = (*Max35Text)(&value)
}

func (t *TradeTransactionStatusAndReason1) SetTradeReference(value string) {
	t.TradeReference = (*Max70Text)(&value)
}

func (t *TradeTransactionStatusAndReason1) SetStatus(value string) {
	t.Status = (*Status2Code)(&value)
}

func (t *TradeTransactionStatusAndReason1) AddRejected() *RejectedStatusReason9Choice {
	newValue := new(RejectedStatusReason9Choice)
	t.Rejected = append(t.Rejected, newValue)
	return newValue
}
