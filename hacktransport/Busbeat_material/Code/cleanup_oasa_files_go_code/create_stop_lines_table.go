package main

import (
	"fmt"
	//"strconv"
	"encoding/csv"
	"strings"
	"os"
)


func main() {

	var stopid, stop_times_stopid, tmpstr string

	//var counter int

	stopsfile, err := os.Open("stops.csv")

	if err != nil {
		fmt.Println(err)
		return
	}


	//counter = 0
	defer stopsfile.Close()

	stopsreader := csv.NewReader(stopsfile)

	stopsreader.FieldsPerRecord = -1

	stopsCSVdata, err := stopsreader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

        /*routescsvfile, err := os.Open("routes.csv")

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
*/
        stop_timescsvfile, err := os.Open("stop_times.csv")

        if err != nil {
                fmt.Println(err)
                return
        }
       
        defer stop_timescsvfile.Close()

        stop_timesreader := csv.NewReader(stop_timescsvfile)

        stop_timesreader.FieldsPerRecord = -1

        stop_timesCSVdata, err := stop_timesreader.ReadAll()

        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

	//stop_timesCSVdata
	//trip_id,stop_id,stop_sequence,pickup_type,drop_off_type
	//6351001-ΤΗΛΕΜΑ-Τ3-Παρασκευή-01,400045,1,0,0

	///routesCSVdata
	////route_id,route_short_name,route_long_name,route_desc,route_type,route_color, route_text_color
	//Τ3-20,Τ3,"ΝΕΟ ΦΑΛΗΡΟ - ΒΟΥΛΑ",,2,153CE0,FFFFFF

	//stopsCSVdata
	//stop_id,stop_code,stop_name,stop_desc,stop_lat,stop_lon,location_type
	//010001,010001,ΣΤΡΟΦΗ,Επί της ΕΛ.ΒΕΝΙΖΕΛΟΥ,37.998608264,23.664984624,0
	//fmt.Printf("%s\n",stop_timesCSVdata[1])

	for _,eachstops := range stopsCSVdata{
		stopid = eachstops[1]
		for _,eachstop_times := range stop_timesCSVdata{
			stop_times_stopid = eachstop_times[1]
			if (strings.EqualFold(stop_times_stopid,stopid)==true) {
			//	for _,each := range routesCSVdata{
			//		routeid = each[1]
			//		if (strings.Contains(eachstop_times[0],routeid)==true){
			//			fmt.Printf("%s,%s,%s,%s,%s\n",stopid,routeid,each[2],eachstop_times[2],eachstop_times[1])
			//			break
			//		}
			//	}
				//fmt.Printf("%s\n",eachstop_times[0][15:])
				if (strings.EqualFold(tmpstr,eachstop_times[0][8:])==false){
					//for _,each := range routesCSVdata{
					//	routeid = each[1]
					//	if (strings.Contains(eachstop_times[0][15:],routeid)==true){
					//		fmt.Printf("%s,%s,%s\n",stopid,routeid,eachstop_times[1])
					fmt.Printf("%s,%s,%s\n",stopid,eachstop_times[0][8:],eachstop_times[2])
					//	}
					//}
					tmpstr = eachstop_times[0][8:]
				}
			}
			//fmt.Printf("%s\n",each[1])
		}
		//counter++
		//if (counter==3){
		//	fmt.Printf("%s\n",eachstops[1])
		//	break
		//}
	}

	//for _,each := range routesCSVdata{
	//	fmt.Printf("%s\n",each[1])
	//}
	//fmt.Printf("%s\n",stopsCSVdata[1])
	
}
