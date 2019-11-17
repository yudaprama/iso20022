package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01200101 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.012.001.01 Document"`
	Message *AgentCAElectionAdviceV01 `xml:"AgtCAElctnAdvc"`
}

func (d *Document01200101) AddMessage() *AgentCAElectionAdviceV01 {
	d.Message = new(AgentCAElectionAdviceV01)
	return d.Message
}

// Scope
// This message is sent by a CSD to the issuer (or its agent) to provide information about the clients' election instruction, the registration details, the delivery details, etc. In case of an election with underlying resource movements, it also confirms that these have been completed. This message may also be sent in case of an amendment of an election, once the CSD has completed the required resource movements.
// Usage
// This message can be used for a new election advice or an amended election advice.
// If this message is used for a new election advice, the function of the message must be 'new election'.
// If this message is used for an amended election advice, the function of the message must be 'option change' and the identification of the previously sent election advice must be present.
// This message can include the cash movements and/or securities movements in the case of an election with underlying resource movements. Additionally, this message can include delivery, certification and beneficial owner details.
// Note: this information can be also sent separately in the Agent Corporate Action Information advice message.
type AgentCAElectionAdviceV01 struct {

	// Identification assigned by the Sender to unambiguously identify the advice.
	Identification *model.DocumentIdentification8 `xml:"Id"`

	// Provides information about the type of election advice and linked messages.
	ElectionAdviceTypeAndLinkage *model.ElectionAdviceFunction1 `xml:"ElctnAdvcTpAndLkg"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionInformation1 `xml:"CorpActnGnlInf"`

	// Provides information about the election(s).
	ElectionDetails *model.CorporateActionElection3 `xml:"ElctnDtls"`

	// Provides additional information about the delivery details, beneficial owner details, etc.
	AdditionalInformation *model.CorporateActionAdditionalInformation1 `xml:"AddtlInf,omitempty"`

	// Contact responsible for the transaction identified in the message.
	ContactDetails *model.ContactPerson1 `xml:"CtctDtls,omitempty"`
}

func (a *AgentCAElectionAdviceV01) AddIdentification() *model.DocumentIdentification8 {
	a.Identification = new(model.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCAElectionAdviceV01) AddElectionAdviceTypeAndLinkage() *model.ElectionAdviceFunction1 {
	a.ElectionAdviceTypeAndLinkage = new(model.ElectionAdviceFunction1)
	return a.ElectionAdviceTypeAndLinkage
}

func (a *AgentCAElectionAdviceV01) AddCorporateActionGeneralInformation() *model.CorporateActionInformation1 {
	a.CorporateActionGeneralInformation = new(model.CorporateActionInformation1)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCAElectionAdviceV01) AddElectionDetails() *model.CorporateActionElection3 {
	a.ElectionDetails = new(model.CorporateActionElection3)
	return a.ElectionDetails
}

func (a *AgentCAElectionAdviceV01) AddAdditionalInformation() *model.CorporateActionAdditionalInformation1 {
	a.AdditionalInformation = new(model.CorporateActionAdditionalInformation1)
	return a.AdditionalInformation
}

func (a *AgentCAElectionAdviceV01) AddContactDetails() *model.ContactPerson1 {
	a.ContactDetails = new(model.ContactPerson1)
	return a.ContactDetails
}
