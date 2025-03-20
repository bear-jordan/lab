package dictionary

import "testing"

func assertError(t *testing.T, got error, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("error found %q", err.Error())
	}
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, key string, want string) {
    t.Helper()

    got, err := dictionary.Search(key)
    assertNoError(t, err)
    assertStrings(t, got, want)
}

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	testTable := []struct {
		description string
		searchKey   string
		want        string
	}{
		{description: "test a valid key", searchKey: "test", want: "this is a test"},
    }

	for _, tt := range testTable {
		t.Run(tt.description, func(t *testing.T) {
			got, err := dictionary.Search(tt.searchKey)

            assertNoError(t, err)
			assertStrings(t, got, tt.want)
		})
	}

    t.Run("check for a missing key", func(t *testing.T) {
        _, err := dictionary.Search("oops")

        assertError(t, err, MissingKeyError)
    })
}

func TestAdd(t *testing.T) {
    t.Run("new value added", func(t *testing.T) {
        dictionary := Dictionary{}
        want := "this is a test"
        dictionary.Add("test", want)
        assertDefinition(t, dictionary, "test", want)
    })

    t.Run("should not overwrite new values", func(t *testing.T) {
        dictionary := Dictionary{}
        want := "should not be overwritten"
        err := dictionary.Add("test", want)
        assertNoError(t, err)
        err = dictionary.Add("test", "fail if this appears")
        assertError(t, err, KeyExistsError)
        assertDefinition(t, dictionary, "test", want)
    })
}

func TestUpdate(t *testing.T) {
    t.Run("update existing key with new value", func(t *testing.T) {
        want := "new value"
        dictionary := Dictionary{"test": "this is a test"}
        err := dictionary.Update("test", want)
        assertNoError(t, err)
        assertDefinition(t, dictionary, "test", want)
    })

    t.Run("cannot update an existing value", func(t *testing.T) {
        want := KeyNotFoundError
        dictionary := Dictionary{}
        err := dictionary.Update("test", "test")
        assertError(t, err, want)
    })
}

func TestDelete(t *testing.T) {
    t.Run("test valid deletion", func(t *testing.T) {
        dictionary := Dictionary{"test": "test value"}
        err := dictionary.Delete("test")
        assertNoError(t, err)
    })

    t.Run("test invalid deletion", func(t *testing.T) {
        dictionary := Dictionary{}
        err := dictionary.Delete("test")
        assertError(t, err, DeleteKeyMissing)
    })
}
