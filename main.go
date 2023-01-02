package main
import (
	"fmt"
	"github.com/gocolly/colly"
)

type Car struct {
	CarName            string
	BodyCar            string
	ReleaseYear        string
	Color              string
	DriveUnit          string
	EngineVolume       string
	Mileage            string
	State              string
	FuelType           string
	CustomsClearedInRT string
	Transmission       string
	Price              string
	Url                string
}

func main() {
	scrapURL := "https://somon.tj/adv/8960821_tesla-model-y-2022/"
	c := colly.NewCollector(
		colly.AllowedDomains("www.somon.tj", "somon.tj"),
	)
	var carAttributes []string

	//Вызывается перед запросом
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting %s\n", r.URL)
	})

	//Вызывается, если во время запроса произошла ошибка
	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Error while scraping: %s\n", e.Error())
	})

	//	//Звонок после получения ответа
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited is success")
	})

	//Вызывается сразу после OnResponse, если полученный контент является HTML
	c.OnHTML(".chars-column li a", func(h *colly.HTMLElement) {
		carAttributes = append(carAttributes, h.Text)
	})

	c.Visit(scrapURL)

	car := Car{}
	car.BodyCar = carAttributes[0]
	car.ReleaseYear = carAttributes[1]
	car.Color = carAttributes[2]
	car.DriveUnit = carAttributes[3]
	car.EngineVolume = carAttributes[4]
	car.Mileage = carAttributes[5]
	car.State = carAttributes[6]
	car.FuelType = carAttributes[7]
	car.Transmission = carAttributes[8]

	fmt.Println("Наша тачка -", car)
}