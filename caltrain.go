package caltrain

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Station string
type Direction int

const (
	// BaseURL is the common url between station realtime web pages.
	BaseURL = "http://www.caltrain.com/schedules/realtime/stations/%sstation-mobile.html"

	// direction of trains
	SouthBound Direction = iota
	NorthBound Direction = iota

	// list of train strations as of June 1, 2016
	SanFrancisco          Station = "sanfrancisco"
	TwentySecondStreet    Station = "22ndstreet"
	Bayshore              Station = "bayshore"
	SouthSanFrancisco     Station = "southsanfrancisco"
	SanBruno              Station = "sanbruno"
	MillbraeTransitCenter Station = "millbraetransitcenter"
	Broadway              Station = "broadway"
	Burlingame            Station = "burlingame"
	SanMateo              Station = "sanmateo"
	HaywardPark           Station = "haywardpark"
	Hillsdale             Station = "hillsdale"
	Belmont               Station = "belmont"
	SanCarlos             Station = "sancarlos"
	RedwoodCity           Station = "redwoodcity"
	Atherton              Station = "atherton"
	MenloPark             Station = "menlopark"
	PaloAlto              Station = "paloalto"
	CaliforniaAve         Station = "californiaave"
	SanAntonio            Station = "sanantonio"
	MountainView          Station = "mountainview"
	Sunnyvale             Station = "sunnyvale"
	Lawrence              Station = "lawrence"
	SantaClara            Station = "santaclara"
	CollegePark           Station = "collegepark"
	SanJoseDiridon        Station = "sanjosediridon"
	Tamien                Station = "tamien"
	Capitol               Station = "capitol"
	BlossomHill           Station = "blossomhill"
	MorganHill            Station = "morganhill"
	SanMartin             Station = "sanmartin"
	Gilroy                Station = "gilroy"

	timingSuffix      = " min."
	directionSelector = ".ipf-st-ip-trains-subtable"
	timingSelector    = ".ipf-st-ip-trains-subtable-td-arrivaltime"
)

// GetRealTimings returns duration of latest trains arriving at the specified
// station and going the specified direction. The Caltrain realtime page
// provides maximum of 3 upcoming trains.
//
// See http://www.caltrain.com/schedules/realtime/stations.html for list of
// stations.
//
// See http://www.caltrain.com/schedules.html if you want scheduled timetables.
func GetRealTimings(s Station, d Direction) (timings []time.Duration, err error) {
	doc, err := goquery.NewDocument(fmt.Sprintf(BaseURL, s))
	if err != nil {
		return nil, err
	}

	// Caltrain realtime page return data in a difficult to parse format:
	//
	//    <table class="ipf-st-ip-trains-subtable">
	//       <tr class="ipf-st-ip-trains-subtable-tr">
	//          ...
	//          <td class="ipf-st-ip-trains-subtable-td-arrivaltime">30 min.</td>
	//       </tr>
	//    </table>
	//    <table class="ipf-st-ip-trains-subtable">
	//       <tr class="ipf-st-ip-trains-subtable-tr">
	//          ...
	//          <td class="ipf-st-ip-trains-subtable-td-arrivaltime">10 min.</td>
	//       </tr>
	//    </table>
	//
	// Since the headers are in a different row than the data, we assume the first
	// table contains SouthBound timings and the last one contains NorthBound.
	//
	// TODO: incorporate alerts

	var index = 0
	if d == NorthBound {
		index = 1
	}

	doc.Find(directionSelector).Eq(index).Each(func(i int, s *goquery.Selection) {
		s.Find(timingSelector).Each(func(j int, t *goquery.Selection) {
			if min, err := parseStrIntoTime(t.Text()); err == nil {
				timings = append(timings, min)
			}
		})
	})

	return timings, nil
}

func parseStrIntoTime(str string) (time.Duration, error) {
	raw := strings.TrimSuffix(str, timingSuffix)
	min, err := strconv.Atoi(raw)
	if err != nil {
		return 0 * time.Minute, err
	}

	return time.Duration(min) * time.Minute, nil
}
