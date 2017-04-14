package jsondb

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"domain"
)

// Schema represents the database structure.
type Schema struct {
	Records []User
}

// Client represents a client to the data store.
type Client struct {
	// Path is the relative filename.
	Path string

	data  *Schema
	mutex sync.RWMutex

	latestID int
}

// NewClient returns a new database client.
func NewClient(path string) *Client {
	c := &Client{
		Path: path,
		data: new(Schema),
	}

	return c
}

// Reads opens/initializes the database.
func (c *Client) Load(output func(err error)) {
	var err error
	var b []byte

	c.mutex.Lock()

	if _, err = os.Stat(c.Path); os.IsNotExist(err) {
		err = ioutil.WriteFile(c.Path, []byte("{}"), 0644)
		if err != nil {
			c.mutex.Unlock()
			//c.output.Error(err)
			output(err)
			return
		}
	}

	b, err = ioutil.ReadFile(c.Path)
	if err != nil {
		c.mutex.Unlock()
		output(err)
		return
	}

	c.data = new(Schema)
	err = json.Unmarshal(b, &c.data)

	c.calculateLatestID()

	c.mutex.Unlock()

	if err != nil {
		output(err)
		return
	}

	output(nil)
}

func (c *Client) calculateLatestID() {
	for _, u := range c.data.Records {
		if u.ID > c.latestID {
			c.latestID = u.ID
		}
	}

}

// Write saves the database.
func (c *Client) Save(output func(err error)) {
	var err error
	var b []byte

	c.mutex.Lock()

	b, err = json.Marshal(c.data)
	if err != nil {
		c.mutex.Unlock()
		//c.output.Error(err)
		output(err)
		return
	}

	err = ioutil.WriteFile(c.Path, b, 0644)
	if err != nil {
		c.mutex.Unlock()
		//c.output.Error(err)
		output(err)
		return
	}

	c.mutex.Unlock()

	//c.output.DidSave()
	output(nil)
}

// AddRecord adds a record to the database.
func (c *Client) AddRecord(user domain.User, output func(user domain.User)) {

	u := (&User{}).from(user)
	c.latestID++ //dummy for incrementing ID's
	u.ID = c.latestID

	c.data.Records = append(c.data.Records, u)

	//c.output.DidAddUser(rec)
	output(u.toDomUser())
}


// Records retrieves all records from the database.
func (c *Client) Records(output func([]domain.User)) {
	//c.output.Records(c.data.Records)
	output(DBtoDomain(c.data.Records))
}
