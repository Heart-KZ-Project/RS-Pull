package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	MIN = 1
	MAX = 401
	LEN = 20
)

var (
	ROOT       = "Random-Statistics-Data"
	MY_PATH    = ""
	SEED       = time.Now().UnixNano()
	files, err = ioutil.ReadDir(ROOT)
	SET        = make(map[int]bool)
)

type Samples struct {
	sample []int
}

type Pull struct {
	fileNames []string
}

func (p *Pull) genPull(obj []Samples, cnt int) {

	f1, err := ioutil.ReadDir(MY_PATH)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(f1); i++ {
		key, _ := strconv.Atoi(strings.Split(f1[i].Name(), "_")[0])
		if contains(obj[cnt].sample, key) {
			p.fileNames = append(p.fileNames, filepath.Join(MY_PATH, f1[i].Name()))
		}
	}
}

func (s *Samples) genSample() {
	for len(s.sample) < LEN {
		temp := rand.Intn((MAX - MIN) + MIN)
		if _, ok := SET[temp]; ok {
			continue
		}
		// if contains(s.sample, temp) {
		// 	continue
		// }
		s.sample = append(s.sample, temp)
		SET[temp] = true
	}
}

func contains(arr []int, ele int) bool {
	for _, v := range arr {
		if v == ele {
			return true
		}
	}
	return false
}

func copyPull(pull []Pull) {
	for _, v1 := range pull {
		for _, v2 := range v1.fileNames {
			dest := filepath.Join("RS-Pull", strings.Split(v2, "/")[1], strings.Split(v2, "/")[2])
			input, _ := ioutil.ReadFile(v2)

			err = ioutil.WriteFile(dest, input, 0644)
			if err != nil {
				fmt.Println("Error creating", dest)
				fmt.Println(err)
				return
			}
		}
	}
}

func main() {
	rand.Seed(SEED)
	if err != nil {
		log.Fatal(err)
	}

	obj := make([]Samples, 0)
	pul := make([]Pull, 0)
	var cnt int

	for _, f := range files {
		if f.IsDir() && f.Name() != ".git" {
			// if cnt < 3 {
			t := Samples{}
			t.genSample()
			obj = append(obj, t)
			// }
			MY_PATH = filepath.Join(ROOT, f.Name())
			p := Pull{}
			p.genPull(obj, cnt)
			pul = append(pul, p)
			cnt++
		}

	}
	fmt.Println(obj)
	copyPull(pul)
}
