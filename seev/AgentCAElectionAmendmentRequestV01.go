package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300101 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.013.001.01 Document"`
	Message *AgentCAElectionAmendmentRequestV01 `xml:"AgtCAElctnAmdmntReq"`
}

func (d *Document01300101) AddMessage() *AgentCAElectionAmendmentRequestV01 {
	d.Message = new(AgentCAElectionAmendmentRequestV01)
	return d.Message
}

// Scope
// This message is sent by a CSD to the issuer (or its agent) to request the authorisation of an amendment of a previously sent Agent Corporate Action Election Advice message.
// Usage
// This message is used to request the amendment of a previously sent Agent Corporate Action Election Advice message.
// Once the amendment request has been accepted by the issuer (or its agent), the CSD will process any resource movement and send an Agent Corporate Action Election Advice message with the function, option change, to confirm that the amendment has been booked at the CSD.
// This message is used when the terms and conditions of the corporate action event allow amendments.
type AgentCAElectionAmendmentRequestV01 struct {

	// Identification assigned by the Sender to unambiguously identify the request.
	Identification *model.DocumentIdentification8 `xml:"Id"`

	// Identification of the linked Agent CA Election Advice for which an amendment is requested.
	AgentCAElectionAdviceIdentification *model.DocumentIdentification8 `xml:"AgtCAElctnAdvcId"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionInformation1 `xml:"CorpActnGnlInf"`

	// Provides information about the account.
	AccountDetails *model.SecuritiesAccount7 `xml:"AcctDtls"`

	// Provides information about the original election to be amended.
	OriginalElectionDetails *model.CorporateActionElection1 `xml:"OrgnlElctnDtls"`

	// Provides information about the amendments to the election.
	AmendedElectionDetails *model.CorporateActionElection2 `xml:"AmddElctnDtls"`

	// Contact responsible for the transaction identified in the message.
	ContactDetails *model.ContactPerson1 `xml:"CtctDtls,omitempty"`
}

func (a *AgentCAElectionAmendmentRequestV01) AddIdentification() *model.DocumentIdentification8 {
	a.Identification = new(model.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCAElectionAmendmentRequestV01) AddAgentCAElectionAdviceIdentification() *model.DocumentIdentification8 {
	a.AgentCAElectionAdviceIdentification = new(model.DocumentIdentification8)
	return a.AgentCAElectionAdviceIdentification
}

func (a *AgentCAElectionAmendmentRequestV01) AddCorporateActionGeneralInformation() *model.CorporateActionInformation1 {
	a.CorporateActionGeneralInformation = new(model.CorporateActionInformation1)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCAElectionAmendmentRequestV01) AddAccountDetails() *model.SecuritiesAccount7 {
	a.AccountDetails = new(model.SecuritiesAccount7)
	return a.AccountDetails
}

func (a *AgentCAElectionAmendmentRequestV01) AddOriginalElectionDetails() *model.CorporateActionElection1 {
	a.OriginalElectionDetails = new(model.CorporateActionElection1)
	return a.OriginalElectionDetails
}

func (a *AgentCAElectionAmendmentRequestV01) AddAmendedElectionDetails() *model.CorporateActionElection2 {
	a.AmendedElectionDetails = new(model.CorporateActionElection2)
	return a.AmendedElectionDetails
}

func (a *AgentCAElectionAmendmentRequestV01) AddContactDetails() *model.ContactPerson1 {
	a.ContactDetails = new(model.ContactPerson1)
	return a.ContactDetails
}
