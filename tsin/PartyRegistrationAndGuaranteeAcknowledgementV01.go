package tsin

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01200101 struct {
	XMLName xml.Name                                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Document"`
	Message *PartyRegistrationAndGuaranteeAcknowledgementV01 `xml:"PtyRegnAndGrntAck"`
}

func (d *Document01200101) AddMessage() *PartyRegistrationAndGuaranteeAcknowledgementV01 {
	d.Message = new(PartyRegistrationAndGuaranteeAcknowledgementV01)
	return d.Message
}

// The message PartyManagementPaymentAcknowledgement is sent from a trade partner to any partner requested through a PartyManagementPaymentAcknowledgemenNotification message to acknowledge the notified factoring service agreement. Depending on legal contexts, the acknowledgement may be required in order for the financial service agreement to become effective.
// The message references related messages and may include referenced data.
// The message can carry digital signatures if required by context.
type PartyRegistrationAndGuaranteeAcknowledgementV01 struct {

	// Set of characteristics that unambiguously identify the acknowlegement, common parameters, documents and identifications.
	Header *model.BusinessLetter1 `xml:"Hdr"`

	// List of party management acknowledgements.
	AcknowledgementList []*model.FinancingAgreementList1 `xml:"AckList"`

	// Number of acknowledgement lists as control value.
	AcknowledgementCount *model.Max15NumericText `xml:"AckCnt"`

	// Total number of individual items in all lists.
	ItemCount *model.Max15NumericText `xml:"ItmCnt,omitempty"`

	// Total of all individual amounts included in all lists, irrespective of currencies or direction.
	ControlSum *model.DecimalNumber `xml:"CtrlSum,omitempty"`

	// Referenced or related business message.
	AttachedMessage []*model.EncapsulatedBusinessMessage1 `xml:"AttchdMsg,omitempty"`
}

func (p *PartyRegistrationAndGuaranteeAcknowledgementV01) AddHeader() *model.BusinessLetter1 {
	p.Header = new(model.BusinessLetter1)
	return p.Header
}

func (p *PartyRegistrationAndGuaranteeAcknowledgementV01) AddAcknowledgementList() *model.FinancingAgreementList1 {
	newValue := new(model.FinancingAgreementList1)
	p.AcknowledgementList = append(p.AcknowledgementList, newValue)
	return newValue
}

func (p *PartyRegistrationAndGuaranteeAcknowledgementV01) SetAcknowledgementCount(value string) {
	p.AcknowledgementCount = (*model.Max15NumericText)(&value)
}

func (p *PartyRegistrationAndGuaranteeAcknowledgementV01) SetItemCount(value string) {
	p.ItemCount = (*model.Max15NumericText)(&value)
}

func (p *PartyRegistrationAndGuaranteeAcknowledgementV01) SetControlSum(value string) {
	p.ControlSum = (*model.DecimalNumber)(&value)
}

func (p *PartyRegistrationAndGuaranteeAcknowledgementV01) AddAttachedMessage() *model.EncapsulatedBusinessMessage1 {
	newValue := new(model.EncapsulatedBusinessMessage1)
	p.AttachedMessage = append(p.AttachedMessage, newValue)
	return newValue
}
