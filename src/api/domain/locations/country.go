package locations

type Country struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	TimeZone       string         `json:"timezone"`
	GeoInformation GeoInformation `json:"geo_information"`
	States         []States       `json:"states"`
}

type GeoInformation struct {
	Location Location `json:"location"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type States struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
