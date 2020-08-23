package database

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type jsonDatabase struct {
	filename string

	// [
	//   {Id: xxx, ...},
	// ]
	raw     []json.RawMessage
	mapping map[interface{}]int
}

func (j *jsonDatabase) List(data interface{}) error {
	b, err := json.Marshal(j.raw)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &data)
}

func (j *jsonDatabase) Patch(key interface{}, b []byte) {
	index, found := j.mapping[key]
	if !found {
		j.Insert(key, b)
		return
	}
	j.raw[index] = json.RawMessage(b)
}

func (j *jsonDatabase) Insert(key interface{}, b []byte) {
	j.raw = append(j.raw, json.RawMessage(b))
	j.mapping[key] = len(j.raw) - 1
}

func (j *jsonDatabase) Commit() error {
	b, err := json.MarshalIndent(j.raw, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(j.filename, b, 0755)
}

func newJSONDatabase(filename string) (*jsonDatabase, error) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	if string(b) == "" {
		b = []byte("[]")
	}

	var raw []json.RawMessage
	if err := json.Unmarshal(b, &raw); err != nil {
		return nil, err
	}

	db := &jsonDatabase{
		filename: filename,
		raw:      raw,
		mapping:  make(map[interface{}]int),
	}

	var idMapping []struct{ Id int }
	if err := json.Unmarshal(b, &idMapping); err != nil {
		return nil, err
	}
	for index, data := range idMapping {
		db.mapping[data.Id] = index
	}

	return db, nil
}
