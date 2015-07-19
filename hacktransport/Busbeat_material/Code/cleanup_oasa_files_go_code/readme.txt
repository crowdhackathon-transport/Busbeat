--------------------------------------------
Readme for the cleanup_oasa_file_go_code
--------------------------------------------

A small go simple app was created to fix 
the dataset and make it more usable. 

The stops_times file had to be cleaned up
from duplicates because it was not usable.

To execute the cleanup execute the following

go run create_stop_lines_table.go 

Make sure all the OASA csv files are inside.


After executing the file will be cleaned up 
but additional cleanup is required i.e. For 
different weekdays the bus stops change. This 
needs to be considered by the app. 

Since this is a proof of concept I did a 
grep exctracting the Sunday routes for today
demo

The following data where used 

http://data.gov.gr/dataset/28/

and must be on the same folder extracted.
