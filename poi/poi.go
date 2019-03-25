package poi

// POI 位置
type PoiLocation struct {
	Lat float64 `json:"lat"` //纬度
	Lgt float64 `json:"lng"` //经度
}

// POI 信息
type Poi struct {
	Name      string      `json:"name"`
	Location  PoiLocation `json:"location"`
	Address   string      `json:"address"`
	Province  string      `json:"province"`
	City      string      `json:"city"`
	Area      string      `json:"area"`
	Street_id string      `json:"street_id"`
	Telephone string      `json:"telephone"`
	//Detail    string      `json:"detail"`
	Uid string `json:"uid"`
}
