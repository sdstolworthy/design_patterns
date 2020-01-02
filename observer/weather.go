package observer

type Observer interface {
	Update()
}

type UserInterface struct {
	WeatherStation WeatherStation
}

type Logger struct {
	WeatherStation WeatherStation
}

type Alert struct {
	WeatherStation WeatherStation
}

type WeatherStation struct {
	Temperature string
	WindSpeed   string
	Pressure    string
	observers   []*Observer
}

type Subject interface {
	registerObserver()
	removeObserver()
	notifyObservers()
}
