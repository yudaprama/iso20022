package reda

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300102 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.003.001.02 Document"`
	Message *PriceReportCorrectionV02 `xml:"reda.003.001.02"`
}

func (d *Document00300102) AddMessage() *PriceReportCorrectionV02 {
	d.Message = new(PriceReportCorrectionV02)
	return d.Message
}

// Scope
// The PriceReportCorrection message is sent by a report provider, eg, a fund accountant, transfer agent, market data provider, or any other interested party, to a report user, eg, a fund management company, a transfer agent, market data provider, regulator or any other interested party.
// The message is used to correct at least one of the prices, of a financial instrument, that was quoted in a previously sent PriceReport message.
// Usage
// The PriceReportCorrection message is used to correct information in a PriceReport message that was previously sent by the fund accountant. If an entire PriceReport message must be corrected, eg, due to an incorrect trade date, it is recommended that a PriceReportCancellation message is used to cancel the entire PriceReport message and a new PriceReport message is sent.
type PriceReportCorrectionV02 struct {

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference3 `xml:"PrvsRef"`

	// Information related to the correction of a price of a financial instrument.
	PriceCorrectionDetails []*model.PriceCorrection2 `xml:"PricCrrctnDtls"`
}

func (p *PriceReportCorrectionV02) AddPoolReference() *model.AdditionalReference3 {
	p.PoolReference = new(model.AdditionalReference3)
	return p.PoolReference
}

func (p *PriceReportCorrectionV02) AddPreviousReference() *model.AdditionalReference3 {
	p.PreviousReference = new(model.AdditionalReference3)
	return p.PreviousReference
}

func (p *PriceReportCorrectionV02) AddPriceCorrectionDetails() *model.PriceCorrection2 {
	newValue := new(model.PriceCorrection2)
	p.PriceCorrectionDetails = append(p.PriceCorrectionDetails, newValue)
	return newValue
}
