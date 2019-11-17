package model

// Describes the time-out consequences.
type TimeOutResult2 struct {

	// Specifies the status of the transaction if no action is taken by the user.
	TransactionFutureStatus *TransactionStatus5 `xml:"TxFutrSts"`
}

func (t *TimeOutResult2) AddTransactionFutureStatus() *TransactionStatus5 {
	t.TransactionFutureStatus = new(TransactionStatus5)
	return t.TransactionFutureStatus
}
