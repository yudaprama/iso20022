package colr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300103 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.003.001.03 Document"`
	Message *MarginCallRequestV03 `xml:"MrgnCallReq"`
}

func (d *Document00300103) AddMessage() *MarginCallRequestV03 {
	d.Message = new(MarginCallRequestV03)
	return d.Message
}

// Scope
// The MarginCallRequest message is sent by the collateral taker or its collateral manager to the collateral giver or its collateral manager
// This message is used to request new collateral at the initiation of an exposure or request additional collateral for an existing exposure. It can also be used to recall collateral upon the collateral giver or its collateral manager's request.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// When sent by the collateral taker the MarginCallRequest message is used to:
// - request new collateral at the initiation of an exposure
// - request additional collateral
// When sent by the collateral giver the MarginCallRequest message is used to:
// - request the return of collateral
type MarginCallRequestV03 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *model.Max35Text `xml:"TxId"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *model.Obligation3 `xml:"Oblgtn"`

	// Agreement details for the over the counter market.
	Agreement *model.Agreement2 `xml:"Agrmt,omitempty"`

	// Provides details about the margin calculation that would be due to party A.
	MarginDetailsDueToA *model.MarginCall1 `xml:"MrgnDtlsDueToA,omitempty"`

	// Provides details about the margin calculation that would be due to party B.
	MarginDetailsDueToB *model.MarginCall1 `xml:"MrgnDtlsDueToB,omitempty"`

	// Amount of expected margin that will be either delivered to party A by party B or recalled to party A from party B.
	RequirementDetailsDueToA *model.MarginRequirement1Choice `xml:"RqrmntDtlsDueToA,omitempty"`

	// Amount of expected margin that will be either delivered to party B by party A or recalled to party B from party A.
	RequirementDetailsDueToB *model.MarginRequirement1Choice `xml:"RqrmntDtlsDueToB,omitempty"`

	// Summation of the call amounts per margin type. It is provided for the purposes of carrying forward for future messages that are used to compare the margin call results to determine whether a call is agreed or full/partially disputed.
	MarginCallResult *model.MarginCallResult3 `xml:"MrgnCallRslt"`

	// Provides details about the type of collateral that will be either delivered to party B by party A or recalled to party B from party A.
	ExpectedCollateralDueToB *model.ExpectedCollateral1Choice `xml:"XpctdCollDueToB,omitempty"`

	// Provides details about the type of collateral that will be either delivered to party A by party B or recalled to party A from party B.
	ExpectedCollateralDueToA *model.ExpectedCollateral1Choice `xml:"XpctdCollDueToA,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MarginCallRequestV03) SetTransactionIdentification(value string) {
	m.TransactionIdentification = (*model.Max35Text)(&value)
}

func (m *MarginCallRequestV03) AddObligation() *model.Obligation3 {
	m.Obligation = new(model.Obligation3)
	return m.Obligation
}

func (m *MarginCallRequestV03) AddAgreement() *model.Agreement2 {
	m.Agreement = new(model.Agreement2)
	return m.Agreement
}

func (m *MarginCallRequestV03) AddMarginDetailsDueToA() *model.MarginCall1 {
	m.MarginDetailsDueToA = new(model.MarginCall1)
	return m.MarginDetailsDueToA
}

func (m *MarginCallRequestV03) AddMarginDetailsDueToB() *model.MarginCall1 {
	m.MarginDetailsDueToB = new(model.MarginCall1)
	return m.MarginDetailsDueToB
}

func (m *MarginCallRequestV03) AddRequirementDetailsDueToA() *model.MarginRequirement1Choice {
	m.RequirementDetailsDueToA = new(model.MarginRequirement1Choice)
	return m.RequirementDetailsDueToA
}

func (m *MarginCallRequestV03) AddRequirementDetailsDueToB() *model.MarginRequirement1Choice {
	m.RequirementDetailsDueToB = new(model.MarginRequirement1Choice)
	return m.RequirementDetailsDueToB
}

func (m *MarginCallRequestV03) AddMarginCallResult() *model.MarginCallResult3 {
	m.MarginCallResult = new(model.MarginCallResult3)
	return m.MarginCallResult
}

func (m *MarginCallRequestV03) AddExpectedCollateralDueToB() *model.ExpectedCollateral1Choice {
	m.ExpectedCollateralDueToB = new(model.ExpectedCollateral1Choice)
	return m.ExpectedCollateralDueToB
}

func (m *MarginCallRequestV03) AddExpectedCollateralDueToA() *model.ExpectedCollateral1Choice {
	m.ExpectedCollateralDueToA = new(model.ExpectedCollateral1Choice)
	return m.ExpectedCollateralDueToA
}

func (m *MarginCallRequestV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
