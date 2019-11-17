package model

// Identifies the transaction with a trade reference and provides its status. If the status is rejected, a reason for this status must be given.
type TradeTransactionStatusAndReason2 struct {

	// Provides the identification of the RegulatoryTransactionReportCancellationRequest document that was previously sent by the reporting institution.
	RelatedReference *Max35Text `xml:"RltdRef"`

	// Unique identification assigned to a trade once it is received or matched by an executing system.
	TradeReference *Max70Text `xml:"TradRef"`

	// Indicates the status of an instruction, request or report message.
	Status *Status2Code `xml:"Sts"`

	// Indicates that the cancellation is rejected and provides a reason why.
	Rejected []*RejectedCancellationStatusReason1Choice `xml:"Rjctd"`
}

func (t *TradeTransactionStatusAndReason2) SetRelatedReference(value string) {
	t.RelatedReference = (*Max35Text)(&value)
}

func (t *TradeTransactionStatusAndReason2) SetTradeReference(value string) {
	t.TradeReference = (*Max70Text)(&value)
}

func (t *TradeTransactionStatusAndReason2) SetStatus(value string) {
	t.Status = (*Status2Code)(&value)
}

func (t *TradeTransactionStatusAndReason2) AddRejected() *RejectedCancellationStatusReason1Choice {
	newValue := new(RejectedCancellationStatusReason1Choice)
	t.Rejected = append(t.Rejected, newValue)
	return newValue
}
