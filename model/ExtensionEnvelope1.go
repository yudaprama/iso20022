package model

// Technical component that contains the validated extension information. This technical envelope allows to segregate the extension information from any other information.
type ExtensionEnvelope1 struct {

	// Technical element that specifies the extension.
	ExtensionContents *ExtensionContents1 `xml:"XtnsnCnts"`
}

func (e *ExtensionEnvelope1) AddExtensionContents() *ExtensionContents1 {
	e.ExtensionContents = new(ExtensionContents1)
	return e.ExtensionContents
}
