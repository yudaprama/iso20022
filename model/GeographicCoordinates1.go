package model

// Location on the Earth specified by two numbers representing vertical and horizontal position.
type GeographicCoordinates1 struct {

	// Latitude measured in degrees, minutes and seconds, following by 'N' for the north and 'S' for the south of the equator (for example 48°51'29" N for the Eiffel Tower latitude).
	Latitude *Max16Text `xml:"Lat"`

	// Angular measurement of the distance of a location on the earth east or west of the Greenwich observatory.
	// The longitude is measured in degrees, minutes and seconds, following by 'E' for the east and 'W' for the west (for example 2°17'40" E for the Eiffel Tower longitude).
	Longitude *Max16Text `xml:"Long"`
}

func (g *GeographicCoordinates1) SetLatitude(value string) {
	g.Latitude = (*Max16Text)(&value)
}

func (g *GeographicCoordinates1) SetLongitude(value string) {
	g.Longitude = (*Max16Text)(&value)
}
