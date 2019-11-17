package colr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400104 struct {
	XMLName xml.Name               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.004.001.04 Document"`
	Message *MarginCallResponseV04 `xml:"MrgnCallRspn"`
}

func (d *Document00400104) AddMessage() *MarginCallResponseV04 {
	d.Message = new(MarginCallResponseV04)
	return d.Message
}

// Scope
// The MarginCallResponse message is sent by the collateral giver or its collateral manager to the collateral taker or its collateral manager or vice versa. This is a response to the MarginCallRequest message. The margin call can be accepted, fully disputed or partially disputed.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// When sent by the collateral giver the MarginCallResponse message is used to:
// - fully accept the MarginCallRequest
// - or partially accept the MarginCallRequest.
//
// When sent by the collateral taker the MarginCallResponse message is used to:
// - fully accept the recall of collateral
// - or partially accept the recall of collateral.
type MarginCallResponseV04 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *model.Max35Text `xml:"TxId"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *model.Obligation4 `xml:"Oblgtn"`

	// Agreement details for the over the counter market.
	Agreement *model.Agreement4 `xml:"Agrmt,omitempty"`

	// Provides details about the margin calculation that would be due to party A.
	MarginDetailsDueToA *model.MarginCall1 `xml:"MrgnDtlsDueToA,omitempty"`

	// Provides details about the margin calculation that would be due to party B.
	MarginDetailsDueToB *model.MarginCall1 `xml:"MrgnDtlsDueToB,omitempty"`

	// Provides details about the agreed amount that would be due to party A.
	AgreedAmountDueToA *model.AgreedAmount1Choice `xml:"AgrdAmtDueToA,omitempty"`

	// Provides details about the agreed amount that would be due to party B.
	AgreedAmountDueToB *model.AgreedAmount1Choice `xml:"AgrdAmtDueToB,omitempty"`

	// Provides response details about the margin call.
	ResponseDetails *model.Response1 `xml:"RspnDtls,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MarginCallResponseV04) SetTransactionIdentification(value string) {
	m.TransactionIdentification = (*model.Max35Text)(&value)
}

func (m *MarginCallResponseV04) AddObligation() *model.Obligation4 {
	m.Obligation = new(model.Obligation4)
	return m.Obligation
}

func (m *MarginCallResponseV04) AddAgreement() *model.Agreement4 {
	m.Agreement = new(model.Agreement4)
	return m.Agreement
}

func (m *MarginCallResponseV04) AddMarginDetailsDueToA() *model.MarginCall1 {
	m.MarginDetailsDueToA = new(model.MarginCall1)
	return m.MarginDetailsDueToA
}

func (m *MarginCallResponseV04) AddMarginDetailsDueToB() *model.MarginCall1 {
	m.MarginDetailsDueToB = new(model.MarginCall1)
	return m.MarginDetailsDueToB
}

func (m *MarginCallResponseV04) AddAgreedAmountDueToA() *model.AgreedAmount1Choice {
	m.AgreedAmountDueToA = new(model.AgreedAmount1Choice)
	return m.AgreedAmountDueToA
}

func (m *MarginCallResponseV04) AddAgreedAmountDueToB() *model.AgreedAmount1Choice {
	m.AgreedAmountDueToB = new(model.AgreedAmount1Choice)
	return m.AgreedAmountDueToB
}

func (m *MarginCallResponseV04) AddResponseDetails() *model.Response1 {
	m.ResponseDetails = new(model.Response1)
	return m.ResponseDetails
}

func (m *MarginCallResponseV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
