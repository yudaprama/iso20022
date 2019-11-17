package model

// Communication details related to the registered currency control contract.
type RegisteredContractCommunication1 struct {

	// Method by which the registered contract document is exchanged.
	Method *CommunicationMethod4Code `xml:"Mtd"`

	// Date of the exchange.
	Date *ISODate `xml:"Dt"`
}

func (r *RegisteredContractCommunication1) SetMethod(value string) {
	r.Method = (*CommunicationMethod4Code)(&value)
}

func (r *RegisteredContractCommunication1) SetDate(value string) {
	r.Date = (*ISODate)(&value)
}
