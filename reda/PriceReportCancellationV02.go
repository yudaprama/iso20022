package reda

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200102 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.002.001.02 Document"`
	Message *PriceReportCancellationV02 `xml:"reda.002.001.02"`
}

func (d *Document00200102) AddMessage() *PriceReportCancellationV02 {
	d.Message = new(PriceReportCancellationV02)
	return d.Message
}

// Scope
// The PriceReportCancellation message is sent by a report provider, eg, a fund accountant, transfer agent, market data provider, or any other interested party, to a report user, eg, a fund management company, a transfer agent, market data provider, regulator or any other interested party.
// The message is used to cancel a previously sent PriceReport message.
// Usage
// The PriceReportCancellation message is used to cancel an entire PriceReport message that was previously sent by the report provider. If only a part of the information needs to be cancelled and replaced, the PriceReportCorrection message must be used.
// This message must contain the reference of the message to be cancelled. This message may also contain details of the message to be cancelled, but this is not recommended.
type PriceReportCancellationV02 struct {

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference3 `xml:"PrvsRef"`

	// Common information related to all the price reports to be cancelled.
	PriceReportToBeCancelled *model.PriceReport1 `xml:"PricRptToBeCanc,omitempty"`
}

func (p *PriceReportCancellationV02) AddPoolReference() *model.AdditionalReference3 {
	p.PoolReference = new(model.AdditionalReference3)
	return p.PoolReference
}

func (p *PriceReportCancellationV02) AddPreviousReference() *model.AdditionalReference3 {
	p.PreviousReference = new(model.AdditionalReference3)
	return p.PreviousReference
}

func (p *PriceReportCancellationV02) AddPriceReportToBeCancelled() *model.PriceReport1 {
	p.PriceReportToBeCancelled = new(model.PriceReport1)
	return p.PriceReportToBeCancelled
}
