package model

// Specifies the Business Application Header of the Business Message.
// Can be used when replying to a query;  can also be used when canceling or amending.
type BusinessApplicationHeader1 struct {

	// Contains the character set of the text-based elements used in the Business Message.
	CharacterSet *UnicodeChartsCode `xml:"CharSet,omitempty"`

	// The sending MessagingEndpoint that has created this Business Message for the receiving MessagingEndpoint that will process this Business Message.
	//
	// Note	the sending MessagingEndpoint might be different from the sending address potentially contained in the transport header (as defined in the transport layer).
	From *Party9Choice `xml:"Fr"`

	// The MessagingEndpoint designated by the sending MessagingEndpoint to be the recipient who will ultimately process this Business Message.
	//
	// Note the receiving MessagingEndpoint might be different from the receiving address potentially contained in the transport header (as defined in the transport layer).
	To *Party9Choice `xml:"To"`

	// Unambiguously identifies the Business Message to the MessagingEndpoint that has created the Business Message.
	BusinessMessageIdentifier *Max35Text `xml:"BizMsgIdr"`

	// Contains the MessageIdentifier that defines the BusinessMessage.
	// It must contain a MessageIdentifier published on the ISO 20022 website.
	//
	// example	camt.001.001.03
	MessageDefinitionIdentifier *Max35Text `xml:"MsgDefIdr"`

	// Specifies the business service agreed between the two MessagingEndpoints under which rules this Business Message is exchanged.
	//  To be used when there is a choice of processing services or processing service levels.
	// Example: E&I
	BusinessService *Max35Text `xml:"BizSvc,omitempty"`

	// Date and time when this Business Message (header) was created.
	// Note    Times must be normalized, using the "Z" annotation.
	CreationDate *ISONormalisedDateTime `xml:"CreDt"`

	// Indicates whether the message is a Copy, a Duplicate or a copy of a duplicate of a previously sent ISO 20022 Message.
	CopyDuplicate *CopyDuplicate1Code `xml:"CpyDplct,omitempty"`

	// Flag indicating if the Business Message exchanged between the MessagingEndpoints is possibly a duplicate.
	// If the receiving MessagingEndpoint  did not receive the original, then this Business Message should be processed as if it were the original.
	//
	// If the receiving MessagingEndpoint did receive the original, then it should perform necessary actions to avoid processing this Business Message again.
	//
	// This will guarantee business idempotent behaviour.
	//
	// NOTE: this is named "PossResend" in FIX - this is an application level resend not a network level retransmission
	PossibleDuplicate *YesNoIndicator `xml:"PssblDplct,omitempty"`

	// Relative indication of the processing precedence of the message over a (set of) Business Messages with assigned priorities.
	Priority *BusinessMessagePriorityCode `xml:"Prty,omitempty"`

	// Contains the digital signature of the Business Entity authorised to sign this Business Message.
	Signature *SignatureEnvelope `xml:"Sgntr,omitempty"`
}

func (b *BusinessApplicationHeader1) SetCharacterSet(value string) {
	b.CharacterSet = (*UnicodeChartsCode)(&value)
}

func (b *BusinessApplicationHeader1) AddFrom() *Party9Choice {
	b.From = new(Party9Choice)
	return b.From
}

func (b *BusinessApplicationHeader1) AddTo() *Party9Choice {
	b.To = new(Party9Choice)
	return b.To
}

func (b *BusinessApplicationHeader1) SetBusinessMessageIdentifier(value string) {
	b.BusinessMessageIdentifier = (*Max35Text)(&value)
}

func (b *BusinessApplicationHeader1) SetMessageDefinitionIdentifier(value string) {
	b.MessageDefinitionIdentifier = (*Max35Text)(&value)
}

func (b *BusinessApplicationHeader1) SetBusinessService(value string) {
	b.BusinessService = (*Max35Text)(&value)
}

func (b *BusinessApplicationHeader1) SetCreationDate(value string) {
	b.CreationDate = (*ISONormalisedDateTime)(&value)
}

func (b *BusinessApplicationHeader1) SetCopyDuplicate(value string) {
	b.CopyDuplicate = (*CopyDuplicate1Code)(&value)
}

func (b *BusinessApplicationHeader1) SetPossibleDuplicate(value string) {
	b.PossibleDuplicate = (*YesNoIndicator)(&value)
}

func (b *BusinessApplicationHeader1) SetPriority(value string) {
	b.Priority = (*BusinessMessagePriorityCode)(&value)
}

func (b *BusinessApplicationHeader1) AddSignature() *SignatureEnvelope {
	b.Signature = new(SignatureEnvelope)
	return b.Signature
}
