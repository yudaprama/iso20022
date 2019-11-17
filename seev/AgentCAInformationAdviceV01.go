package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02300101 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.023.001.01 Document"`
	Message *AgentCAInformationAdviceV01 `xml:"AgtCAInfAdvc"`
}

func (d *Document02300101) AddMessage() *AgentCAInformationAdviceV01 {
	d.Message = new(AgentCAInformationAdviceV01)
	return d.Message
}

// Scope
// This message is sent by a CSD to an issuer (or its agent) to provide information about the certification and/or the delivery details to the issuer (or its agent).
// Usage
// This message can be used in the case of a corporate action event without an election.
// This message can also be used in the case of a corporate action event with election when the election details and the additional information cannot be provided in the same message. In this case, the Agent Corporate Action Election Advice Identification must be used to link this message to the election advice for which additional information is provided.
type AgentCAInformationAdviceV01 struct {

	// Identification assigned by the Sender to unambiguously identify the advice.
	Identification *model.DocumentIdentification8 `xml:"Id"`

	// Identification of the linked Agent CA Election Advice Identification
	AgentCAElectionAdviceIdentification *model.DocumentIdentification8 `xml:"AgtCAElctnAdvcId,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionInformation1 `xml:"CorpActnGnlInf"`

	// Information about the account for which additional information is provided.
	AccountDetails *model.SecuritiesAccount7 `xml:"AcctDtls"`

	// Additional information about the corporate action such as the delivery details.
	CorporateActionAdditionalInformation *model.CorporateActionAdditionalInformation1 `xml:"CorpActnAddtlInf"`

	// Contact responsible for the transaction identified in the message.
	ContactDetails *model.ContactPerson1 `xml:"CtctDtls,omitempty"`
}

func (a *AgentCAInformationAdviceV01) AddIdentification() *model.DocumentIdentification8 {
	a.Identification = new(model.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCAInformationAdviceV01) AddAgentCAElectionAdviceIdentification() *model.DocumentIdentification8 {
	a.AgentCAElectionAdviceIdentification = new(model.DocumentIdentification8)
	return a.AgentCAElectionAdviceIdentification
}

func (a *AgentCAInformationAdviceV01) AddCorporateActionGeneralInformation() *model.CorporateActionInformation1 {
	a.CorporateActionGeneralInformation = new(model.CorporateActionInformation1)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCAInformationAdviceV01) AddAccountDetails() *model.SecuritiesAccount7 {
	a.AccountDetails = new(model.SecuritiesAccount7)
	return a.AccountDetails
}

func (a *AgentCAInformationAdviceV01) AddCorporateActionAdditionalInformation() *model.CorporateActionAdditionalInformation1 {
	a.CorporateActionAdditionalInformation = new(model.CorporateActionAdditionalInformation1)
	return a.CorporateActionAdditionalInformation
}

func (a *AgentCAInformationAdviceV01) AddContactDetails() *model.ContactPerson1 {
	a.ContactDetails = new(model.ContactPerson1)
	return a.ContactDetails
}
