package model

// Choice of format for the instruction processing status.
type InstructionProcessingReason2Choice struct {

	// Specifies the reason of the InRepairStatus.
	Reason []*RepairReason5 `xml:"Rsn"`

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`
}

func (i *InstructionProcessingReason2Choice) AddReason() *RepairReason5 {
	newValue := new(RepairReason5)
	i.Reason = append(i.Reason, newValue)
	return newValue
}

func (i *InstructionProcessingReason2Choice) SetNoSpecifiedReason(value string) {
	i.NoSpecifiedReason = (*NoReasonCode)(&value)
}
