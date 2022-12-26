package singleton

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

var once sync.Once
var instance *singletonDatabase

// think of a module as a singleton
type Database interface {
	GetPopulation(name string) int
}

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// 2 ways: sync.Once init() -- thread safety
// laziness

func readData(path string) (map[string]int, error) {
	ex, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file, err := os.Open(ex + path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		caps, e := readData("/singleton/capitals.txt")
		db := singletonDatabase{capitals: caps}
		if e == nil {
			db.capitals = caps
		}
		instance = &db
	})
	return instance
}

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase().GetPopulation(city)
	}
	return result
}

func GetTotalPopulationEx(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	}
	return result
}

type DummyDatabase struct {
	dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3}
	}
	return d.dummyData[name]
}

func ExecuteSingleton() {
	/*
		Singleton is not a bad pattern, the bad part is depending directly on the singleton as opposed to
		depending on some interface
	*/

	// pop := GetTotalPopulationEx(GetSingletonDatabase(), "Seoul")
	// fmt.Println("Pop of Seoul = ", pop)

	names := []string{"alpha", "gamma"}
	tp := GetTotalPopulationEx(&DummyDatabase{}, names)
	fmt.Println(tp == 4)
}
