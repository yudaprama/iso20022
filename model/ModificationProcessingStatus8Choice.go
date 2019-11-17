package model

// Choice of format for the processing status.
type ModificationProcessingStatus8Choice struct {

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus30Choice `xml:"AckdAccptd"`

	// Modification is pending processing.
	PendingProcessing *PendingProcessingStatus16Choice `xml:"PdgPrcg"`

	// Modification request will not be executed.
	Denied *DeniedStatus19Choice `xml:"Dnd"`

	// Modification request has been rejected for further processing.
	Rejected *RejectionStatus24Choice `xml:"Rjctd"`

	// Modification request is accepted but in repair.
	Repaired *RepairStatus17Choice `xml:"Rprd"`

	// Instruction has been modified.
	Modified *ModificationStatus5Choice `xml:"Modfd"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason7 `xml:"Prtry"`
}

func (m *ModificationProcessingStatus8Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus30Choice {
	m.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus30Choice)
	return m.AcknowledgedAccepted
}

func (m *ModificationProcessingStatus8Choice) AddPendingProcessing() *PendingProcessingStatus16Choice {
	m.PendingProcessing = new(PendingProcessingStatus16Choice)
	return m.PendingProcessing
}

func (m *ModificationProcessingStatus8Choice) AddDenied() *DeniedStatus19Choice {
	m.Denied = new(DeniedStatus19Choice)
	return m.Denied
}

func (m *ModificationProcessingStatus8Choice) AddRejected() *RejectionStatus24Choice {
	m.Rejected = new(RejectionStatus24Choice)
	return m.Rejected
}

func (m *ModificationProcessingStatus8Choice) AddRepaired() *RepairStatus17Choice {
	m.Repaired = new(RepairStatus17Choice)
	return m.Repaired
}

func (m *ModificationProcessingStatus8Choice) AddModified() *ModificationStatus5Choice {
	m.Modified = new(ModificationStatus5Choice)
	return m.Modified
}

func (m *ModificationProcessingStatus8Choice) AddProprietary() *ProprietaryStatusAndReason7 {
	m.Proprietary = new(ProprietaryStatusAndReason7)
	return m.Proprietary
}
