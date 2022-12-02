package solid

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0

// Journal has only one responsibility, manage entries
type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)

	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

// Separation of concerns
// God Objects (BAD) we are going to see how to break the SRP

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {
	// br
}

/*
	To this point journal is doing many things
	we are braking the single responsibility principle of journal, why?
	because adding these functions journal has another responsibility, persistence
	the responsibility oh the journal is not to deal with persistence can be handled by a separate component
*/

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func ExecuteSRP() {
	j := Journal{}
	j.AddEntry("I cried today")
	j.AddEntry("I ate a bug")
	fmt.Println(j.String())

	//
	// SaveToFile(&j, "journaltxt")

	//
	p := Persistence{lineSeparator: "\r\n"}
	p.SaveToFile(&j, "journal.txt")
}
