package domain

type Location struct {
	Name         string
	LocationType int
	Latitude     float64
	Longitude    float64
	Content      string
}

type MyLocation struct {
	Latitude  float64
	Longitude float64
}
type LocationDto struct {
	ID       string
	Location Location
}

type LocationUrlParameter struct {
	ID string `uri:"locationId"`
}

type LocationTypeUrlParameter struct {
	LocationType int `uri:"locationType"`
}
