package secl

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400103 struct {
	XMLName xml.Name        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Document"`
	Message *NetPositionV03 `xml:"NetPos"`
}

func (d *Document00400103) AddMessage() *NetPositionV03 {
	d.Message = new(NetPositionV03)
	return d.Message
}

// Scope
// The Net Position Report message is sent by the central counterparty (CCP) to a clearing member to confirm the net position of all trade legs reported during the day.
//
// The message definition is intended for use with the ISO 20022 Business Application Header.
//
// Usage
// The central counterparty (CCP) nets all the positions per clearing account and sends the Net Position report message to the Clearing member.
type NetPositionV03 struct {

	// Provides parameters of the margin report such as the creation date and time, the report currency or the calculation date and time.
	ReportParameters *model.ReportParameters1 `xml:"RptParams"`

	// Provides information about the number of used pages.
	Pagination *model.Pagination `xml:"Pgntn"`

	// Provides the identification of the account owner, that is the clearing member (individual clearing member or general clearing member).
	ClearingMember *model.PartyIdentification35Choice `xml:"ClrMmb"`

	// Clearing organisation that will clear the trade.
	//
	// Note: This field allows Clearing Member Firm to segregate flows coming from clearing counterparty's clearing system. Indeed, Clearing Member Firms receive messages from the same system (same sender) and this field allows them to know if the message is related to equities or derivatives.
	ClearingSegment *model.PartyIdentification35Choice `xml:"ClrSgmt,omitempty"`

	// Provides the net position details such as the average deal price and net quantity.
	NetPositionReport []*model.NetPosition3 `xml:"NetPosRpt"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (n *NetPositionV03) AddReportParameters() *model.ReportParameters1 {
	n.ReportParameters = new(model.ReportParameters1)
	return n.ReportParameters
}

func (n *NetPositionV03) AddPagination() *model.Pagination {
	n.Pagination = new(model.Pagination)
	return n.Pagination
}

func (n *NetPositionV03) AddClearingMember() *model.PartyIdentification35Choice {
	n.ClearingMember = new(model.PartyIdentification35Choice)
	return n.ClearingMember
}

func (n *NetPositionV03) AddClearingSegment() *model.PartyIdentification35Choice {
	n.ClearingSegment = new(model.PartyIdentification35Choice)
	return n.ClearingSegment
}

func (n *NetPositionV03) AddNetPositionReport() *model.NetPosition3 {
	newValue := new(model.NetPosition3)
	n.NetPositionReport = append(n.NetPositionReport, newValue)
	return newValue
}

func (n *NetPositionV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	n.SupplementaryData = append(n.SupplementaryData, newValue)
	return newValue
}
