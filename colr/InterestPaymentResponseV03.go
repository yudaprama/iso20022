package colr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01400103 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 Document"`
	Message *InterestPaymentResponseV03 `xml:"IntrstPmtRspn"`
}

func (d *Document01400103) AddMessage() *InterestPaymentResponseV03 {
	d.Message = new(InterestPaymentResponseV03)
	return d.Message
}

// Scope
// This InterestPaymentResponse message is sent by either;
// - the collateral taker or its collateral manager to the collateral giver or its collateral manager, or
// - the collateral giver or its collateral manager to the collateral taker or its collateral manager
// This is a response to the InterestPaymentRequest message and the amount of interest requested or advised can be accepted or rejected.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// The InterestPaymentResponse message is sent in response to the InterestPaymentRequest in order to accept or reject the amount of interest requested or advised. A rejection reason and information can be provide if the InterestPaymentRequest is being rejected.
type InterestPaymentResponseV03 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *model.Max35Text `xml:"TxId"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *model.Obligation3 `xml:"Oblgtn"`

	// Agreement details for the over the counter market.
	Agreement *model.Agreement2 `xml:"Agrmt"`

	// Provides details on the interest amount due to party A.
	InterestDueToA *model.InterestAmount2 `xml:"IntrstDueToA,omitempty"`

	// Provides details on the interest amount due to party B.
	InterestDueToB *model.InterestAmount2 `xml:"IntrstDueToB,omitempty"`

	// Provides details on the response to the interest payment request.
	InterestResponse *model.InterestResponse1 `xml:"IntrstRspn"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *InterestPaymentResponseV03) SetTransactionIdentification(value string) {
	i.TransactionIdentification = (*model.Max35Text)(&value)
}

func (i *InterestPaymentResponseV03) AddObligation() *model.Obligation3 {
	i.Obligation = new(model.Obligation3)
	return i.Obligation
}

func (i *InterestPaymentResponseV03) AddAgreement() *model.Agreement2 {
	i.Agreement = new(model.Agreement2)
	return i.Agreement
}

func (i *InterestPaymentResponseV03) AddInterestDueToA() *model.InterestAmount2 {
	i.InterestDueToA = new(model.InterestAmount2)
	return i.InterestDueToA
}

func (i *InterestPaymentResponseV03) AddInterestDueToB() *model.InterestAmount2 {
	i.InterestDueToB = new(model.InterestAmount2)
	return i.InterestDueToB
}

func (i *InterestPaymentResponseV03) AddInterestResponse() *model.InterestResponse1 {
	i.InterestResponse = new(model.InterestResponse1)
	return i.InterestResponse
}

func (i *InterestPaymentResponseV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
