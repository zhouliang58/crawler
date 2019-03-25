package poi

type PoiResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Total   int    `json:"total"`
	Results []Poi  `json:"results"`
}
