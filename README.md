# microtest-isam-to-innodb
This is the repository where I will work to prove out a golang script that will alter mysql tables on the fly from myisam to innodb

## Binaries vs Code
So I created the binary file here using `go build`. The source code is in the `main.go` (with the dependencies in the src/github folder).
If you want to get just the binary- utilize the  **altertbl** binary here.

## How to use altertbl
So the script requires command line arguments to work. All of them are mandatory. They are:

* host
* table
* user
* pwd
* db

Basically the idea is that you put the table name and login information there one at a time and we can then run the script. The code should look like this: `./altertbl -host="" -table=totalexperiment -user=fern -pwd=password -db=ferndb`

The script should return the exit message at the end.

## How to update the script
So if you make alterations to the main.go script (or any other script you add in) then first check that everything works (from the root directory here): `go run src/main.go`
Now, assuming everything worked, let's rebuild the binary: `cd src/ && go build -o altertbl src/main.go` and then run it with `./altertbl -host="" -table=totalexperiment -user=fern -pwd=password -db=ferndb`

## A few other notes

Don't put quotation marks around the command line arguments. 
Contact fernincornwall@gmail.com with questions




