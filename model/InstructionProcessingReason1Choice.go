package model

// Choice of format for the instruction processing status.
type InstructionProcessingReason1Choice struct {

	// Specifies the reason of the RejectedStatus.
	Reason []*RejectionReason9 `xml:"Rsn"`

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`
}

func (i *InstructionProcessingReason1Choice) AddReason() *RejectionReason9 {
	newValue := new(RejectionReason9)
	i.Reason = append(i.Reason, newValue)
	return newValue
}

func (i *InstructionProcessingReason1Choice) SetNoSpecifiedReason(value string) {
	i.NoSpecifiedReason = (*NoReasonCode)(&value)
}
