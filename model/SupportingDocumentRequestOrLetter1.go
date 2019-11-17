package model

// Currency control related document or letter supporting the contract registration.
type SupportingDocumentRequestOrLetter1 struct {

	// Unique and unambiguous identification of the supporting document request or the letter.
	RequestOrLetterIdentification *Max35Text `xml:"ReqOrLttrId"`

	// Date of the supporting document request or the letter.
	Date *ISODate `xml:"Dt,omitempty"`

	// Sender of the request or letter.
	Sender *Party28Choice `xml:"Sndr,omitempty"`

	// Receiver of the request or letter.
	Receiver *Party28Choice `xml:"Rcvr,omitempty"`

	// Provides the references of the original underlying message(s) for which supporting documents are requested or for which the letter is sent.
	OriginalReferences []*OriginalMessage2 `xml:"OrgnlRefs,omitempty"`

	// Subject of the letter or supporting document.
	Subject *Max140Text `xml:"Sbjt"`

	// Provides the type of supporting document requested.
	Type *SupportDocumentType1Code `xml:"Tp"`

	// Further free format description of the request.
	Description *Max1025Text `xml:"Desc,omitempty"`

	// Flag to indicate whether a response is required or not.
	//
	// Usage: when the request is used to send a letter, there is no response required.
	ResponseRequired *TrueFalseIndicator `xml:"RspnReqrd"`

	// Date by which the response to the request is expected.
	DueDate *ISODate `xml:"DueDt,omitempty"`

	// Documents provided as attachments to the supporting document request or letter.
	Attachment []*DocumentGeneralInformation3 `xml:"Attchmnt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SupportingDocumentRequestOrLetter1) SetRequestOrLetterIdentification(value string) {
	s.RequestOrLetterIdentification = (*Max35Text)(&value)
}

func (s *SupportingDocumentRequestOrLetter1) SetDate(value string) {
	s.Date = (*ISODate)(&value)
}

func (s *SupportingDocumentRequestOrLetter1) AddSender() *Party28Choice {
	s.Sender = new(Party28Choice)
	return s.Sender
}

func (s *SupportingDocumentRequestOrLetter1) AddReceiver() *Party28Choice {
	s.Receiver = new(Party28Choice)
	return s.Receiver
}

func (s *SupportingDocumentRequestOrLetter1) AddOriginalReferences() *OriginalMessage2 {
	newValue := new(OriginalMessage2)
	s.OriginalReferences = append(s.OriginalReferences, newValue)
	return newValue
}

func (s *SupportingDocumentRequestOrLetter1) SetSubject(value string) {
	s.Subject = (*Max140Text)(&value)
}

func (s *SupportingDocumentRequestOrLetter1) SetType(value string) {
	s.Type = (*SupportDocumentType1Code)(&value)
}

func (s *SupportingDocumentRequestOrLetter1) SetDescription(value string) {
	s.Description = (*Max1025Text)(&value)
}

func (s *SupportingDocumentRequestOrLetter1) SetResponseRequired(value string) {
	s.ResponseRequired = (*TrueFalseIndicator)(&value)
}

func (s *SupportingDocumentRequestOrLetter1) SetDueDate(value string) {
	s.DueDate = (*ISODate)(&value)
}

func (s *SupportingDocumentRequestOrLetter1) AddAttachment() *DocumentGeneralInformation3 {
	newValue := new(DocumentGeneralInformation3)
	s.Attachment = append(s.Attachment, newValue)
	return newValue
}

func (s *SupportingDocumentRequestOrLetter1) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
