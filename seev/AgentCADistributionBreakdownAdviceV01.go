package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01600101 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.016.001.01 Document"`
	Message *AgentCADistributionBreakdownAdviceV01 `xml:"AgtCADstrbtnBrkdwnAdvc"`
}

func (d *Document01600101) AddMessage() *AgentCADistributionBreakdownAdviceV01 {
	d.Message = new(AgentCADistributionBreakdownAdviceV01)
	return d.Message
}

// Scope
// This message is sent by a CSD to an issuer (or its agent) to provide distribution breakdown information for the proceeds that are to be delivered outside the CSD (e.g. when the proceeds are not eligible in the CSD).
// Usage
// This message is used to provide distribution breakdown information (securities and/or cash) per account for a specific corporate action option.
// Note: the delivery details are sent through the Agent Corporate Action Information Advice.
type AgentCADistributionBreakdownAdviceV01 struct {

	// Identification assigned by the Sender to unambiguously identify the advice.
	Identification *model.DocumentIdentification8 `xml:"Id"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionInformation1 `xml:"CorpActnGnlInf"`

	// Provides information about the CA option and the entitlements.
	CorporateActionDistributionDetails *model.EntitlementAdvice1 `xml:"CorpActnDstrbtnDtls"`
}

func (a *AgentCADistributionBreakdownAdviceV01) AddIdentification() *model.DocumentIdentification8 {
	a.Identification = new(model.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCADistributionBreakdownAdviceV01) AddCorporateActionGeneralInformation() *model.CorporateActionInformation1 {
	a.CorporateActionGeneralInformation = new(model.CorporateActionInformation1)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCADistributionBreakdownAdviceV01) AddCorporateActionDistributionDetails() *model.EntitlementAdvice1 {
	a.CorporateActionDistributionDetails = new(model.EntitlementAdvice1)
	return a.CorporateActionDistributionDetails
}
