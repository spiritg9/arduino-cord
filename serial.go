package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/jacobsa/go-serial/serial"
)

func main() {
	// Set up options.
	options := serial.OpenOptions{
		PortName:          "/dev/ttyUSB0",
		BaudRate:          9600,
		DataBits:          8,
		StopBits:          1,
		MinimumReadSize:   1,
		RTSCTSFlowControl: false,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	// Make sure to close it later.
	defer port.Close()

	allon := flag.Bool("allon", false, "turn all relays on")
	alloff := flag.Bool("alloff", false, "turn all relays off")
	status := flag.Bool("status", false, "get status of relays")
	set := flag.Int("set", -1, "which relay to turn on/off")
	toggle := flag.Int("toggle", -1, "which relay to toggle")
	on := flag.Bool("on", false, "use with set, to turn relays off")
	off := flag.Bool("off", false, "use with set, to turn relays off")

	flag.Parse()

	if *allon {
		b := []byte("allon#")
		_, err := port.Write(b)
		if err != nil {
			log.Fatalf("port.Write: %v", err)
		}
	}

	if *alloff {
		b := []byte("alloff#")
		_, err := port.Write(b)
		if err != nil {
			log.Fatalf("port.Write: %v", err)
		}
	}

	if *status {
		b := []byte("status#")
		_, err := port.Write(b)
		if err != nil {
			log.Fatalf("port.Write: %v", err)
		}
		buf := make([]byte, 4)
		_, err = port.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Status: '%s'\n", buf)
	}

	if *toggle >= 1 && *toggle <= 4 {
		messageString := fmt.Sprintf("toggle%d#", *toggle)
		_, err := port.Write([]byte(messageString))
		if err != nil {
			log.Fatalf("port.Write: %v", err)
		}
	}

	if *set >= 1 && *set <= 4 {
		messageString := fmt.Sprintf("set%d", *set)
		if *on && *off {
			fmt.Println("Ambiguous command")
		} else if *off {
			messageString += "off#"
			_, err := port.Write([]byte(messageString))
			if err != nil {
				log.Fatalf("port.Write: %v", err)
			}
		} else if *on {
			messageString += "on#"
			_, err := port.Write([]byte(messageString))
			if err != nil {
				log.Fatalf("port.Write: %v", err)
			}
		} else {
			fmt.Println("on or off?")
		}
	}

	time.Sleep(200 * time.Millisecond)
}
