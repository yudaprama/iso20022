package model

// Information related to the reject of a message from an ATM or an ATM manager.
type ATMReject1 struct {

	// Identification of the entity sending the reject message.
	RejectInitiatorIdentification *Max35Text `xml:"RjctInitrId,omitempty"`

	// High level information allowing the sender of a request or an advice to know the types of error, and handle them accordingly.
	RejectReason *RejectReason1Code `xml:"RjctRsn"`

	// Additional information related to the sending of a reject message in response to a request or an advice.
	// For logging purpose, in order to allow further analysis, statistics and deferred processing on the success or the failure of the request processing.
	AdditionalInformation *Max500Text `xml:"AddtlInf,omitempty"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand1 `xml:"Cmd,omitempty"`

	// Received message that has been rejected.
	MessageInError *Max100KBinary `xml:"MsgInErr,omitempty"`
}

func (a *ATMReject1) SetRejectInitiatorIdentification(value string) {
	a.RejectInitiatorIdentification = (*Max35Text)(&value)
}

func (a *ATMReject1) SetRejectReason(value string) {
	a.RejectReason = (*RejectReason1Code)(&value)
}

func (a *ATMReject1) SetAdditionalInformation(value string) {
	a.AdditionalInformation = (*Max500Text)(&value)
}

func (a *ATMReject1) AddCommand() *ATMCommand1 {
	newValue := new(ATMCommand1)
	a.Command = append(a.Command, newValue)
	return newValue
}

func (a *ATMReject1) SetMessageInError(value string) {
	a.MessageInError = (*Max100KBinary)(&value)
}
