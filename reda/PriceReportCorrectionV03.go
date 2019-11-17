package reda

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300103 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.003.001.03 Document"`
	Message *PriceReportCorrectionV03 `xml:"PricRptCrrctnV03"`
}

func (d *Document00300103) AddMessage() *PriceReportCorrectionV03 {
	d.Message = new(PriceReportCorrectionV03)
	return d.Message
}

// Scope
// A report provider, eg, a transfer agent, fund accountant or market data provider, sends the PriceReportCorrection message to the report recipient, eg, a fund management company, transfer agent, market data provider, regulator or any other interested party to correct at least one of the prices for a financial instrument that was reported in a previously sent PriceReport message.
// Usage
// The PriceReportCorrection message is used to correct information reported in a previously sent PriceReport.
// If an entire PriceReport message must be corrected, eg, due to an incorrect trade date, it is recommended that a PriceReportCancellation message is used to cancel the entire PriceReport message and a new PriceReport sent.
type PriceReportCorrectionV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Pagination of the message.
	MessagePagination *model.Pagination `xml:"MsgPgntn"`

	// Information related to the correction of a price of a financial instrument.
	PriceCorrectionDetails []*model.PriceCorrection3 `xml:"PricCrrctnDtls"`
}

func (p *PriceReportCorrectionV03) AddMessageIdentification() *model.MessageIdentification1 {
	p.MessageIdentification = new(model.MessageIdentification1)
	return p.MessageIdentification
}

func (p *PriceReportCorrectionV03) AddPoolReference() *model.AdditionalReference3 {
	p.PoolReference = new(model.AdditionalReference3)
	return p.PoolReference
}

func (p *PriceReportCorrectionV03) AddPreviousReference() *model.AdditionalReference3 {
	p.PreviousReference = new(model.AdditionalReference3)
	return p.PreviousReference
}

func (p *PriceReportCorrectionV03) AddMessagePagination() *model.Pagination {
	p.MessagePagination = new(model.Pagination)
	return p.MessagePagination
}

func (p *PriceReportCorrectionV03) AddPriceCorrectionDetails() *model.PriceCorrection3 {
	newValue := new(model.PriceCorrection3)
	p.PriceCorrectionDetails = append(p.PriceCorrectionDetails, newValue)
	return newValue
}
