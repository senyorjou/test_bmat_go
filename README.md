##README.  test_bmat_go

### Set Up
Clone repo into a golang workspace and from within directory:  
	
	$ go get
	$ go run *.go
	
A couple of executables will be provided for OSX and Linux architectures

Program tries to connect to a standard mongodb port and exits if not available, if succesful all data will be persisted on a newly created collection called `bmat-test`.
