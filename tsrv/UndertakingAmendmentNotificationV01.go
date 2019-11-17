package tsrv

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700101 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.007.001.01 Document"`
	Message *UndertakingAmendmentNotificationV01 `xml:"UdrtkgAmdmntNtfctn"`
}

func (d *Document00700101) AddMessage() *UndertakingAmendmentNotificationV01 {
	d.Message = new(UndertakingAmendmentNotificationV01)
	return d.Message
}

// The UndertakingAmendmentNotification message is sent by the party that issued the undertaking to the applicant to notify it of the contents of a proposed amendment to the undertaking (issued electronically or on paper). The undertaking could be a demand guarantee, standby letter of credit, counter-undertaking (counter-guarantee or counter-standby), or suretyship undertaking. In addition to providing the proposed terms of the amendment and details on proposed changes to the undertaking, the message may provide other supporting information from the sender. It may also be used to notify the proposed termination or cancellation of the undertaking.
type UndertakingAmendmentNotificationV01 struct {

	// Details related to the notification of the proposed amended undertaking.
	UndertakingAmendmentNotificationDetails *model.Amendment6 `xml:"UdrtkgAmdmntNtfctnDtls"`

	// Digital signature of the notification.
	DigitalSignature *model.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (u *UndertakingAmendmentNotificationV01) AddUndertakingAmendmentNotificationDetails() *model.Amendment6 {
	u.UndertakingAmendmentNotificationDetails = new(model.Amendment6)
	return u.UndertakingAmendmentNotificationDetails
}

func (u *UndertakingAmendmentNotificationV01) AddDigitalSignature() *model.PartyAndSignature2 {
	u.DigitalSignature = new(model.PartyAndSignature2)
	return u.DigitalSignature
}
