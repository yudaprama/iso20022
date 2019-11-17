package model

// Choice of format for the processing status.
type ModificationProcessingStatus7Choice struct {

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus23Choice `xml:"AckdAccptd"`

	// Modification is pending processing.
	PendingProcessing *PendingProcessingStatus13Choice `xml:"PdgPrcg"`

	// Modification request will not be executed.
	Denied *DeniedStatus15Choice `xml:"Dnd"`

	// Modification request has been rejected for further processing.
	Rejected *RejectionStatus18Choice `xml:"Rjctd"`

	// Modification request is accepted but in repair.
	Repaired *RepairStatus13Choice `xml:"Rprd"`

	// Instruction has been modified.
	Modified *ModificationStatus4Choice `xml:"Modfd"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason6 `xml:"Prtry"`
}

func (m *ModificationProcessingStatus7Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus23Choice {
	m.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus23Choice)
	return m.AcknowledgedAccepted
}

func (m *ModificationProcessingStatus7Choice) AddPendingProcessing() *PendingProcessingStatus13Choice {
	m.PendingProcessing = new(PendingProcessingStatus13Choice)
	return m.PendingProcessing
}

func (m *ModificationProcessingStatus7Choice) AddDenied() *DeniedStatus15Choice {
	m.Denied = new(DeniedStatus15Choice)
	return m.Denied
}

func (m *ModificationProcessingStatus7Choice) AddRejected() *RejectionStatus18Choice {
	m.Rejected = new(RejectionStatus18Choice)
	return m.Rejected
}

func (m *ModificationProcessingStatus7Choice) AddRepaired() *RepairStatus13Choice {
	m.Repaired = new(RepairStatus13Choice)
	return m.Repaired
}

func (m *ModificationProcessingStatus7Choice) AddModified() *ModificationStatus4Choice {
	m.Modified = new(ModificationStatus4Choice)
	return m.Modified
}

func (m *ModificationProcessingStatus7Choice) AddProprietary() *ProprietaryStatusAndReason6 {
	m.Proprietary = new(ProprietaryStatusAndReason6)
	return m.Proprietary
}
