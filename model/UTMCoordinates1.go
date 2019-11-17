package model

// Location on the Earth specified by the Universal Transverse Mercator coordinate system, using the WGS84 geodesic system.
type UTMCoordinates1 struct {

	// UTM grid zone combination of the longitude zone (1 to 60) and the latitude band, C to X, excluding I and O (for example Eiffel tower UTM zone is 31U).
	UTMZone *Max16Text `xml:"UTMZone"`

	// X-coordinate of the Universal Transverse Mercator coordinate system in meters (for example 448 265m for Eiffel Tower X-coordinate).
	UTMEastward *Number `xml:"UTMEstwrd"`

	// Y-coordinate of the Universal Transverse Mercator coordinate system (for example 5 411 920m for Eiffel Tower Y-coordinate).
	UTMNorthward *Number `xml:"UTMNrthwrd"`
}

func (u *UTMCoordinates1) SetUTMZone(value string) {
	u.UTMZone = (*Max16Text)(&value)
}

func (u *UTMCoordinates1) SetUTMEastward(value string) {
	u.UTMEastward = (*Number)(&value)
}

func (u *UTMCoordinates1) SetUTMNorthward(value string) {
	u.UTMNorthward = (*Number)(&value)
}
