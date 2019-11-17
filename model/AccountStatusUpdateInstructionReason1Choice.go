package model

// Choice of formats for a reason for an instruction to change the status of an account.
type AccountStatusUpdateInstructionReason1Choice struct {

	// There is no reason available or to report for the instruction to change the account status.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the instruction to change the account status.
	Reason []*AccountStatusUpdateInstructionReason1 `xml:"Rsn"`
}

func (a *AccountStatusUpdateInstructionReason1Choice) SetNoSpecifiedReason(value string) {
	a.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (a *AccountStatusUpdateInstructionReason1Choice) AddReason() *AccountStatusUpdateInstructionReason1 {
	newValue := new(AccountStatusUpdateInstructionReason1)
	a.Reason = append(a.Reason, newValue)
	return newValue
}
