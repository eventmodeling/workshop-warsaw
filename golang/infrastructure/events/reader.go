package events

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/eventmodeling/workshop-warsaw/register/app/register"

	"github.com/pkg/errors"
)

type Reader struct {
	EventsDirectory string
}

// ByName returns all events for a given event name, sorted in ascending order by creation date.
func (r Reader) ByName(eventName string) (register.EventsCursor, error) {
	matches, err := filepath.Glob(fmt.Sprintf("%s/%s_*", r.EventsDirectory, eventName))
	if err != nil {
		return nil, errors.Wrap(err, "could not find events by name")
	}

	if len(matches) == 0 {
		return nil, register.ErrNoEventsByName
	}

	return &Cursor{
		current:   0,
		filenames: matches,
	}, nil
}

// Cursor iterates over events and allows to unmarshal them
type Cursor struct {
	current   int
	filenames []string
}

// Len returns the length of the cursor.
func (c *Cursor) Len() int {
	return len(c.filenames)
}

// Next bumps the iterator and returns false if there is no more events
func (c *Cursor) Next() bool {
	c.current++
	if c.current == len(c.filenames) {
		return false
	}
	return true
}

// Unmarshal unmarshals the current event on the given structure.
func (c *Cursor) Unmarshal(val interface{}) error {
	f, err := os.Open(c.filenames[c.current])
	if err != nil {
		return errors.Wrap(err, "error reading event")
	}
	defer func() {
		closeErr := f.Close()
		if closeErr != nil {
			log.Printf("error closing file: %+v", err)
		}
	}()

	b, err := ioutil.ReadFile(c.filenames[c.current])
	if err != nil {
		return errors.Wrap(err, "error reading event")
	}

	err = json.Unmarshal(b, val)
	if err != nil {
		return errors.Wrap(err, "error unmarshaling event")
	}
	return nil
}
