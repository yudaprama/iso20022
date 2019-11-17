package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document05500101 struct {
	XMLName xml.Name             `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.055.001.01 Document"`
	Message *PartyEventAdviceV01 `xml:"PtyEvtAdvc"`
}

func (d *Document05500101) AddMessage() *PartyEventAdviceV01 {
	d.Message = new(PartyEventAdviceV01)
	return d.Message
}

// The PartyEventAdvice message can be sent by any party to any other party. It is used for business letters containing information for which treatment is not formally defined in order to keep partners informed and to maintain business traces, for example confirmations of information exchanged out-of band such as announcing a postal letter, confirming a telephone call or the exchange of contractual information. It can also be sent to verify the technical interoperability of the communicating IT-systems.
// The message can reference other message and include data from referenced messages.
// The message can report several events.
// The message can carry digital signatures if required by context.
type PartyEventAdviceV01 struct {

	// Set of characteristics that unambiguously identify the event, common parameters, documents and identifications.
	Header *model.BusinessLetter1 `xml:"Hdr"`

	// Description of the event.
	EventNotice []*model.EventDescription1 `xml:"EvtNtce"`

	// Number of events as control value.
	EventCount *model.Max15NumericText `xml:"EvtCnt,omitempty"`

	// Referenced or related business message.
	AttachedMessage []*model.EncapsulatedBusinessMessage1 `xml:"AttchdMsg,omitempty"`
}

func (p *PartyEventAdviceV01) AddHeader() *model.BusinessLetter1 {
	p.Header = new(model.BusinessLetter1)
	return p.Header
}

func (p *PartyEventAdviceV01) AddEventNotice() *model.EventDescription1 {
	newValue := new(model.EventDescription1)
	p.EventNotice = append(p.EventNotice, newValue)
	return newValue
}

func (p *PartyEventAdviceV01) SetEventCount(value string) {
	p.EventCount = (*model.Max15NumericText)(&value)
}

func (p *PartyEventAdviceV01) AddAttachedMessage() *model.EncapsulatedBusinessMessage1 {
	newValue := new(model.EncapsulatedBusinessMessage1)
	p.AttachedMessage = append(p.AttachedMessage, newValue)
	return newValue
}
