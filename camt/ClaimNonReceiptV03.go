package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02700103 struct {
	XMLName xml.Name            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Document"`
	Message *ClaimNonReceiptV03 `xml:"ClmNonRct"`
}

func (d *Document02700103) AddMessage() *ClaimNonReceiptV03 {
	d.Message = new(ClaimNonReceiptV03)
	return d.Message
}

// Scope
// The Claim Non Receipt message is sent by a case creator/case assigner to a case assignee.
// This message is used to initiate an investigation for missing funds at the creditor (missing credit entry to its account) or at an agent along the processing chain (missing cover for a received payment instruction).
// Usage
// The claim non receipt case occurs in two situations:
// - The creditor is expecting funds from a particular debtor and cannot find the corresponding credit entry on its account. In this situation, it is understood that the creditor will contact its debtor, and that the debtor will trigger the claim non receipt case on its behalf. A workflow where the creditor directly addresses a Claim Non Receipt message to its account servicing institution is not retained.
// - An agent in the processing chain cannot find a cover payment corresponding to a received payment instruction. In this situation, the agent may directly trigger the investigation by sending a Claim Non Receipt message to the sender of the original payment instruction.
// The Claim Non Receipt message covers one and only one payment instruction at a time. If several expected payment instructions/cover instructions are found missing, then multiple Claim Non Receipt messages must be sent.
// Depending on the result of the investigation by a case assignee (incorrect routing, errors/omissions when processing the instruction or even the absence of an error) and the stage at which the payment instruction is being process, the claim non receipt case may lead to a:
// - Request To Cancel Payment message, sent to the subsequent agent in the payment processing chain, if the original payment instruction has been incorrectly routed through the chain of agents. (This also implies that a new, corrected, payment instruction is issued).
// - Request To Modify Payment message, sent to the subsequent agent in the payment processing chain, if a truncation or omission has occurred during the processing of the original payment instruction.
// If the above situations occur, the assignee wanting to request a payment cancellation or payment modification should first send out a Resolution Of Investigation with a confirmation status that indicates that either cancellation (CWFW) modification (MWFW) or unable to apply (UWFW) will follow. (See section on Resolution Of Investigation for more details).
// In the cover is missing, the case assignee may also simply issue the omitted cover payment or when the initial cover information was incorrect, update the cover (through modification and/or cancellation as required) with the correction information provided in the ClaimNonReceipt message. The case assignee will issue a Resolution Of Investigation message with the CorrectionTransaction element mentioning the references of the cover payment.
// The Claim Non Receipt message may be forwarded to subsequent case assignees.
type ClaimNonReceiptV03 struct {

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *model.CaseAssignment2 `xml:"Assgnmt"`

	// Identifies the investigation case.
	Case *model.Case2 `xml:"Case"`

	// Identifies the payment instruction for which the Creditor has not received the funds.
	// Usage: In case of a missing cover, it must be the identification of the related payment instruction.
	// In case of a claim non receipt initiated by the debtor, it must be the identification of the instruction.
	Underlying *model.UnderlyingTransaction1Choice `xml:"Undrlyg"`

	// Provides the cover related information of a claim non receipt investigation. The absence of the component means that the message is not a cover related investigation.
	CoverDetails *model.MissingCover2 `xml:"CoverDtls,omitempty"`
}

func (c *ClaimNonReceiptV03) AddAssignment() *model.CaseAssignment2 {
	c.Assignment = new(model.CaseAssignment2)
	return c.Assignment
}

func (c *ClaimNonReceiptV03) AddCase() *model.Case2 {
	c.Case = new(model.Case2)
	return c.Case
}

func (c *ClaimNonReceiptV03) AddUnderlying() *model.UnderlyingTransaction1Choice {
	c.Underlying = new(model.UnderlyingTransaction1Choice)
	return c.Underlying
}

func (c *ClaimNonReceiptV03) AddCoverDetails() *model.MissingCover2 {
	c.CoverDetails = new(model.MissingCover2)
	return c.CoverDetails
}
