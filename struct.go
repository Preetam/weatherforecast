package weatherforecast

type WeatherForecast struct {
	Code    string  `json:"cod"`
	Message float32 `json:"message"`
	City    struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Coordinates struct {
			Longitude float32 `json:"lon"`
			Latitiude float32 `json:"lat"`
		} `json:"coord"`
		Country    string `json:"country"`
		Population int    `json:"population"`
		Sys        struct {
			Population int `json:"population"`
		} `json:"sys"`
	} `json:"city"`
	Count int `json:"cnt"`
	List  []struct {
		Timestamp int `json:"dt"`
		Main      struct {
			Temp        float32 `json:"temp"`
			TempMin     float32 `json:"temp_min"`
			TempMax     float32 `json:"temp_max"`
			Pressure    float32 `json:"pressure"`
			SeaLevel    float32 `json:"sea_level"`
			GroundLevel float32 `json:"grnd_level"`
			Humidity    int     `json:"humidity"`
			TempKf      float32 `json:"temp_kf"`
		} `json:"main"`
		Weather []struct {
			Id          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Rain struct {
			ThreeHour int `json:"3h"`
		} `json:"rain"`
		Wind struct {
			Speed     float32 `json:"speed"`
			Direction float32 `json:"deg"`
		} `json:"wind"`
		Sys struct {
			Pod string `json:"pod"`
		} `json:"sys"`
		DtTxt string `json:"dt_txt"`
	} `json:"list"`
}
