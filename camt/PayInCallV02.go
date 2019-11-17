package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document06100102 struct {
	XMLName xml.Name      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.061.001.02 Document"`
	Message *PayInCallV02 `xml:"PayInCall"`
}

func (d *Document06100102) AddMessage() *PayInCallV02 {
	d.Message = new(PayInCallV02)
	return d.Message
}

// The PayInCall message is sent by a central settlement system to request additional funding from a settlement member impacted by a failure situation.
type PayInCallV02 struct {

	// Party for which the PayInCall is generated.
	PartyIdentification *model.PartyIdentification73Choice `xml:"PtyId"`

	// Contains  the report generation information and the report items.
	ReportData *model.ReportData5 `xml:"RptData"`

	// To indicate the requested CLS Settlement Session that the related trade is part of.
	SettlementSessionIdentifier *model.Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PayInCallV02) AddPartyIdentification() *model.PartyIdentification73Choice {
	p.PartyIdentification = new(model.PartyIdentification73Choice)
	return p.PartyIdentification
}

func (p *PayInCallV02) AddReportData() *model.ReportData5 {
	p.ReportData = new(model.ReportData5)
	return p.ReportData
}

func (p *PayInCallV02) SetSettlementSessionIdentifier(value string) {
	p.SettlementSessionIdentifier = (*model.Exact4AlphaNumericText)(&value)
}

func (p *PayInCallV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}
