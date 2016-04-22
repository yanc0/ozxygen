# ozxygen
Generate random OSM URL for load testing purpose

```
Usage of ozxygen:
  -baseurl string
    	HTTP base url (default "http://127.0.0.1/osm/")
  -extension string
    	End of the URI (default ".png")
  -maxlat float
    	Max latitude (default 49.2)
  -maxlon float
    	Max longitude (default 2.98)
  -minlat float
    	Min latitude (default 48.5)
  -minlon float
    	Min longitude (default 1.69)
  -num int
    	Number of results (default 1)
  -z int
    	Zoom level (default 18)

```

```
ozxygen -baseurl http://tileserver.com/ -num 5 -z 10

http://tileserver.com/10/519/353.png
http://tileserver.com/10/519/351.png
http://tileserver.com/10/519/351.png
http://tileserver.com/10/516/351.png
http://tileserver.com/10/519/353.png

```
