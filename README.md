# Go Coding Challenge

### Create an API with two endpoints:
1. countcaps/{someString}
    a. Takes a string as input
    b. Returns the number of capital letters in the string.
    c. Write unit tests.
 2. weather/favorites
    a.  Takes no input
    b. Looks up the weather in 5 zipcodes using openweather.org
    c. Use go routines
    d. Use channels
    e. Returns a JSON with an array of zipcode/tempature elements
___
## How do I run it?

  - First things first: create a .env file at the root of the directory contatining the OWM_DEVELOPER_KEY
  - While in the root of the directory, start the server with the following:
```sh
$ go run main.go
```

## Routes:
* (GET) localhost:8080/api/weather/favorites
* (GET) localhost:8080/api/countcaps/{word}
