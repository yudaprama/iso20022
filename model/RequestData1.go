package model

// Contains the meta data for a netting cut off update request: message identification, request servicer and a request type.
type RequestData1 struct {

	// Unique identification of the message
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Description of the type of request.
	RequestType *Max4Text `xml:"ReqTp"`

	// Specifies the business date on which the new netting cut off(s) is (are) to be activated.
	RequestedActivationDate *ISODate `xml:"ReqdActvtnDt"`

	// Describes the central system servicing the request.
	RequestServicer *PartyIdentification73Choice `xml:"ReqSvcr,omitempty"`

	// Describes the participant issuing the request.
	NetServiceParticipantIdentification *PartyIdentification73Choice `xml:"NetSvcPtcptId"`

	// Describes the type of netting service supporting the net report.
	NetServiceType *Max35Text `xml:"NetSvcTp,omitempty"`
}

func (r *RequestData1) SetMessageIdentification(value string) {
	r.MessageIdentification = (*Max35Text)(&value)
}

func (r *RequestData1) SetRequestType(value string) {
	r.RequestType = (*Max4Text)(&value)
}

func (r *RequestData1) SetRequestedActivationDate(value string) {
	r.RequestedActivationDate = (*ISODate)(&value)
}

func (r *RequestData1) AddRequestServicer() *PartyIdentification73Choice {
	r.RequestServicer = new(PartyIdentification73Choice)
	return r.RequestServicer
}

func (r *RequestData1) AddNetServiceParticipantIdentification() *PartyIdentification73Choice {
	r.NetServiceParticipantIdentification = new(PartyIdentification73Choice)
	return r.NetServiceParticipantIdentification
}

func (r *RequestData1) SetNetServiceType(value string) {
	r.NetServiceType = (*Max35Text)(&value)
}
