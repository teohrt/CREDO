# Go Coding Challenge

### Create an API with two endpoints:
1. countcaps/{someString}
    * Takes a string as input
    * Returns the number of capital letters in the string.
    * Write unit tests.
 2. weather/favorites
    *  Takes no input
    * Looks up the weather in 5 zipcodes using openweather.org
    * Use go routines
    * Use channels
    * Returns a JSON with an array of zipcode/tempature elements
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
