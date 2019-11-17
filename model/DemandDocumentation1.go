package model

// Document presented for examination.
type DemandDocumentation1 struct {

	// Indication as to whether the presentation is complete.
	CompleteIndicator *YesNoIndicator `xml:"CmpltInd"`

	// Information related to an incomplete presentation.
	CompletionInformation *Max2000Text `xml:"CmpltnInf,omitempty"`

	// Document or template enclosed in the demand.
	EnclosedFile []*Document9 `xml:"NclsdFile,omitempty"`

	// Narrative text constituting the demand.
	DemandNarrative *Max20000Text `xml:"DmndNrrtv,omitempty"`
}

func (d *DemandDocumentation1) SetCompleteIndicator(value string) {
	d.CompleteIndicator = (*YesNoIndicator)(&value)
}

func (d *DemandDocumentation1) SetCompletionInformation(value string) {
	d.CompletionInformation = (*Max2000Text)(&value)
}

func (d *DemandDocumentation1) AddEnclosedFile() *Document9 {
	newValue := new(Document9)
	d.EnclosedFile = append(d.EnclosedFile, newValue)
	return newValue
}

func (d *DemandDocumentation1) SetDemandNarrative(value string) {
	d.DemandNarrative = (*Max20000Text)(&value)
}
