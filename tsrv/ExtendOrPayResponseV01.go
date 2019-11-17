package tsrv

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01500101 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.015.001.01 Document"`
	Message *ExtendOrPayResponseV01 `xml:"XtndOrPayRspn"`
}

func (d *Document01500101) AddMessage() *ExtendOrPayResponseV01 {
	d.Message = new(ExtendOrPayResponseV01)
	return d.Message
}

// The ExtendOrPayResponse message is sent by the party that requested issuance of the undertaking (applicant or obligor) to the party that issued the undertaking, in response to the issuer's request for the applicant's response to the beneficiary’s request to extend or pay.
type ExtendOrPayResponseV01 struct {

	// Details of the extend or pay response.
	ExtendOrPayResponseDetails *model.ExtendOrPayQuery2 `xml:"XtndOrPayRspnDtls"`

	// Digital signature of the response.
	DigitalSignature *model.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (e *ExtendOrPayResponseV01) AddExtendOrPayResponseDetails() *model.ExtendOrPayQuery2 {
	e.ExtendOrPayResponseDetails = new(model.ExtendOrPayQuery2)
	return e.ExtendOrPayResponseDetails
}

func (e *ExtendOrPayResponseV01) AddDigitalSignature() *model.PartyAndSignature2 {
	e.DigitalSignature = new(model.PartyAndSignature2)
	return e.DigitalSignature
}
