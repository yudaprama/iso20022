package cain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300101 struct {
	XMLName xml.Name           `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 Document"`
	Message *AcquirerRejection `xml:"AcqrrRjctn"`
}

func (d *Document01300101) AddMessage() *AcquirerRejection {
	d.Message = new(AcquirerRejection)
	return d.Message
}

// The AcquirerRejection message is sent by any party, to reject an Acquirer to Issuer message.
type AcquirerRejection struct {

	// Information related to the protocol management.
	Header *model.Header19 `xml:"Hdr"`

	// Information related to the reject.
	Reject *model.AcceptorRejection4 `xml:"Rjct"`
}

func (a *AcquirerRejection) AddHeader() *model.Header19 {
	a.Header = new(model.Header19)
	return a.Header
}

func (a *AcquirerRejection) AddReject() *model.AcceptorRejection4 {
	a.Reject = new(model.AcceptorRejection4)
	return a.Reject
}
