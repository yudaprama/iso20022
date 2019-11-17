package tsrv

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300101 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.003.001.01 Document"`
	Message *UndertakingIssuanceNotificationV01 `xml:"UdrtkgIssncNtfctn"`
}

func (d *Document00300101) AddMessage() *UndertakingIssuanceNotificationV01 {
	d.Message = new(UndertakingIssuanceNotificationV01)
	return d.Message
}

// The UndertakingIssuanceNotification message is sent by the party that issued the undertaking to the applicant to notify it of the contents of an undertaking issued electronically or on paper. The undertaking that is notified could be a demand guarantee, standby letter of credit, counter-undertaking (counter-guarantee or counter-standby), or suretyship undertaking. In addition to containing details on the applicable rules, expiry date, the amount, required documents, and terms and conditions of the undertaking, the message may provide information from the sender such as confirmation details.
type UndertakingIssuanceNotificationV01 struct {

	// Details related to the notification of the issued undertaking.
	UndertakingIssuanceNotificationDetails *model.UndertakingAdvice2 `xml:"UdrtkgIssncNtfctnDtls"`

	// Digital signature of the undertaking notification.
	DigitalSignature *model.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (u *UndertakingIssuanceNotificationV01) AddUndertakingIssuanceNotificationDetails() *model.UndertakingAdvice2 {
	u.UndertakingIssuanceNotificationDetails = new(model.UndertakingAdvice2)
	return u.UndertakingIssuanceNotificationDetails
}

func (u *UndertakingIssuanceNotificationV01) AddDigitalSignature() *model.PartyAndSignature2 {
	u.DigitalSignature = new(model.PartyAndSignature2)
	return u.DigitalSignature
}
