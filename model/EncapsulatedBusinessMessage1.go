package model

// Defines an encapsulated form of an ISO 20022 message and, if present, its associated Business Application Header. The encapsulation guarantees uniqueness of ID/IDREFs though the use of the Prefix element. This element can be added during message preparation to ID/IDREFs. In order to verify the signature in the Hdr element or inside the encapsulated message, for each occurrence of an ID orIDREF that possesses the same value as a prefix, the prefix part is removed before signature verification. This is not done for surrounding signatures.
type EncapsulatedBusinessMessage1 struct {

	// The Business Application Header associated to the encapsulated message if it exists.
	Header *BusinessApplicationHeader1 `xml:"Hdr,omitempty"`

	// Prefix of ID/IDREFs in the encapsulated message to be removed before signature verification.
	Prefix *ID `xml:"Prfx,omitempty"`

	// If yes, the Msg element contains only a subset of the original message.
	Partial *YesNoIndicator `xml:"Prtl"`

	// The encapsulated ISO 20022 message.
	Message *StrictPayload `xml:"Msg"`
}

func (e *EncapsulatedBusinessMessage1) AddHeader() *BusinessApplicationHeader1 {
	e.Header = new(BusinessApplicationHeader1)
	return e.Header
}

func (e *EncapsulatedBusinessMessage1) SetPrefix(value string) {
	e.Prefix = (*ID)(&value)
}

func (e *EncapsulatedBusinessMessage1) SetPartial(value string) {
	e.Partial = (*YesNoIndicator)(&value)
}

func (e *EncapsulatedBusinessMessage1) AddMessage() *StrictPayload {
	e.Message = new(StrictPayload)
	return e.Message
}
