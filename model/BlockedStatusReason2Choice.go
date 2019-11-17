package model

// Choice of formats for a blocked status reason.
type BlockedStatusReason2Choice struct {

	// There is no reason available or to report for the closed account status.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Transaction type for which the account is blocked and the underlying reason.
	Reason []*BlockedStatusReason2 `xml:"Rsn"`
}

func (b *BlockedStatusReason2Choice) SetNoSpecifiedReason(value string) {
	b.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (b *BlockedStatusReason2Choice) AddReason() *BlockedStatusReason2 {
	newValue := new(BlockedStatusReason2)
	b.Reason = append(b.Reason, newValue)
	return newValue
}
