package model

// Instructs the POI (Point Of Interaction) how to contact the host of the terminal management system (TMS), to initiate the maintenance of the terminal.
type TMSTrigger1 struct {

	// Level of urgency in contacting the maintenance.
	TMSContactLevel *TMSContactLevel1Code `xml:"TMSCtctLvl"`

	// Identification of the host to contact for the maintenance.
	TMSIdentification *Max35Text `xml:"TMSId,omitempty"`

	// Date and time for calling the maintenance.
	TMSContactDateTime *ISODateTime `xml:"TMSCtctDtTm,omitempty"`
}

func (t *TMSTrigger1) SetTMSContactLevel(value string) {
	t.TMSContactLevel = (*TMSContactLevel1Code)(&value)
}

func (t *TMSTrigger1) SetTMSIdentification(value string) {
	t.TMSIdentification = (*Max35Text)(&value)
}

func (t *TMSTrigger1) SetTMSContactDateTime(value string) {
	t.TMSContactDateTime = (*ISODateTime)(&value)
}
