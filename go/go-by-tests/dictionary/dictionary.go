package dictionary

const (
    MissingKeyError = DictionaryError("search key not found in dictionary")
    KeyExistsError = DictionaryError("key already exists, cannot add")
    KeyNotFoundError = DictionaryError("cannot update a key that does not exist")
    DeleteKeyMissing = DictionaryError("cannot delete a key that does not exist")
)

type Dictionary map[string]string
type DictionaryError string

func (e DictionaryError) Error() string {
    return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
    if value := d[key]; value == "" {
        return "", MissingKeyError
    }

	return d[key], nil
}

func (d Dictionary) Add(key string, value string) error {
    _, err := d.Search(key)
    switch err {
    case MissingKeyError:
        d[key] = value
    case nil:
        return KeyExistsError
    default:
        return err
    }

    return nil
}

func (d Dictionary) Update(key string, value string) error {
    _, err := d.Search(key)
    switch err {
    case MissingKeyError:
        return KeyNotFoundError
    case nil:
        d[key] = value
    default:
        return err
    }

    return nil
}

func (d Dictionary) Delete(key string) error {
    _, err := d.Search(key)
    switch err {
    case nil:
        delete(d, key)
    case MissingKeyError:
        return DeleteKeyMissing
    default:
        return err
    }

    return nil
}
