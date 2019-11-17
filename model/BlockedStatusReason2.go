package model

// Reason for a blocked status.
type BlockedStatusReason2 struct {

	// Type of transaction for which the account has a blocked status.
	TransactionType *TransactionType5Choice `xml:"TxTp"`

	// Indicates whether the account is blocked.
	Blocked *YesNoIndicator `xml:"Blckd"`

	// Reason for the blocked status.
	Reason []*BlockedReason2Choice `xml:"Rsn,omitempty"`

	// Additional information about the blocked account status.
	AdditionalInformation *Max350Text `xml:"AddtlInf"`
}

func (b *BlockedStatusReason2) AddTransactionType() *TransactionType5Choice {
	b.TransactionType = new(TransactionType5Choice)
	return b.TransactionType
}

func (b *BlockedStatusReason2) SetBlocked(value string) {
	b.Blocked = (*YesNoIndicator)(&value)
}

func (b *BlockedStatusReason2) AddReason() *BlockedReason2Choice {
	newValue := new(BlockedReason2Choice)
	b.Reason = append(b.Reason, newValue)
	return newValue
}

func (b *BlockedStatusReason2) SetAdditionalInformation(value string) {
	b.AdditionalInformation = (*Max350Text)(&value)
}
