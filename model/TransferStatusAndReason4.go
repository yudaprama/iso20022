package model

// Information about the status of a transfer instruction and its reason.
type TransferStatusAndReason4 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identification of a transfer, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Status of the transfer instruction.
	TransferStatus *TransferStatus2Choice `xml:"TrfSts"`

	// Date and time at which the transfer was executed.
	TradeDate *ISODate `xml:"TradDt,omitempty"`

	// Date on which the document, for example, the application form, was sent.
	SendOutDate *ISODate `xml:"SndOutDt,omitempty"`

	// Party that initiates the status.
	StatusInitiator *PartyIdentification70Choice `xml:"StsInitr,omitempty"`
}

func (t *TransferStatusAndReason4) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferStatusAndReason4) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *TransferStatusAndReason4) AddClientReference() *AdditionalReference7 {
	t.ClientReference = new(AdditionalReference7)
	return t.ClientReference
}

func (t *TransferStatusAndReason4) SetCancellationReference(value string) {
	t.CancellationReference = (*Max35Text)(&value)
}

func (t *TransferStatusAndReason4) AddTransferStatus() *TransferStatus2Choice {
	t.TransferStatus = new(TransferStatus2Choice)
	return t.TransferStatus
}

func (t *TransferStatusAndReason4) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *TransferStatusAndReason4) SetSendOutDate(value string) {
	t.SendOutDate = (*ISODate)(&value)
}

func (t *TransferStatusAndReason4) AddStatusInitiator() *PartyIdentification70Choice {
	t.StatusInitiator = new(PartyIdentification70Choice)
	return t.StatusInitiator
}
