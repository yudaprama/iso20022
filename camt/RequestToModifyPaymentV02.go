package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document08700102 struct {
	XMLName xml.Name                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.087.001.02 Document"`
	Message *RequestToModifyPaymentV02 `xml:"ReqToModfyPmt"`
}

func (d *Document08700102) AddMessage() *RequestToModifyPaymentV02 {
	d.Message = new(RequestToModifyPaymentV02)
	return d.Message
}

// Scope
// The Request To Modify Payment message is sent by a case creator/case assigner to a case assignee.
// This message is used to request the modification of characteristics of an original payment instruction.
// Usage
// The Request To Modify Payment message must be answered with a:
// - Resolution Of Investigation message with a positive final outcome when the case assignee can perform the requested modification
// - Resolution Of Investigation message with a negative final outcome when the case assignee may perform the requested modification but fails to do so (too late, irrevocable instruction, one requested element cannot be modified, ...)
// - Reject Case Assignment message when the case assignee is unable or not authorised to perform the requested modification
// - Notification Of Case Assignment message to indicate whether the case assignee will take on the case himself or reassign the case to a subsequent party in the payment processing chain.
// The Request To Modify Payment message covers one and only one original instruction at a time. If several original payment instructions need to be modified, then multiple Request To Modify Payment messages must be sent.
// The Request To Modify Payment message can be sent to request the modification of one or several elements of the original payment instruction. If many elements need to be modified, it is recommended to cancel the original payment instruction and initiate a new one.
// The Request To Modify Payment must be processed on an all or nothing basis. If one of the elements to be modified cannot be altered, the assignment must be rejected in full by means of a negative Resolution Of Investigation message. (See section on Resolution Of Investigation for more details.)
// The Request To Modify Payment message must never be sent to request the modification of the currency of the original payment instruction. If the currency is wrong, use Request To Cancel Payment message to cancel it and issue and a new payment instruction.
// The Request To Modify Payment message may be forwarded to subsequent case assignee(s).
// When a Request To Modify Payment message is used to decrease the amount of the original payment instruction, the modification will trigger a return of funds from the case assignee to the case creator. The assignee may indicate, within the Resolution Of Investigation message, the amount to be returned, the date it is or will be returned and the channel through which the return will be done.
// The Request To Modify Payment message must never be sent to request the increase of the amount of the original payment instruction. To increase the amount in a payment, the debtor can do one of the following:
// - Cancel the first payment using a Request To Cancel Payment message and make a new payment with a higher and correct amount.
// - Simply send a second payment with the supplementary amount.
// Depending on the requested modification(s) and the processing stage of the original payment instruction, the processing of a request to modify payment case may end with one of the following:
// - an Additional Payment Information message sent to the creditor of the original payment instruction
// - a Debit Authorisation Request message sent to the creditor of the original payment instruction
// - a Request To Cancel Payment message sent to a subsequent case assignee
// The Request To Modify Payment message can be sent to correct characteristics of an original payment instruction following receipt of an Unable To Apply message. In this scenario, the case identification will remain the same.
// The RequestToModifyPayment message has the following main characteristics:
// The case creator assigns a unique case identification. This information will be passed unchanged to all subsequent case assignee(s).
// Lowering the amount of an original payment instruction for which cover is provided by a separate instruction will systematically mean the modification of the whole transaction, including the cover. The case assignee performing the amount modification must initiate the return of funds in excess to the case creator.
// The modification of the agent's or agents' information on an original payment instruction for which cover is provided by a separate instruction will systematically mean the whole transaction is modified, i.e., the cover is executed through the agent(s) mentioned in the Request To Modify Payment message. The cover payment must not be modified separately.
// The modification of a payment instruction can be initiated by either the debtor or any subsequent agent in the payment processing chain.
// The case creator provides the information to be modified in line with agreements made with the case assignee. If the case assignee needs in turn to assign the case to a subsequent case assignee, the requested modification(s) must be in line with the agreement made with the next case assignee and a Notification Of Case Assignment message must be sent to the case assigner. Otherwise, the request to modify payment case must be rejected (by means of a negative Resolution Of Investigation message).
type RequestToModifyPaymentV02 struct {

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *model.CaseAssignment3 `xml:"Assgnmt"`

	// Identifies the investigation case.
	Case *model.Case3 `xml:"Case"`

	// Identifies the payment transaction to be modified.
	Underlying *model.UnderlyingTransaction2Choice `xml:"Undrlyg"`

	// Identifies the list of modifications requested.
	Modification *model.RequestedModification4 `xml:"Mod"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (r *RequestToModifyPaymentV02) AddAssignment() *model.CaseAssignment3 {
	r.Assignment = new(model.CaseAssignment3)
	return r.Assignment
}

func (r *RequestToModifyPaymentV02) AddCase() *model.Case3 {
	r.Case = new(model.Case3)
	return r.Case
}

func (r *RequestToModifyPaymentV02) AddUnderlying() *model.UnderlyingTransaction2Choice {
	r.Underlying = new(model.UnderlyingTransaction2Choice)
	return r.Underlying
}

func (r *RequestToModifyPaymentV02) AddModification() *model.RequestedModification4 {
	r.Modification = new(model.RequestedModification4)
	return r.Modification
}

func (r *RequestToModifyPaymentV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	r.SupplementaryData = append(r.SupplementaryData, newValue)
	return newValue
}
