package main

import "testing"


func TestSearch(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		
		dictionary := Dictionary{}

		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil) 
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

// func TestAdd(t *testing.T) {
	
// 	dictionary := Dictionary{}
// 	word := "test"
// 	definition := "this is just a test"

// 	dictionary.Add(word, definition);

// 	assertDefinition(t, dictionary, word, definition)
// 	// want := "this is just a test"

// 	// got, err := dictionary.Search("test")

// 	// if err != nil {
// 	// 	t.Fatal("should find added word:", err)
// 	// }

// 	// if got != want {
// 	// 	t.Errorf("got %q want %q", got, want)
// 	// }
// }

// func TestSearch(t *testing.T) {
	
// 	dictionary := Dictionary{"test": "this is just a test"}

// 	t.Run("known word", func(t *testing.T) {

// 		got, _ := dictionary.Search("test")

// 		want := "this is just a test"

// 		assertStrings(t, got, want)
// 	})

// 	t.Run("unknown word", func(t *testing.T) {

// 		_, got := dictionary.Search("unknown")

// 		// if err == nil {
// 		// 	t.Fatal("expected to get an error.")
// 		// }

// 		// assertStrings(t, err.Error(), want)

// 		assertError(t, got, ErrNotFound)
// 	})

// 	// got := dictionary.Search("test")

// 	// want := "this is just a test"

// 	// assertStrings(t, got, want)
// }
// func TestSearch(t *testing.T) {

// 	dictionary := map[string]string{"test": "this is just a test"}

// 	got := Search(dictionary, "test")

// 	want := "this is just a test"

// 	if got != want {
// 		t.Errorf("got %q want %q given, %q", got, want, "test")
// 	}
// }

func assertStrings(t testing.TB, got, want string) {

	t.Helper()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {

	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {

	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word", err)
	}

	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}
}