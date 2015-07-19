package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"net/http"
        "github.com/kellydunn/golang-geo"
	"strconv"
	"encoding/csv"
	"os"
	"strings"
)



func main() {

	var bootstrap_start, html_end, stopsdromologio string
	//var html_start string
	var tmpcount int

	//html_start = "<html><head><style>img.top{vertical-align:text-top}img.bottom{vertical-align:text-bottom}img.middle{vertical-align:text-middle}.Absolute-Center{margin:auto;position:absolute;top:0;left:0;bottom:0;right:0}.demo{color:red;text-align:center;font-size:100%}</style><script type='text/javascript'>function loadUrl(newLocation){window.location=newLocation;return false;}</script></head><body><div id='busbeat' align='center'><img src='http://www.busbeat.com/busbeat.png' onclick='getLocation()'/></div><p id='demo' class='demo'></p><p id='myDiv' class='demo'></p><script>var x=document.getElementById('demo');function getLocation(){if(navigator.geolocation){navigator.geolocation.getCurrentPosition(showPosition);}else{x.innerHTML='Geolocation is not supported by this browser.';}}function showPosition(position){url='http://88.198.100.100:3000/stasis?lon='+position.coords.longitude+'&lat='+position.coords.latitude;loadUrl(url);}</script>"
	html_end = "</ul>Build with <span class='glyphicon glyphicon-heart' aria-hidden='true'></span> for Crowdhackathon</div></body></html>"


	bootstrap_start="<!DOCTYPE html><html lang='en'><head><title>Busbeat</title><meta charset='utf-8'><meta name='viewport' content='width=device-width, initial-scale=1'><link rel='stylesheet' href='http://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css'><script src='https://ajax.googleapis.com/ajax/libs/jquery/1.11.3/jquery.min.js'></script><script src='http://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js'></script></head><body><p id='demo'></p><div class='container'><div class='text-center'><a href='http://www.busbeat.com/beta'><img src='http://www.busbeat.com/busbeat.png' alt='Busbeat'/></a></div>"



	csvfile, err := os.Open("stops.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()


	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	

	stopsdromologiafile, err := os.Open("stop_dromologia_sunday_ola_koma.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer stopsdromologiafile.Close()

	stopsdromologiareader := csv.NewReader(stopsdromologiafile)

	stopsdromologiareader.FieldsPerRecord = -1

	stopsdromologiaCSVdata, err := stopsdromologiareader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}


        routescsvfile, err := os.Open("routes.csv")

        if err != nil {
                fmt.Println(err)
                return
        }

        defer routescsvfile.Close()

        routesreader := csv.NewReader(routescsvfile)

        routesreader.FieldsPerRecord = -1

        routesCSVdata, err := routesreader.ReadAll()

        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }


	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Get("/stasis", func(req *http.Request) string {

		var response string
		var i int
		var row_min int

		response = "GET\n\n"

		user_lon := req.URL.Query().Get("lon")
		if len(user_lon) != 0 {
			user_lat := req.URL.Query().Get("lat")
			if len(user_lat) != 0 {
     				// Make a few points
				f_lon,_:=strconv.ParseFloat(user_lon,64)
				f_lat,_:=strconv.ParseFloat(user_lat,64)
     				p := geo.NewPoint(f_lat, f_lon)
				f2_lon,_ := strconv.ParseFloat(rawCSVdata[1][5],64)
				f2_lat,_ := strconv.ParseFloat(rawCSVdata[1][4],64)
				p2 := geo.NewPoint(f2_lat,f2_lon)
				dist_min := p.GreatCircleDistance(p2)
				i=1
				row_min=i
				for _,each := range rawCSVdata{
					f2_lon,_:=strconv.ParseFloat(each[5],64)
					f2_lat,_:=strconv.ParseFloat(each[4],64)
					p2 := geo.NewPoint(f2_lat,f2_lon)
					dist := p.GreatCircleDistance(p2)
					if dist<dist_min {
						dist_min = dist
						row_min = i
					}
					i++
				}
				//response = fmt.Sprintf("great circle distance: %.1fKm %s\n", dist,rawCSVdata[1])
				response = bootstrap_start + fmt.Sprintf("<h3 align='center'>%s, %s (%.1fKm)</h4><ul class='list-group'>", rawCSVdata[row_min][2],rawCSVdata[row_min][3],dist_min)


/*<h2>List Group With Badges</h2>
  <ul class="list-group">
    <li class="list-group-item"><span class="badge">12</span> New</li>
    <li class="list-group-item"><span class="badge">5</span> Deleted</li>
    <li class="list-group-item"><span class="badge">3</span> Warnings</li>
  </ul>
</div>
*/
				for _,stopsdromologiaeach := range stopsdromologiaCSVdata{
					if(len(stopsdromologiaeach) == 6) {
						stopsdromologio = stopsdromologiaeach[0]
						if (strings.EqualFold(stopsdromologio,rawCSVdata[row_min][0])==true){
							tmpcount = 0
							for _,eachroute := range routesCSVdata{
								if (strings.EqualFold(stopsdromologiaeach[2],eachroute[1])==true){
									response += fmt.Sprintf("<li class='list-group-item'><span class='badge'>%s λεπτά</span><span class='badge'>3 προσφορές</span>%s %s</li>",stopsdromologiaeach[5],stopsdromologiaeach[2],eachroute[2])
									tmpcount = 1
								}
							}
							if (tmpcount == 0){
									response += fmt.Sprintf("<li class='list-group-item'><span class='badge'>%s λεπτά</span>%s</li>",stopsdromologiaeach[5],stopsdromologiaeach[2])
							}
						}
					}

                                        if(len(stopsdromologiaeach) == 7) {
                                                stopsdromologio = stopsdromologiaeach[0]
                                                if (strings.EqualFold(stopsdromologio,rawCSVdata[row_min][0])==true){
							tmpcount = 0
							for _,secondeachroute := range routesCSVdata{
                                                                if (strings.EqualFold(stopsdromologiaeach[2],secondeachroute[1])==true){
                                                        		response += fmt.Sprintf("<li class='list-group-item'><span class='badge'>%s λεπτά</span>%s %s %s</li>",stopsdromologiaeach[6],stopsdromologiaeach[2],secondeachroute[2])
									tmpcount = 1
								}

                                                               	if (strings.EqualFold(stopsdromologiaeach[3],secondeachroute[1])==true){
                                                        			response += fmt.Sprintf("<li class='list-group-item'><span class='badge'>%s λεπτά</span>%s %s</li>",stopsdromologiaeach[6],stopsdromologiaeach[3],secondeachroute[2])
										tmpcount = 2
								}
							}
							if (tmpcount == 0){
									response += fmt.Sprintf("<li class='list-group-item'><span class='badge'>%s λεπτά</span>%s </li></center>",stopsdromologiaeach[6],stopsdromologiaeach[2])
									response += fmt.Sprintf("<li class='list-group-item'><span class='badge'>%s λεπτά</span>%s </li></center>",stopsdromologiaeach[6],stopsdromologiaeach[3])
							}

							if (tmpcount == 1){
									response += fmt.Sprintf("<li class='list-group-item'><span class='badge'>%s λεπτά</span>%s </li></center>",stopsdromologiaeach[6],stopsdromologiaeach[3])
							}
                                                }
                                        }

				}
				response += html_end
			}
		}
		return response
	})

	m.Run()
}

/*func search_closest_stop(lon,lat float64)
{
	p := geo.NewPoint(lon,lat)

	return "Κοντινότερη στάση, Απόσταση"
}*/

/*
func get_buses(stopid)
{
	return "Λεωφορείο"
}

func get_trip(tripid)
{
	return "Δρομολόγιο"
}

func get_events(tripid)
{
	return "Προσφορές στο τρέχον δρομολόγιο"
}*/
