package events

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/eventmodeling/workshop-warsaw/register/app/register"

	"github.com/pkg/errors"
)

type Reader struct {
	EventsDirectory string
}

type eventFile struct {
	filename  string
	createdAt time.Time
}

type eventFiles []eventFile

func (ef eventFiles) Len() int           { return len(ef) }
func (ef eventFiles) Swap(i, j int)      { ef[i], ef[j] = ef[j], ef[i] }
func (ef eventFiles) Less(i, j int) bool { return ef[i].createdAt.Before(ef[j].createdAt) }

// ByName returns all events for a given event name, sorted in ascending order by creation date.
func (r Reader) ByName(eventName string) (register.EventsCursor, error) {
	matches, err := filepath.Glob(fmt.Sprintf("%s/%s_*", r.EventsDirectory, eventName))
	if err != nil {
		return nil, errors.Wrap(err, "could not find events by name")
	}

	if len(matches) == 0 {
		return nil, register.ErrNoEventsByName
	}

	eventFiles := make(eventFiles, len(matches))
	for i, filename := range matches {
		fi, err := os.Stat(filename)
		if err != nil {
			return nil, errors.Wrap(err, "could not stat a file")
		}

		eventFiles[i] = eventFile{
			filename: filename,
			// use ModTime, Ctime would require os-dependent stuff, we don't wanna go there yet
			// events should not be modified once created, anyway
			createdAt: fi.ModTime(),
		}
	}
	sort.Sort(eventFiles)

	return &Cursor{
		current:    0,
		eventFiles: eventFiles,
	}, nil
}

// Cursor iterates over events and allows to unmarshal them
type Cursor struct {
	current    int
	eventFiles eventFiles
}

// Len returns the length of the cursor.
func (c *Cursor) Len() int {
	return len(c.eventFiles)
}

// Next bumps the iterator and returns false if there is no more events
func (c *Cursor) Next() bool {
	c.current++
	if c.current == len(c.eventFiles) {
		return false
	}
	return true
}

// Unmarshal unmarshals the current event on the given structure.
func (c *Cursor) Unmarshal(val interface{}) error {
	f, err := os.Open(c.eventFiles[c.current].filename)
	if err != nil {
		return errors.Wrap(err, "error reading event")
	}
	defer func() {
		closeErr := f.Close()
		if closeErr != nil {
			log.Printf("error closing file: %+v", err)
		}
	}()

	b, err := ioutil.ReadFile(c.eventFiles[c.current].filename)
	if err != nil {
		return errors.Wrap(err, "error reading event")
	}

	err = json.Unmarshal(b, val)
	if err != nil {
		return errors.Wrap(err, "error unmarshaling event")
	}
	return nil
}
