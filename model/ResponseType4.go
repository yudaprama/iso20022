package model

// Trace of response by the entities in the path from the issuer to the ATM.
type ResponseType4 struct {

	// Identification of the responder.
	ResponderIdentification *Max35Text `xml:"RspndrId"`

	// Codification of the response (for instance ISO 8583, IFX).
	Codification *Max35Text `xml:"Cdfctn,omitempty"`

	// Result of the request withdrawal message.
	Response *Max35Text `xml:"Rspn"`

	// Detail of the response.
	ResponseReason *Max35Text `xml:"RspnRsn,omitempty"`

	// Additional information to be logged for further examination.
	AdditionalResponseInformation *Max35Text `xml:"AddtlRspnInf,omitempty"`
}

func (r *ResponseType4) SetResponderIdentification(value string) {
	r.ResponderIdentification = (*Max35Text)(&value)
}

func (r *ResponseType4) SetCodification(value string) {
	r.Codification = (*Max35Text)(&value)
}

func (r *ResponseType4) SetResponse(value string) {
	r.Response = (*Max35Text)(&value)
}

func (r *ResponseType4) SetResponseReason(value string) {
	r.ResponseReason = (*Max35Text)(&value)
}

func (r *ResponseType4) SetAdditionalResponseInformation(value string) {
	r.AdditionalResponseInformation = (*Max35Text)(&value)
}
