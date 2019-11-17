package pain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00600101 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.006.001.01 Document"`
	Message *PaymentCancellationRequestV01 `xml:"pain.006.001.01"`
}

func (d *Document00600101) AddMessage() *PaymentCancellationRequestV01 {
	d.Message = new(PaymentCancellationRequestV01)
	return d.Message
}

// Scope
// The PaymentCancellationRequest message is sent by the initiating party or any agent, to the next party in the payment chain. It is used to request the cancellation of an instruction previously sent.
// Usage
// The PaymentCancellationRequest message is exchanged between non-financial institution customers and agents to request the cancellation of a payment initiation message previously sent (i.e. the CustomerCreditTransferInitiation message and the CustomerDirectDebitInitiation messages).
// The PaymentCancellationRequest message can be used to request the cancellation of single instructions or multiple instructions, from one or multiple files.
// The PaymentCancellationRequest message can be used in domestic and cross-border scenarios.
// The PaymentCancellationRequest message refers to the original instruction(s) by means of references only or by means of references and a set of elements from the original instruction.
// The PaymentCancellationRequest message exchanged between non-financial institution customers and agents is identified in the schema as follows: urn:iso:std:iso:20022:tech:xsd:pain.006.001.01
type PaymentCancellationRequestV01 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader7 `xml:"GrpHdr"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *model.OriginalGroupInformation4 `xml:"OrgnlGrpInf"`

	// Information concerning the original transactions, to which the cancellation request message refers.
	TransactionInformation []*model.PaymentTransactionInformation3 `xml:"TxInf,omitempty"`
}

func (p *PaymentCancellationRequestV01) AddGroupHeader() *model.GroupHeader7 {
	p.GroupHeader = new(model.GroupHeader7)
	return p.GroupHeader
}

func (p *PaymentCancellationRequestV01) AddOriginalGroupInformation() *model.OriginalGroupInformation4 {
	p.OriginalGroupInformation = new(model.OriginalGroupInformation4)
	return p.OriginalGroupInformation
}

func (p *PaymentCancellationRequestV01) AddTransactionInformation() *model.PaymentTransactionInformation3 {
	newValue := new(model.PaymentTransactionInformation3)
	p.TransactionInformation = append(p.TransactionInformation, newValue)
	return newValue
}
