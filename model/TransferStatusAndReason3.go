package model

// Information about the status of a transfer instruction and its reason.
type TransferStatusAndReason3 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identification of a transfer, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Status of the transfer instruction.
	TransferStatus *TransferStatus1Choice `xml:"TrfSts"`

	// Date and time at which the transfer was executed.
	TradeDate *ISODate `xml:"TradDt,omitempty"`

	// Date on which the document, for example, the application form, was sent.
	SendOutDate *ISODate `xml:"SndOutDt,omitempty"`

	// Party that initiates the status.
	StatusInitiator *PartyIdentification2Choice `xml:"StsInitr,omitempty"`
}

func (t *TransferStatusAndReason3) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferStatusAndReason3) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *TransferStatusAndReason3) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *TransferStatusAndReason3) SetCancellationReference(value string) {
	t.CancellationReference = (*Max35Text)(&value)
}

func (t *TransferStatusAndReason3) AddTransferStatus() *TransferStatus1Choice {
	t.TransferStatus = new(TransferStatus1Choice)
	return t.TransferStatus
}

func (t *TransferStatusAndReason3) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *TransferStatusAndReason3) SetSendOutDate(value string) {
	t.SendOutDate = (*ISODate)(&value)
}

func (t *TransferStatusAndReason3) AddStatusInitiator() *PartyIdentification2Choice {
	t.StatusInitiator = new(PartyIdentification2Choice)
	return t.StatusInitiator
}
