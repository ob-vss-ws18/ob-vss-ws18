package moviestore

import (
	"errors"
)

type Age uint8
type FSK uint8
type Serial uint

var i Serial = 0
var j UserID = 0

const (
	FSK0  = 0
	FSK6  = 6
	FSK12 = 12
	FSK16 = 16
	FSK18 = 18
)

type Moviestore interface {
	AddMovie(string, FSK) Serial
	AddUser(string, Age) UserID
	Rent(Serial, UserID) (User, Movie, error)
	Return(Serial) (User, Movie, error)
	RentedByUser(UserID) ([]Movie, error)
}

type moviestoreImpl struct {
	available  map[Serial]Movie
	rented     map[UserID][]Movie
	users      map[UserID]User
	nextSerial Serial
	nextUserID UserID
}

type Movie struct {
	Title  string
	Fsk    FSK
	Serial Serial
}

func NewMoviestore() Moviestore {
	return &moviestoreImpl{}
}

func AllowedAtAge(m *Movie, age Age) bool {
	if uint8(m.Fsk) >= uint8(age) {
		return true
	}
	return false
}

func (ms *moviestoreImpl) AddMovie(title string, fsk FSK) Serial {
	var temp = i
	i = i + 1
	ms.available[temp] = Movie{Title: title, Fsk: fsk, Serial: temp}
	return temp
}

func (ms *moviestoreImpl) AddUser(name string, age Age) UserID {
	var temp = j
	j = j + 1
	ms.users[temp] = User{Name: name, Age: age, UserID: temp}
	return temp
}

func (ms *moviestoreImpl) Rent(serial Serial, userID UserID) (User, Movie, error) {
	var err error
	var user User
	var movie Movie
	return user, movie, err
}

func (ms *moviestoreImpl) RentedByUser(userID UserID) ([]Movie, error) {
	var err = errors.New("userID unknown")
	var moviesList = ms.rented[userID]
	if(moviesList) == nil {
		return moviesList, err
	}
	return moviesList, nil
}

func (ms *moviestoreImpl) Return(serial Serial) (User, Movie, error) {
	var movieToReturn Movie
	var userToReturn User
	for userID, movies := range ms.rented {
		if contains(movies, serial) {
			movieToReturn = movies[getSlicePos(movies, serial)]
			userToReturn = ms.users[userID]
			movies[getSlicePos(movies, serial)] = movies[len(movies)-1]
			movies = movies[:len(movies)-1]
			return userToReturn, movieToReturn, nil
		}
	}
	var err = errors.New("Movie not found")
	return userToReturn, movieToReturn, err
}

// **************************
// additional helper methods
// **************************
func contains(s []Movie, e Serial) bool {
	for _, a := range s {
		if a.Serial == e {
			return true
		}
	}
	return false
}

func getSlicePos(m []Movie, s Serial) int {
	var pos = 0
	for _, movie := range m {
		if movie.Serial == s {
			return pos
		}
		pos += 1
	}
	return -1
}

func main() {

	m := NewMoviestore()
	m.AddUser("a", 4)

}
