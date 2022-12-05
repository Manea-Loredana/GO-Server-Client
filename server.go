package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net"
	"regexp"
	"strconv"
	"strings"
)

func reverse_int(n int) int {
	new_int := 0
	for n > 0 {
		remainder := n % 10
		new_int *= 10
		new_int += remainder
		n /= 10
	}
	return new_int
}

func findArraySum(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res += reverse_int(arr[i])
	}
	return res
}

func convertArraytoInt(arr []string) []int {

	var t2 = []int{}
	for _, i := range arr {
		j, err2 := strconv.Atoi(i)
		if err2 != nil {
			panic(err2)
		}
		t2 = append(t2, j)
	}
	return t2
}

func isPerfectSquare(no int) bool {
	int_root := int(math.Sqrt(float64(no)))
	return int_root*int_root == no
}

func countPerfectSquare(arr []int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if isPerfectSquare(arr[i]) {
			count += 1
		}

	}
	return count
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func countPrimeDigits(arr []int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if isPrime(arr[i]) {
			for arr[i] > 0 {
				arr[i] = arr[i] / 10
				count++
			}

		}

	}
	return count
}

func doubleFirstDigit(a int) int {
	c := 0
	s1 := strconv.Itoa(a)
	s2 := strconv.Itoa(a / 10)
	s := s2 + s1
	c, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		panic(err)
	}

	return c
}

func sumOfDoubleDigit(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res += doubleFirstDigit(arr[i])
	}
	return res
}

type Data struct {
	Server_motd     string
	Max_connections int
	Port            string
}

var count = 0

var content, err = ioutil.ReadFile("./config.json")

var payload Data

var erar = json.Unmarshal(content, &payload)

func handleConnection(c net.Conn) {

	id_client := strconv.Itoa(count)

	fmt.Println("Clientul " + id_client + " a intrat pe server!")

	if count == payload.Max_connections {
		count--
		fmt.Println("Clientul " + id_client + " a iesit de pe server!")
		c.Close()
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		fmt.Println("Clientul " + id_client + " a facut requestul cu datele: " + temp)

		switch temp {
		case "1":
			for {
				primire := "Server-ul a primit requestul tau. Se proceseaza..." + "\n"
				c.Write([]byte(string(primire)))
				procesare := "Scrie numere 1: " + "\n"
				c.Write([]byte(string(procesare)))
				netData2, err := bufio.NewReader(c).ReadString('\n')
				if err != nil {
					fmt.Println(err)
					return
				}
				temp2 := strings.TrimSpace(string(netData2))
				val := strings.Fields(temp2)
				res := convertArraytoInt(val)
				s := strconv.Itoa(findArraySum(res)) + "\n"
				fmt.Println("Clientul " + id_client + " a primit raspunsul " + string(strconv.Itoa(findArraySum(res))))
				c.Write([]byte(string(s)))
				break
			}
		case "2":
			for {
				primire := "Server-ul a primit requestul tau. Se proceseaza..." + "\n"
				c.Write([]byte(string(primire)))
				procesare := "Scrie datele tale 2: " + "\n"
				c.Write([]byte(string(procesare)))
				netData3, err := bufio.NewReader(c).ReadString('\n')
				if err != nil {
					fmt.Println(err)
					return
				}
				temp3 := strings.TrimSpace(string(netData3))
				re := regexp.MustCompile(`[a-z]+`)
				ceva := re.ReplaceAllLiteralString(temp3, "")
				val2 := strings.Fields(ceva)
				res2 := convertArraytoInt(val2)
				rezultat := countPerfectSquare(res2)
				s2 := strconv.Itoa(rezultat) + "\n"
				c.Write([]byte(string(s2)))
				fmt.Println("Clientul " + id_client + " a primit raspunsul " + strconv.Itoa(rezultat))
				break
			}
		case "3":
			for {
				primire := "Server-ul a primit requestul tau. Se proceseaza..." + "\n"
				c.Write([]byte(string(primire)))
				procesare := "Scrie datele tale 3: " + "\n"
				c.Write([]byte(string(procesare)))
				netData4, err := bufio.NewReader(c).ReadString('\n')
				if err != nil {
					fmt.Println(err)
					return
				}
				temp4 := strings.TrimSpace(string(netData4))
				val2 := strings.Fields(temp4)
				res2 := convertArraytoInt(val2)

				rezultat2 := countPrimeDigits(res2)
				s2 := strconv.Itoa(rezultat2) + "\n"
				fmt.Println("Clientul " + id_client + " a primit raspunsul " + strconv.Itoa(rezultat2))
				c.Write([]byte(string(s2)))
				break
			}
		case "4":
			for {
				primire := "Server-ul a primit requestul tau. Se proceseaza..." + "\n"
				c.Write([]byte(string(primire)))
				procesare := "Scrie datele tale 4: " + "\n"
				c.Write([]byte(string(procesare)))
				netData5, err := bufio.NewReader(c).ReadString('\n')
				if err != nil {
					fmt.Println(err)
					return
				}
				temp5 := strings.TrimSpace(string(netData5))
				val3 := strings.Fields(temp5)
				res3 := convertArraytoInt(val3)
				rezultat2 := sumOfDoubleDigit(res3)
				s2 := strconv.Itoa(rezultat2) + "\n"
				fmt.Println("Clientul " + id_client + " a primit raspunsul " + strconv.Itoa(rezultat2))
				c.Write([]byte(string(s2)))
			}
		default:
			primire := payload.Server_motd + "\n"
			fmt.Println("Clientul " + id_client + " a primit raspunsul default")
			c.Write([]byte(string(primire)))
		}

		if temp == "STOP" {
			count--
			fmt.Println("Clientul " + id_client + " a iesit de pe server!")
			break
		}
	}
	c.Close()
}

func main() {

	PORT := ":" + payload.Port
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
		count++
	}
}
