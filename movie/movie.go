package movie

import (
        "encoding/json"
        "fmt"
        "going/httpclient"
        "reflect"
)

// Container generic
type Container []interface{}

// MovieInterfaces interface
type MovieInterfaces interface {
        Rename(newName string)
}

// MovieInfo is Movie Information Struct
type MovieInfo struct {
        title    string
        director string
        cast     string
        plot     string
        cost     int32
        revenue  int32
        rating   float32
        rated    int
}

// Actor of the movie
type Actor struct {
        firstName string
        lastName  string
        age       int
        gender    string
}

// RatingChecker - for ratings
type RatingChecker func(float32)

// Weekday stuff
type Weekday int

const (
        Sunday Weekday = iota + 1
        Monday
        Tuesday
        Wednesday
        Thursday
        Friday
        Staurday
)

const (
        MOVIETYPE = "Comedy"  // MOVIETYPE is just that
        LANGUAGE  = "English" // Language
        RATING    = "R"       // Rating
)

var title, director, cast, plot string
var planeEaterWikiURL = "https://en.wikipedia.org/wiki/Michel_Lotito"

func RunMain() {

        printCross(15)

        var ratingCalc = makeRatingFunction()

        var thisMovie = initializeMovie()

        otherItem1, otherItem2, otherItem3, otherItem4, otherItem5 := alsoAte()

        fmt.Println("Title: \t\t", thisMovie.title)
        fmt.Println("Directed by:\t", thisMovie.director)
        fmt.Println("Featuring:\t", thisMovie.cast)
        fmt.Println("Synopsis:\t", thisMovie.plot)

        fmt.Println("He also ate:\t", otherItem1, otherItem2, otherItem3, otherItem4, otherItem5)

        fmt.Println("Language:\t", LANGUAGE)
        fmt.Println("Rated For:\t\t", getRatingText(thisMovie.rated))

        fmt.Println("User Rating:\t\t", thisMovie.rating)
        fmt.Println("Released on Day:\t", Tuesday)

        fmt.Println("Profits:\t", thisMovie.revenue-thisMovie.cost)

        // fmt.Println("Mem Addr for aString:\t", &aString)
        // fmt.Println("Mem Addr for thisMovie:\t", &thisMovie)

        //fmt.Println("makeRatingFunction.Returned:\t", makeRatingFunction(thisMovie.title))

        ratingCalc(thisMovie.rating)

        printType(1.5)
        paintStars(10)

        testMap()

        messWithSlices()

        //var myClient = httpclient.Client

        remoteData := httpclient.HttpGet(planeEaterWikiURL)
        fmt.Println(remoteData)

        paintStars(50)

        thisMovie.PrintMovie()

        //      fmt.Println("%", thisMovie)
        PrintObject(thisMovie)

        //RenameMovie()

        ba, err := json.Marshal(thisMovie)
        if err != nil {
                fmt.Println(err)
                return
        }
        //fmt.Println("%+v", b)
        //fmt.Println("%+v", object)
        fmt.Println(string(ba))
        fmt.Println("typeof: ", reflect.TypeOf(thisMovie))
}

// PrintObject converts objects to JSON string and prints them.
func PrintObject(object interface{}) {

        fmt.Println("type: ", reflect.TypeOf(object))

        b, err := json.Marshal(object)
        if err != nil {
                fmt.Println(err)
                return
        }
        //fmt.Println("%+v", b)
        //fmt.Println("%+v", object)
        fmt.Println(string(b))
}

