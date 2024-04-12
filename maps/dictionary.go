package maps

const (
	ErrNotFound         = DictErr("could not find the word you were looking for")
	ErrWordExists       = DictErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictErr("cannot update word because it does not exist")
)

type DictErr string

func (e DictErr) Error() string {
	return string(e)
}

type Dict map[string]string

func (d Dict) Search(key string) (string, error) {
	val, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}

	return val, nil
}

func (d Dict) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dict) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}
