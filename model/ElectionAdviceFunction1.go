package model

// Provides information about the type of election advice and linked messages.
type ElectionAdviceFunction1 struct {

	// Type of election advice.
	ElectionType *ElectionType1Code `xml:"ElctnTp"`

	// Identification of the previous Agent Corporate Action Election Advice that is being amended.
	PreviousAgentCAElectionAdviceIdentification *DocumentIdentification8 `xml:"PrvsAgtCAElctnAdvcId,omitempty"`

	// Identification of the Agent Corporate Action Election Status Advice by which the issuer (agent) accepts the election amendment request.
	AgentCAElectionStatusAdviceIdentification *DocumentIdentification8 `xml:"AgtCAElctnStsAdvcId,omitempty"`

	// Identification of the Agent Corporate Action Election Amendment Request by which the CSD request the authorisation to amend an election.
	AgentCAElectionAmendmentRequestIdentification *DocumentIdentification8 `xml:"AgtCAElctnAmdmntReqId,omitempty"`
}

func (e *ElectionAdviceFunction1) SetElectionType(value string) {
	e.ElectionType = (*ElectionType1Code)(&value)
}

func (e *ElectionAdviceFunction1) AddPreviousAgentCAElectionAdviceIdentification() *DocumentIdentification8 {
	e.PreviousAgentCAElectionAdviceIdentification = new(DocumentIdentification8)
	return e.PreviousAgentCAElectionAdviceIdentification
}

func (e *ElectionAdviceFunction1) AddAgentCAElectionStatusAdviceIdentification() *DocumentIdentification8 {
	e.AgentCAElectionStatusAdviceIdentification = new(DocumentIdentification8)
	return e.AgentCAElectionStatusAdviceIdentification
}

func (e *ElectionAdviceFunction1) AddAgentCAElectionAmendmentRequestIdentification() *DocumentIdentification8 {
	e.AgentCAElectionAmendmentRequestIdentification = new(DocumentIdentification8)
	return e.AgentCAElectionAmendmentRequestIdentification
}
