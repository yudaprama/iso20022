package admi

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01700101 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:admi.017.001.01 Document"`
	Message *ProcessingRequestV01 `xml:"PrcgReq"`
}

func (d *Document01700101) AddMessage() *ProcessingRequestV01 {
	d.Message = new(ProcessingRequestV01)
	return d.Message
}

// The Processing Request message is sent by a participant to a central system to request the initiation of a system process suported by a central system.
type ProcessingRequestV01 struct {

	// Unique and unambiguous identifier for the message, as assigned by the sender.
	MessageIdentification *model.Max35Text `xml:"MsgId"`

	// Indicates the requested CLS Settlement Session that the related trade is part of.
	SettlementSessionIdentifier *model.Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// Contains the details of the processing request.
	Request *model.RequestDetails19 `xml:"Req"`
}

func (p *ProcessingRequestV01) SetMessageIdentification(value string) {
	p.MessageIdentification = (*model.Max35Text)(&value)
}

func (p *ProcessingRequestV01) SetSettlementSessionIdentifier(value string) {
	p.SettlementSessionIdentifier = (*model.Exact4AlphaNumericText)(&value)
}

func (p *ProcessingRequestV01) AddRequest() *model.RequestDetails19 {
	p.Request = new(model.RequestDetails19)
	return p.Request
}
