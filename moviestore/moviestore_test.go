package moviestore

import "testing"

func newMoviestoreImpl() *moviestoreImpl {
	ms := new(moviestoreImpl)
	ms.available = make(map[Serial]Movie)
	ms.users = make(map[UserID]User)
	ms.rented = make(map[UserID][]Movie)
	return ms
}

func TestMoviestore_AddMovie(t *testing.T) {
	cases := []struct {
		title string
		fsk   FSK
	}{
		{"Am Limit", FSK0},
		{"Texas Chainsaw Massacre", FSK18},
		{"Inglourious Basterds", FSK16},
	}
	moviestore := newMoviestoreImpl()
	for _, c := range cases {
		gotSerial := moviestore.AddMovie(c.title, c.fsk)
		gotMovie, present := moviestore.available[gotSerial]
		if !present || gotMovie.Title != c.title || gotMovie.Fsk != c.fsk {
			t.Errorf("Added %v, but got %v", c, gotMovie)
		}
	}
}

func TestMoviestore_AddUser(t *testing.T) {
	cases := []struct {
		name string
		age Age
	}{
		{"Hugo", 12},
		{"Helga", 19},
		{"Ronja", 17},
	}
	moviestore := newMoviestoreImpl()
	for _, c := range cases {
		gotUserID := moviestore.AddUser(c.name, c.age)
		gotUser, present := moviestore.users[gotUserID]
		if !present || gotUser.Name != c.name || gotUser.Age != c.age {
			t.Errorf("Added %v, but got %v", c, gotUser)
		}
	}
}
