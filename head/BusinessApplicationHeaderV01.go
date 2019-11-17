package head

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100101 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 Document"`
	Message *BusinessApplicationHeaderV01 `xml:"AppHdr"`
}

func (d *Document00100101) AddMessage() *BusinessApplicationHeaderV01 {
	d.Message = new(BusinessApplicationHeaderV01)
	return d.Message
}

// The Business Layer deals with Business Messages. The behaviour of the Business Messages is fully described by the Business Transaction and the structure of the Business Messages is fully described by the Message Definitions and related Message Rules, Rules and Market Practices. All of which are registered in the ISO 20022 Repository.
// A single new Business Message (with its accompagnying business application header) is created - by the sending MessagingEndpoint - for each business event; that is each interaction in a Business Transaction. A Business Message adheres to the following principles:
// " A Business Message (and its business application header) must not contain information about the Message Transport System or the mechanics or mechanism of message sending, transportation, or receipt.
// " A Business Message must be comprehensible outside of the context of the Transport Message. That is the Business Message must not require knowledge of the Transport Message to be understood.
// " A Business Message may contain headers, footers, and envelopes that are meaningful for the business. When present, they are treated as any other message content, which means that they are considered part of the Message Definition of the Business Message and as such will be part of the ISO 20022 Repository.
// " A Business Message refers to Business Actors by their Name. Each instance of a Business Actor has one Name. The Business Actor must not be referred to in the Transport Layer.
// Specific usage of this BusinessMessageHeader may be defined by the relevant SEG.
type BusinessApplicationHeaderV01 struct {

	// Contains the character set of the text-based elements used in the Business Message.
	CharacterSet *model.UnicodeChartsCode `xml:"CharSet,omitempty"`

	// The sending MessagingEndpoint that has created this Business Message for the receiving MessagingEndpoint that will process this Business Message.
	//
	// Note	the sending MessagingEndpoint might be different from the sending address potentially contained in the transport header (as defined in the transport layer).
	From *model.Party9Choice `xml:"Fr"`

	// The MessagingEndpoint designated by the sending MessagingEndpoint to be the recipient who will ultimately process this Business Message.
	//
	// Note the receiving MessagingEndpoint might be different from the receiving address potentially contained in the transport header (as defined in the transport layer).
	To *model.Party9Choice `xml:"To"`

	// Unambiguously identifies the Business Message to the MessagingEndpoint that has created the Business Message.
	BusinessMessageIdentifier *model.Max35Text `xml:"BizMsgIdr"`

	// Contains the MessageIdentifier that defines the BusinessMessage.
	// It must contain a MessageIdentifier published on the ISO 20022 website.
	//
	// example	camt.001.001.03
	MessageDefinitionIdentifier *model.Max35Text `xml:"MsgDefIdr"`

	// Specifies the business service agreed between the two MessagingEndpoints under which rules this Business Message is exchanged.
	//  To be used when there is a choice of processing services or processing service levels.
	// Example: E&I
	BusinessService *model.Max35Text `xml:"BizSvc,omitempty"`

	// Date and time when this Business Message (header) was created.
	// Note    Times must be normalized, using the "Z" annotation.
	CreationDate *model.ISONormalisedDateTime `xml:"CreDt"`

	// Indicates whether the message is a Copy, a Duplicate or a copy of a duplicate of a previously sent ISO 20022 Message.
	CopyDuplicate *model.CopyDuplicate1Code `xml:"CpyDplct,omitempty"`

	// Flag indicating if the Business Message exchanged between the MessagingEndpoints is possibly a duplicate.
	// If the receiving MessagingEndpoint  did not receive the original, then this Business Message should be processed as if it were the original.
	//
	// If the receiving MessagingEndpoint did receive the original, then it should perform necessary actions to avoid processing this Business Message again.
	//
	// This will guarantee business idempotent behaviour.
	//
	// NOTE: this is named "PossResend" in FIX - this is an application level resend not a network level retransmission
	PossibleDuplicate *model.YesNoIndicator `xml:"PssblDplct,omitempty"`

	// Relative indication of the processing precedence of the message over a (set of) Business Messages with assigned priorities.
	Priority *model.BusinessMessagePriorityCode `xml:"Prty,omitempty"`

	// Contains the digital signature of the Business Entity authorised to sign this Business Message.
	Signature *model.SignatureEnvelope `xml:"Sgntr,omitempty"`

	// Specifies the Business Application Header of the Business Message to which this Business Message relates.
	// Can be used when replying to a query;  can also be used when canceling or amending.
	Related *model.BusinessApplicationHeader1 `xml:"Rltd,omitempty"`
}

func (b *BusinessApplicationHeaderV01) SetCharacterSet(value string) {
	b.CharacterSet = (*model.UnicodeChartsCode)(&value)
}

func (b *BusinessApplicationHeaderV01) AddFrom() *model.Party9Choice {
	b.From = new(model.Party9Choice)
	return b.From
}

func (b *BusinessApplicationHeaderV01) AddTo() *model.Party9Choice {
	b.To = new(model.Party9Choice)
	return b.To
}

func (b *BusinessApplicationHeaderV01) SetBusinessMessageIdentifier(value string) {
	b.BusinessMessageIdentifier = (*model.Max35Text)(&value)
}

func (b *BusinessApplicationHeaderV01) SetMessageDefinitionIdentifier(value string) {
	b.MessageDefinitionIdentifier = (*model.Max35Text)(&value)
}

func (b *BusinessApplicationHeaderV01) SetBusinessService(value string) {
	b.BusinessService = (*model.Max35Text)(&value)
}

func (b *BusinessApplicationHeaderV01) SetCreationDate(value string) {
	b.CreationDate = (*model.ISONormalisedDateTime)(&value)
}

func (b *BusinessApplicationHeaderV01) SetCopyDuplicate(value string) {
	b.CopyDuplicate = (*model.CopyDuplicate1Code)(&value)
}

func (b *BusinessApplicationHeaderV01) SetPossibleDuplicate(value string) {
	b.PossibleDuplicate = (*model.YesNoIndicator)(&value)
}

func (b *BusinessApplicationHeaderV01) SetPriority(value string) {
	b.Priority = (*model.BusinessMessagePriorityCode)(&value)
}

func (b *BusinessApplicationHeaderV01) AddSignature() *model.SignatureEnvelope {
	b.Signature = new(model.SignatureEnvelope)
	return b.Signature
}

func (b *BusinessApplicationHeaderV01) AddRelated() *model.BusinessApplicationHeader1 {
	b.Related = new(model.BusinessApplicationHeader1)
	return b.Related
}
