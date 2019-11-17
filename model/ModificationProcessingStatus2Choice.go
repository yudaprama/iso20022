package model

// Choice of format for the processing status.
type ModificationProcessingStatus2Choice struct {

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus14Choice `xml:"AckdAccptd"`

	// The modification is pending processing.
	PendingProcessing *PendingProcessingStatus7Choice `xml:"PdgPrcg"`

	// Modification request will not be executed.
	Denied *DeniedStatus6Choice `xml:"Dnd"`

	// Modification request has been rejected for further processing.
	Rejected *RejectionStatus8Choice `xml:"Rjctd"`

	// Modification request is accepted but in repair.
	Repaired *RepairStatus8Choice `xml:"Rprd"`

	// Instruction has been modified.
	Modified *ModificationStatus2Choice `xml:"Modfd"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`
}

func (m *ModificationProcessingStatus2Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus14Choice {
	m.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus14Choice)
	return m.AcknowledgedAccepted
}

func (m *ModificationProcessingStatus2Choice) AddPendingProcessing() *PendingProcessingStatus7Choice {
	m.PendingProcessing = new(PendingProcessingStatus7Choice)
	return m.PendingProcessing
}

func (m *ModificationProcessingStatus2Choice) AddDenied() *DeniedStatus6Choice {
	m.Denied = new(DeniedStatus6Choice)
	return m.Denied
}

func (m *ModificationProcessingStatus2Choice) AddRejected() *RejectionStatus8Choice {
	m.Rejected = new(RejectionStatus8Choice)
	return m.Rejected
}

func (m *ModificationProcessingStatus2Choice) AddRepaired() *RepairStatus8Choice {
	m.Repaired = new(RepairStatus8Choice)
	return m.Repaired
}

func (m *ModificationProcessingStatus2Choice) AddModified() *ModificationStatus2Choice {
	m.Modified = new(ModificationStatus2Choice)
	return m.Modified
}

func (m *ModificationProcessingStatus2Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	m.Proprietary = new(ProprietaryStatusAndReason1)
	return m.Proprietary
}
