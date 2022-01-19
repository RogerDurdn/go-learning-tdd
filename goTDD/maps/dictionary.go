package dictionary

type Dictionary map[string]string

const (
	ErrorWordDoesNotExist = DictionaryError("could not find word")
	ErrorWordExists       = DictionaryError("the word already exists")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorWordDoesNotExist
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorWordDoesNotExist:
		d[word] = definition
	case nil:
		return ErrorWordExists
	default:
		return err

	}
	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	if _, err := d.Search(word); err != nil {
		return err
	}
	d[word] = newDefinition
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
