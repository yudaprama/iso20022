package tsrv

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01000101 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.010.001.01 Document"`
	Message *UndertakingNonExtensionRequestV01 `xml:"UdrtkgNonXtnsnReq"`
}

func (d *Document01000101) AddMessage() *UndertakingNonExtensionRequestV01 {
	d.Message = new(UndertakingNonExtensionRequestV01)
	return d.Message
}

// The UndertakingNonExtensionRequest message is sent by the party that requested issuance of the undertaking (applicant or obligor) to the party that issued the undertaking. It is used to request no further automatic extensions to the expiry of the referenced undertaking.
type UndertakingNonExtensionRequestV01 struct {

	// Details of the non extension request.
	UndertakingNonExtensionRequestDetails *model.UndertakingNonExtensionRequest1 `xml:"UdrtkgNonXtnsnReqDtls"`

	// Digital signature of the request.
	DigitalSignature *model.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (u *UndertakingNonExtensionRequestV01) AddUndertakingNonExtensionRequestDetails() *model.UndertakingNonExtensionRequest1 {
	u.UndertakingNonExtensionRequestDetails = new(model.UndertakingNonExtensionRequest1)
	return u.UndertakingNonExtensionRequestDetails
}

func (u *UndertakingNonExtensionRequestV01) AddDigitalSignature() *model.PartyAndSignature2 {
	u.DigitalSignature = new(model.PartyAndSignature2)
	return u.DigitalSignature
}
