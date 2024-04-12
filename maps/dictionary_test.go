package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := Dict{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dict{}
		k := "test"
		v := "this is just a test"

		err := dict.Add(k, v)

		assertDefinition(t, dict, k, v)
		assertError(t, err, nil)
	})
	t.Run("existing word", func(t *testing.T) {
		k := "test"
		v := "this is just a test"
		dict := Dict{k: v}

		err := dict.Add(k, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, k, v)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		k := "test"
		v := "this is just a test"
		dict := Dict{k: v}
		newV := "new value"

		err := dict.Update(k, newV)

		assertError(t, err, nil)
		assertDefinition(t, dict, k, newV)
	})
	t.Run("new word", func(t *testing.T) {
		k := "test"
		v := "this is just a test"
		dict := Dict{}

		err := dict.Update(k, v)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("ogt %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dict Dict, k, v string) {
	t.Helper()

	got, err := dict.Search(k)

	if err != nil {
		t.Fatal("shold find added word: ", err)
	}

	assertStrings(t, got, v)
}
