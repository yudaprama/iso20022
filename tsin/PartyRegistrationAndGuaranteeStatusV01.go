package tsin

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01000101 struct {
	XMLName xml.Name                                `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.010.001.01 Document"`
	Message *PartyRegistrationAndGuaranteeStatusV01 `xml:"PtyRegnAndGrntSts"`
}

func (d *Document01000101) AddMessage() *PartyRegistrationAndGuaranteeStatusV01 {
	d.Message = new(PartyRegistrationAndGuaranteeStatusV01)
	return d.Message
}

// The message PartyRegistrationAndGuaranteeStatus is either sent by a factoring service to a financing client to indicate the status of a factoring service agreement or from a guarantee issuer to a factoring client or a factoring service to indicate the guarantee covering a requested factoring service agreement. The message can also be sent to an interested party.
// The factoring service or guarantee issuer may include references to a corresponding PartyRegistrationAndGuaranteeRequest message or other related messages and may include referenced data.
// The message contains information about other parties to be notified about the financial service agreement or the guarantee and whether these parties are required to acknowledge the agreement.
// The message contains information returned by the financial institution indicating acceptance or rejection of the trade partner; a positive response is necessary before being able to assign financial items concerning the trade party.
// This message contains identifications of cash accounts to enable payer and payee to treat the transferred payment obligations.
// The message can carry digital signatures if required by context.
type PartyRegistrationAndGuaranteeStatusV01 struct {

	// Set of characteristics that unambiguously identify the status, common parameters, documents and identifications.
	Header *model.BusinessLetter1 `xml:"Hdr"`

	// List of agreements.
	AgreementList []*model.FinancingAgreementList1 `xml:"AgrmtList"`

	// Number of agreement lists as control value.
	AgreementCount *model.Max15NumericText `xml:"AgrmtCnt,omitempty"`

	// Total number of individual items in all lists.
	ItemCount *model.Max15NumericText `xml:"ItmCnt,omitempty"`

	// Total of all individual amounts included in all lists, irrespective of currencies or direction.
	ControlSum *model.DecimalNumber `xml:"CtrlSum,omitempty"`

	// Referenced or related business message.
	AttachedMessage []*model.EncapsulatedBusinessMessage1 `xml:"AttchdMsg,omitempty"`
}

func (p *PartyRegistrationAndGuaranteeStatusV01) AddHeader() *model.BusinessLetter1 {
	p.Header = new(model.BusinessLetter1)
	return p.Header
}

func (p *PartyRegistrationAndGuaranteeStatusV01) AddAgreementList() *model.FinancingAgreementList1 {
	newValue := new(model.FinancingAgreementList1)
	p.AgreementList = append(p.AgreementList, newValue)
	return newValue
}

func (p *PartyRegistrationAndGuaranteeStatusV01) SetAgreementCount(value string) {
	p.AgreementCount = (*model.Max15NumericText)(&value)
}

func (p *PartyRegistrationAndGuaranteeStatusV01) SetItemCount(value string) {
	p.ItemCount = (*model.Max15NumericText)(&value)
}

func (p *PartyRegistrationAndGuaranteeStatusV01) SetControlSum(value string) {
	p.ControlSum = (*model.DecimalNumber)(&value)
}

func (p *PartyRegistrationAndGuaranteeStatusV01) AddAttachedMessage() *model.EncapsulatedBusinessMessage1 {
	newValue := new(model.EncapsulatedBusinessMessage1)
	p.AttachedMessage = append(p.AttachedMessage, newValue)
	return newValue
}
