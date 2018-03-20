package main

import (
	"fmt"
	"os/exec"
	"strings"

	"io/ioutil"
	"strconv"
	"sync"
)

func main() {
	command := `curl`
	url:= "http://0.0.0.0:8000/authenticate"

	data, err := ioutil.ReadFile("/home/kutepov/Documents/notes/CustomersSG")
	//data, err := ioutil.ReadFile("/home/kutepov/customers_loc.txt")
	check(err)
	customers := string(data)
	lines := strings.Split(string(customers), "\n")
	ln := len(lines)

	finished := 0
	var wg sync.WaitGroup
	circles := 10

	wg.Add(circles * ln)
	for k := 0; k < circles; k++ {
		fmt.Println("NEW LOOP")
		for i, _ := range lines {
			go func(j int) {
				defer wg.Done()

				url2 := url + lines[j]
				_, err := exec.Command(command, url2).Output()
				var out []byte
				if err != nil {
					out = []byte(" ERROR" + err.Error())
				} /* else {
					out = " SUCCESS"
				}*/
				fmt.Print(strconv.Itoa(j) + ": " + lines[j] + " " + url2 + " " + string(out) + " OK\n")
				finished++
			}(i)
		}
	}
	wg.Wait()
	fmt.Print(finished)

	/*	for i := range lines{
		url2 := url + lines[i]
		fmt.Println(url2)
	}*/

	//var out []byte

	//size := 100

	/*wg := awg.AdvancedWaitGroup{}

	for i, el := range lines {
		wg.Add(func() error {
			fmt.Print("Step" + strconv.Itoa(i) + "\n")
			url2 := url + el
			exec.Command(command, url2)
			finished++
			fmt.Print(el)
			time.Sleep(1 * time.Second)
			return nil
		})

	*/ /*

	   		go func() {
	   //			out2, err := exec.Command(command, url).Output()
	   			exec.Command(command, url)
	   			finished++

	   //			out = out2
	   			if err != nil {
	   				fmt.Println(err)
	   			}
	   		}()
	*/ /*
		}
		wg.Start()*/
	/*	for i:= 0; i<10; i++ {
		time.Sleep(1*time.Second)
		//fmt.Println(finished)
	}*/
	/*	for finished < size {
		fmt.Println(finished)
		time.Sleep(1*time.Second)
	}*/
	//fmt.Println(string(out))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
