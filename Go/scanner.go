package main
import (
	"context"
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
	"bufio"
	"golang.org/x/sync/semaphore"
)

type Scanner struct {
	ip string
	lock *semaphore.Weighted
}

func Ulimit() int64 {
	out, err := exec.Command("ulimit", "-n").Output()
	if err != nil {
		panic(err)
	}

	s := strings.TrimSpace(string(out))
	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		panic(err)
	}
	return i
}

func ScanPort(ip, string, port int, timeout time.Duration) {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files"){
			time.Sleep(timeout)
			ScanPort(ip, port, timeout)
		} else {
			fmt.Println(port, "closed")
		}
		return
	}
	conn.Close()
	fmt.Println(port, "open")
}

func (ps *Scanner) Start(f, l int, timeout time.Duration) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for port := f; port <= l; port++ {
		wg.Add(1)
		ps.lock.Acquire(context.TODO(), 1)

		go func(port int){
			defer ps.lock.Release(1)
			defer wg.Done()
			ScanPort(ps.ip, port, timeout)
		} (port)
	}
}

func main(){
	// user input
	reader := bufio.NewReader(os.Stdin)
	var name string
	fmt.Println("What is your target?")
	name, _ := reader.readString("\n")
	ips, err := net.LookupIP(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v/n", err)
		os. Exit(1)

	}
	for _, ip := range ips {
		fmt.Printf(ip.String())
	}



	ps := &Scanner {
		ip: ip
		lock: semaphore.NewWeighted(Ulimit()),
	}
	ps.Start(1, 65535, 500*time.Miillisecond)
}