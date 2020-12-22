package climacell

import "time"

type Geometry struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates"`
}

type TimelineListOptions struct {
	Location  Geometry
	Fields    []string
	StartTime string
	EndTime   string
	TimeSteps []string
}

type TimelineList struct {
	Data []Timeline `json:"data"`
}

type Timeline struct {
	Timestep  string     `json:"timestep"`
	StartTime time.Time  `json:"startTime"`
	EndTime   time.Time  `json:"endTime"`
	Intervals []Interval `json:"intervals"`
}

type Interval struct {
	StartTime time.Time `json:"startTime"`
	Values    Values    `json:"values"`
}

type Values struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type Core struct {
}

type Temperature struct {
	TemperatureMax     float64   `json:"temperatureMax"`
	TemperatureMin     float64   `json:"temperatureMin"`
	TemperatureAvg     float64   `json:"temperatureAvg"`
	TemperatureMaxTime time.Time `json:"temperatureMaxTime"`
	TemperatureMinTime time.Time `json:"temperatureMinTime"`
}

type TemperatureApparent struct {
	TemperatureApparentMax     float64   `json:"temperatureApparentMax"`
	TemperatureApparentMin     float64   `json:"temperatureApparentMin"`
	TemperatureApparentAvg     float64   `json:"temperatureApparentAvg"`
	TemperatureApparentMaxTime time.Time `json:"temperatureApparentMaxTime"`
	TemperatureApparentMinTime time.Time `json:"temperatureApparentMinTime"`
}

type DewPoint struct {
	DewPointMax     float64   `json:"dewPointMax"`
	DewPointMin     float64   `json:"dewPointMin"`
	DewPointAvg     float64   `json:"dewPointAvg"`
	DewPointMaxTime time.Time `json:"dewPointMaxTime"`
	DewPointMinTime time.Time `json:"dewPointMinTime"`
}
type Humidity struct {
	HumidityMax     float64   `json:"humidityMax"`
	HumidityMin     float64   `json:"humidityMin"`
	HumidityAvg     float64   `json:"humidityAvg"`
	HumidityMaxTime time.Time `json:"humidityMaxTime"`
	HumidityMinTime time.Time `json:"humidityMinTime"`
}
type WindSpeed struct {
	WindSpeedMax     float64   `json:"windSpeedMax"`
	WindSpeedMin     float64   `json:"windSpeedMin"`
	WindSpeedAvg     float64   `json:"windSpeedAvg"`
	WindSpeedMaxTime time.Time `json:"windSpeedMaxTime"`
	WindSpeedMinTime time.Time `json:"windSpeedMinTime"`
}

type WindDirection struct {
	WindDirectionMax     float64   `json:"windDirectionMax"`
	WindDirectionMin     float64   `json:"windDirectionMin"`
	WindDirectionAvg     float64   `json:"windDirectionAvg"`
	WindDirectionMaxTime time.Time `json:"windDirectionMaxTime"`
	WindDirectionMinTime time.Time `json:"windDirectionMinTime"`
}

type WindGust struct {
	WindGustMax     float64   `json:"windGustMax"`
	WindGustMin     float64   `json:"windGustMin"`
	WindGustAvg     float64   `json:"windGustAvg"`
	WindGustMaxTime time.Time `json:"windGustMaxTime"`
	WindGustMinTime time.Time `json:"windGustMinTime"`
}

type PressureSurfaceLevel struct {
	PressureSurfaceLevelMax     float64   `json:"pressureSurfaceLevelMax"`
	PressureSurfaceLevelMin     float64   `json:"pressureSurfaceLevelMin"`
	PressureSurfaceLevelAvg     float64   `json:"pressureSurfaceLevelAvg"`
	PressureSurfaceLevelMaxTime time.Time `json:"pressureSurfaceLevelMaxTime"`
	PressureSurfaceLevelMinTime time.Time `json:"pressureSurfaceLevelMinTime"`
}

type PressureSeaLevel struct {
	PressureSeaLevelMax     float64   `json:"pressureSeaLevelMax"`
	PressureSeaLevelMin     float64   `json:"pressureSeaLevelMin"`
	PressureSeaLevelAvg     float64   `json:"pressureSeaLevelAvg"`
	PressureSeaLevelMaxTime time.Time `json:"pressureSeaLevelMaxTime"`
	PressureSeaLevelMinTime time.Time `json:"pressureSeaLevelMinTime"`
}

type PrecipitationIntensity struct {
	PrecipitationIntensityMax     float64   `json:"precipitationIntensityMax"`
	PrecipitationIntensityMin     float64   `json:"precipitationIntensityMin"`
	PrecipitationIntensityAvg     float64   `json:"precipitationIntensityAvg"`
	PrecipitationIntensityMaxTime time.Time `json:"precipitationIntensityMaxTime"`
	PrecipitationIntensityMinTime time.Time `json:"precipitationIntensityMinTime"`
}

type PrecipitationProbability struct {
	PrecipitationProbabilityMax     float64   `json:"precipitationProbabilityMax"`
	PrecipitationProbabilityMin     float64   `json:"precipitationProbabilityMin"`
	PrecipitationProbabilityAvg     float64   `json:"precipitationProbabilityAvg"`
	PrecipitationProbabilityMaxTime time.Time `json:"precipitationProbabilityMaxTime"`
	PrecipitationProbabilityMinTime time.Time `json:"precipitationProbabilityMinTime"`
}

type PrecipitationType struct {
	PrecipitationTypeMax     float64   `json:"precipitationTypeMax"`
	PrecipitationTypeMin     float64   `json:"precipitationTypeMin"`
	PrecipitationTypeAvg     float64   `json:"precipitationTypeAvg"`
	PrecipitationTypeMaxTime time.Time `json:"precipitationTypeMaxTime"`
	PrecipitationTypeMinTime time.Time `json:"precipitationTypeMinTime"`
}

type SunriseTime struct {
	sunriseTime time.Time `json:"sunriseTime"`
}

type SunsetTime struct {
	sunsetTime time.Time `json:"sunsetTime"`
}

type SolarGHI struct {
	SolarGHIMax     float64   `json:"solarGHIMax"`
	SolarGHIMin     float64   `json:"solarGHIMin"`
	SolarGHIAvg     float64   `json:"solarGHIAvg"`
	SolarGHIMaxTime time.Time `json:"solarGHIMaxTime"`
	SolarGHIMinTime time.Time `json:"solarGHIMinTime"`
}

type Visibility struct {
	VisibilityMax     float64   `json:"visibilityMax"`
	VisibilityMin     float64   `json:"visibilityMin"`
	VisibilityAvg     float64   `json:"visibilityAvg"`
	VisibilityMaxTime time.Time `json:"visibilityMaxTime"`
	VisibilityMinTime time.Time `json:"visibilityMinTime"`
}

type CloudCover struct {
	CloudCoverMax     float64   `json:"cloudCoverMax"`
	CloudCoverMin     float64   `json:"cloudCoverMin"`
	CloudCoverAvg     float64   `json:"cloudCoverAvg"`
	CloudCoverMaxTime time.Time `json:"cloudCoverMaxTime"`
	CloudCoverMinTime time.Time `json:"cloudCoverMinTime"`
}

type CloudBase struct {
	CloudBaseMax     float64   `json:"cloudBaseMax"`
	CloudBaseMin     float64   `json:"cloudBaseMin"`
	CloudBaseAvg     float64   `json:"cloudBaseAvg"`
	CloudBaseMaxTime time.Time `json:"cloudBaseMaxTime"`
	CloudBaseMinTime time.Time `json:"cloudBaseMinTime"`
}

type CloudCeiling struct {
	CloudCeilingMax     float64   `json:"cloudCeilingMax"`
	CloudCeilingMin     float64   `json:"cloudCeilingMin"`
	CloudCeilingAvg     float64   `json:"cloudCeilingAvg"`
	CloudCeilingMaxTime time.Time `json:"cloudCeilingMaxTime"`
	CloudCeilingMinTime time.Time `json:"cloudCeilingMinTime"`
}

type MoonPhase struct {
	MoonPhase int `json:"moonPhase"`
}

type WeatherCode struct {
	WeatherCode int `json:"weatherCode"`
}
