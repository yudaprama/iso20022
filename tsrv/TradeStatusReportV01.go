package tsrv

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01800101 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.018.001.01 Document"`
	Message *TradeStatusReportV01 `xml:"TradStsRpt"`
}

func (d *Document01800101) AddMessage() *TradeStatusReportV01 {
	d.Message = new(TradeStatusReportV01)
	return d.Message
}

// The TradeStatusReport message is exchanged between parties involved in the trade finance domain to report the transaction level status of a transaction previously received. It informs the sender about the positive or negative status of the referenced transaction, such as acceptance or rejection resulting from technical validation performed by the parser and/or front-office applications. It can be used, for example, to acknowledge receipt of a transaction, to report a syntactical error, to report an unrecognised digital signature, to indicate that further processing is pending, and to indicate that a transaction has been technically accepted for processing by the back-office application. Multiple TradeStatusReport messages may be progressively sent in response to the incremental processing of a single transaction.
type TradeStatusReportV01 struct {

	// Details of the trade status report.
	TradeStatusAdviceDetails *model.TradeStatusReport1 `xml:"TradStsAdvcDtls"`

	// Digital signature of the report.
	DigitalSignature *model.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (t *TradeStatusReportV01) AddTradeStatusAdviceDetails() *model.TradeStatusReport1 {
	t.TradeStatusAdviceDetails = new(model.TradeStatusReport1)
	return t.TradeStatusAdviceDetails
}

func (t *TradeStatusReportV01) AddDigitalSignature() *model.PartyAndSignature2 {
	t.DigitalSignature = new(model.PartyAndSignature2)
	return t.DigitalSignature
}
