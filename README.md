# ozxygen
Generate random OSM URL for load testing purpose

## Get Ozxygen

* Install a valid Go environnement
* `go get github.com/yanc0/ozxygen`

## Usage

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

## Example

    ozxygen -baseurl http://tileserver.com/ -num 5 -z 10
    
    http://tileserver.com/10/519/353.png
    http://tileserver.com/10/519/351.png
    http://tileserver.com/10/519/351.png
    http://tileserver.com/10/516/351.png
    http://tileserver.com/10/519/353.png

## Licence

    The MIT License (MIT)
    
    Copyright (c) 2016 Yann Coleu
    
    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:
    
    The above copyright notice and this permission notice shall be included in all
    copies or substantial portions of the Software.
    
    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
    SOFTWARE.
