package namepicker

import (
	"github.com/athom/rtrand"
	"io/ioutil"
	"os"
	"strings"
)

const (
	ALL_LAST_NAME_DB    = "dist.all.last"
	FEMAL_FIRST_NAME_DB = "dist.female.first"
	MALE_FIRST_NAME_DB  = "dist.male.first"
)

var dataDir = os.Getenv("GOPATH") + "/src/github.com/athom/namepicker/data"

func RandomName() (r string) {
	f, l := RandomNames()
	return f + ` ` + l
}

func RandomFemaleName() (r string) {
	f, l := RandomFemaleNames()
	return f + ` ` + l
}

func RandomMaleName() (r string) {
	f, l := RandomMaleNames()
	return f + ` ` + l
}

func RandomNames() (firstName string, lastName string) {
	if rtrand.Intn(2) == 0 {
		return RandomMaleNames()
	}
	return RandomFemaleNames()
}

func RandomFemaleNames() (firstName string, lastName string) {
	firstName = firstFemaleName()
	lastName = lastAllName()
	return
}

func RandomMaleNames() (firstName string, lastName string) {
	firstName = firstMaleName()
	lastName = lastAllName()
	return
}

func firstMaleName() (r string) {
	return pick(MALE_FIRST_NAME_DB)
}

func firstFemaleName() (r string) {
	return pick(FEMAL_FIRST_NAME_DB)
}

func lastAllName() (r string) {
	return pick(ALL_LAST_NAME_DB)
}

func pick(db string) (r string) {
	c := getFileContent(dataDir + "/" + db)
	rows := strings.Split(c, "\n")
	index := rtrand.Int() % (len(rows) - 2)
	row := rows[index]
	cols := strings.Split(row, " ")
	name := strings.ToLower(cols[0])
	r = strings.ToUpper(name[0:1]) + name[1:]
	return
}

func getFileContent(filename string) (r string) {
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	contents, err := ioutil.ReadAll(fd)
	if err != nil {
		panic(err)
	}
	r = (string)(contents)
	return
}
