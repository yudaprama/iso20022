package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document05600104 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.04 Document"`
	Message *FIToFIPaymentCancellationRequestV04 `xml:"FIToFIPmtCxlReq"`
}

func (d *Document05600104) AddMessage() *FIToFIPaymentCancellationRequestV04 {
	d.Message = new(FIToFIPaymentCancellationRequestV04)
	return d.Message
}

// Scope
// The FIToFI Payment Cancellation Request message is sent by a case creator/case assigner to a case assignee.
// This message is used to request the cancellation of an original payment instruction. The FIToFI Payment Cancellation Request message is exchanged between the instructing agent and the instructed agent to request the cancellation of a interbank payment message previously sent (such as FIToFICustomerCreditTransfer, FIToFICustomerDirectDebit or FinancialInstitutionCreditTransfer).
// Usage
// The FIToFI Payment Cancellation Request message must be answered with a:
// - Resolution Of Investigation message with a positive final outcome when the case assignee can perform the requested cancellation
// - Resolution Of Investigation message with a negative final outcome when the case assignee may perform the requested cancellation but fails to do so (too late, irrevocable instruction, ...)
// - Reject Investigation message when the case assignee is unable or not authorised to perform the requested cancellation
// - Notification Of Case Assignment message to indicate whether the case assignee will take on the case himself or reassign the case to a subsequent party in the payment processing chain.
// A FIToFI Payment Cancellation Request message concerns one and only one original payment instruction at a time.
// When a case assignee successfully performs a cancellation, it must return the corresponding funds to the case assigner. It may provide some details about the return in the Resolution Of Investigation message.
// The processing of a FIToFI Payment Cancellation Request message case may lead to a Debit Authorisation Request message sent to the creditor by its account servicing institution.
// The FIToFI Payment Cancellation Request message may be used to escalate a case after an unsuccessful request to modify the payment. In this scenario, the case identification remains the same as in the original FIToFI Payment Cancellation Request message and the element ReopenCaseIndication is set to 'Yes' or 'true'.
// The FIToFI Payment Cancellation Request message has the following main characteristics:
// Case Identification:
// The case creator assigns a unique case identification and the reason code for the cancellation request. This information will be passed unchanged to all subsequent case assignee(s). For the FIToFI Payment Cancellation Request message has been made optional, as the message might be used outside of a case management environment where the case identification is not relevant.
// Moreover, the case identification may be present at different levels:
// - One unique case is defined per cancellation request message: If multiple underlying groups or transactions are present in the message and the case assignee has already forwarded the transaction for which the cancellation is requested, the case cannot be forwarded to the next party in the chain (see rule on uniqueness of the case) and the case creator will have to issue individual cancellation requests for each underlying individual transaction. In response to this cancellation request, the case must also be present at the message level in the Resolution of Investigation message.
// - One case per original group or transaction present in the cancellation request: For each group or transaction, a unique case has been assigned. This means, when a payment instruction has already been forwarded by the case assignee, the cancellation request may be forwarded to next party in the payment chain, with the unique case assigned to the transaction. When the group can only be cancelled partially, new cancellation requests need however to be issued for the individual transactions within the group for which the cancellation request has not been successful. In response to this cancellation request, the case must be present in the cancellation details identifying the original group or transaction in the Resolution of Investigation message.
// - No case used in cancellation request message.
// Cancellation of a cover payment:
// The cancellation of a payment instruction for which cover is provided by a separate instruction always results in the cancellation of the whole transaction, including the cover. The case assignee performing the cancellation must initiate the return of funds to the case creator. The case assigner must not request the cancellation of the cover separately.
// Cancellation request initiators:
// The cancellation of a payment instruction can be initiated by either the debtor/creditor or any subsequent agent in the payment instruction processing chain.
type FIToFIPaymentCancellationRequestV04 struct {

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *model.CaseAssignment3 `xml:"Assgnmt"`

	// Identifies the investigation case.
	Case *model.Case3 `xml:"Case,omitempty"`

	// Provides details on the number of transactions and the control sum of the message.
	ControlData *model.ControlData1 `xml:"CtrlData,omitempty"`

	// Identifies the payment instruction to be cancelled.
	Underlying []*model.UnderlyingTransaction10 `xml:"Undrlyg"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (f *FIToFIPaymentCancellationRequestV04) AddAssignment() *model.CaseAssignment3 {
	f.Assignment = new(model.CaseAssignment3)
	return f.Assignment
}

func (f *FIToFIPaymentCancellationRequestV04) AddCase() *model.Case3 {
	f.Case = new(model.Case3)
	return f.Case
}

func (f *FIToFIPaymentCancellationRequestV04) AddControlData() *model.ControlData1 {
	f.ControlData = new(model.ControlData1)
	return f.ControlData
}

func (f *FIToFIPaymentCancellationRequestV04) AddUnderlying() *model.UnderlyingTransaction10 {
	newValue := new(model.UnderlyingTransaction10)
	f.Underlying = append(f.Underlying, newValue)
	return newValue
}

func (f *FIToFIPaymentCancellationRequestV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}
