package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00900105 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.05 Document"`
	Message *RequestForTransferStatusReportV05 `xml:"ReqForTrfStsRpt"`
}

func (d *Document00900105) AddMessage() *RequestForTransferStatusReportV05 {
	d.Message = new(RequestForTransferStatusReportV05)
	return d.Message
}

// Scope
// An instructing party, for example, an investment manager or its authorised representative, sends the RequestForTransferStatusReport to the executing party, for example, a transfer agent to request the status of a previously instructed transfer.
// Usage
// The RequestForTransferStatusReport is used to request either:
// - the status of one or several transfer instructions or,
// - the status of one or several transfer cancellation instructions.
type RequestForTransferStatusReportV05 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Information to identify the transfer for which the status is requested.
	//
	RequestDetails []*model.MessageAndBusinessReference8 `xml:"ReqDtls"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *model.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (r *RequestForTransferStatusReportV05) AddMessageIdentification() *model.MessageIdentification1 {
	r.MessageIdentification = new(model.MessageIdentification1)
	return r.MessageIdentification
}

func (r *RequestForTransferStatusReportV05) AddRequestDetails() *model.MessageAndBusinessReference8 {
	newValue := new(model.MessageAndBusinessReference8)
	r.RequestDetails = append(r.RequestDetails, newValue)
	return newValue
}

func (r *RequestForTransferStatusReportV05) AddMarketPracticeVersion() *model.MarketPracticeVersion1 {
	r.MarketPracticeVersion = new(model.MarketPracticeVersion1)
	return r.MarketPracticeVersion
}

func (r *RequestForTransferStatusReportV05) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}
