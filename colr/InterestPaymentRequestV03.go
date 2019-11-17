package colr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300103 struct {
	XMLName xml.Name                   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.013.001.03 Document"`
	Message *InterestPaymentRequestV03 `xml:"IntrstPmtReq"`
}

func (d *Document01300103) AddMessage() *InterestPaymentRequestV03 {
	d.Message = new(InterestPaymentRequestV03)
	return d.Message
}

// Scope
// This InterestPaymentRequest message is sent by either;
// - the collateral taker or its collateral manager to the collateral giver or its collateral manager, or
// - the collateral giver or its collateral manager to the collateral taker or its collateral manager
// It is used to request payment or advise the amount due for interest calculated for a specified period. The interest is based on the amount of collateral that has been held during the calculation period. It is possible to send these messages on a bi-lateral basis for matching.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// The InterestPaymentRequest message is used to advise the interest amount calculated for a specific period or request payment for an interest amount for a specific period.
type InterestPaymentRequestV03 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *model.Max35Text `xml:"TxId"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *model.Obligation3 `xml:"Oblgtn"`

	// Agreement details for the over the counter market.
	Agreement *model.Agreement2 `xml:"Agrmt"`

	// Provides details on the interest amount due to party A.
	InterestDueToA *model.InterestAmount1 `xml:"IntrstDueToA,omitempty"`

	// Provides details on the interest amount due to party B.
	InterestDueToB *model.InterestAmount1 `xml:"IntrstDueToB,omitempty"`

	// Provides the net interest amount due to A or due to B in case of collateral held and posted during a period.
	NetAmountDetails *model.InterestResult1 `xml:"NetAmtDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *InterestPaymentRequestV03) SetTransactionIdentification(value string) {
	i.TransactionIdentification = (*model.Max35Text)(&value)
}

func (i *InterestPaymentRequestV03) AddObligation() *model.Obligation3 {
	i.Obligation = new(model.Obligation3)
	return i.Obligation
}

func (i *InterestPaymentRequestV03) AddAgreement() *model.Agreement2 {
	i.Agreement = new(model.Agreement2)
	return i.Agreement
}

func (i *InterestPaymentRequestV03) AddInterestDueToA() *model.InterestAmount1 {
	i.InterestDueToA = new(model.InterestAmount1)
	return i.InterestDueToA
}

func (i *InterestPaymentRequestV03) AddInterestDueToB() *model.InterestAmount1 {
	i.InterestDueToB = new(model.InterestAmount1)
	return i.InterestDueToB
}

func (i *InterestPaymentRequestV03) AddNetAmountDetails() *model.InterestResult1 {
	i.NetAmountDetails = new(model.InterestResult1)
	return i.NetAmountDetails
}

func (i *InterestPaymentRequestV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
