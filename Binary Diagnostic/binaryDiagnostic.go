package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := os.Args[1]

	data, _ := os.ReadFile(fileName)
	measurements := strings.Split(string(data), "\n")

	measurements = measurements[:len(measurements)-1]

	gammaRate := ""
	epsilonRate := ""

	for i := 0; i < len(measurements[0]); i++ {
		count0 := 0
		count1 := 0

		for j := 0; j < len(measurements); j++ {
			if measurements[j][i] == '0' {
				count0 = count0 + 1
			} else {
				count1 = count1 + 1
			}
		}

		if count0 > count1 {
			gammaRate = gammaRate + "0"
			epsilonRate = epsilonRate + "1"
		} else {
			gammaRate = gammaRate + "1"
			epsilonRate = epsilonRate + "0"
		}
	}

	gammaVal, _ := strconv.ParseInt(gammaRate, 2, 64)
	epsilonVal, _ := strconv.ParseInt(epsilonRate, 2, 64)

	powerConsumption := gammaVal * epsilonVal

	fmt.Println("Part 1, Power Consumption: ", powerConsumption)

	// Part 2
	oxygenGenerator := measurements
	co2Scrubber := measurements

	// for oxygen generator
	for i := 0; i < len(measurements[0]); i++ {
		if len(oxygenGenerator) == 1 {
			break
		}

		count0 := 0
		count1 := 0

		for j := 0; j < len(oxygenGenerator); j++ {
			if oxygenGenerator[j][i] == '0' {
				count0 = count0 + 1
			} else {
				count1 = count1 + 1
			}
		}

		o2measurements := []string{}

		if count1 >= count0 {
			for l := 0; l < len(oxygenGenerator); l++ {
				if oxygenGenerator[l][i] == '1' {
					o2measurements = append(o2measurements, oxygenGenerator[l])
				}
			}
		} else {
			for l := 0; l < len(oxygenGenerator); l++ {
				if oxygenGenerator[l][i] == '0' {
					o2measurements = append(o2measurements, oxygenGenerator[l])
				}
			}
		}

		oxygenGenerator = o2measurements
	}

	// For Co2 scrubber
	for i := 0; i < len(measurements[0]); i++ {
		if len(co2Scrubber) == 1 {
			break
		}

		count0 := 0
		count1 := 0

		for j := 0; j < len(co2Scrubber); j++ {
			if co2Scrubber[j][i] == '0' {
				count0 = count0 + 1
			} else {
				count1 = count1 + 1
			}
		}

		co2scrubberMeasurements := []string{}

		if count1 < count0 {
			for l := 0; l < len(co2Scrubber); l++ {
				if co2Scrubber[l][i] == '1' {
					co2scrubberMeasurements = append(co2scrubberMeasurements, co2Scrubber[l])
				}
			}
		} else {
			for l := 0; l < len(co2Scrubber); l++ {
				if co2Scrubber[l][i] == '0' {
					co2scrubberMeasurements = append(co2scrubberMeasurements, co2Scrubber[l])
				}
			}
		}

		co2Scrubber = co2scrubberMeasurements
	}

	o2rating, _ := strconv.ParseInt(oxygenGenerator[0], 2, 64)
	co2scrubberRating, _ := strconv.ParseInt(co2Scrubber[0], 2, 64)

	lifeSupportRating := o2rating * co2scrubberRating

	fmt.Println(oxygenGenerator, co2Scrubber, lifeSupportRating)
}
