package day

import (
	"fmt"
	"hyoa/aoc2023/internal/utils"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
)

type mapCoordinate struct {
	source      int
	destination int
	lenght      int
}

func init() {
	DayCollection["5"] = Day5
}

type day5 struct {
	input1 []string
}

func Day5(input1, input2 string) (any, any) {
	d := day5{
		input1: utils.ReadTextFileLinesAsString(input1),
	}

	return d.step1(), d.step2()
}

func (d day5) step1() any {
	result := 0

	seeds, seedsToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation := getMappings(d.input1)

	for _, seed := range seeds {
		// to soil
		d := getValueFromMapping(seed, seedsToSoil)

		// to fertilizer
		d = getValueFromMapping(d, soilToFertilizer)

		// to water
		d = getValueFromMapping(d, fertilizerToWater)

		// to light
		d = getValueFromMapping(d, waterToLight)

		// to temperature
		d = getValueFromMapping(d, lightToTemperature)

		// to humidity
		d = getValueFromMapping(d, temperatureToHumidity)

		// to location
		d = getValueFromMapping(d, humidityToLocation)

		if result == 0 {
			result = d
		} else if d < result {
			result = d
		}
	}

	return result
}

func (d day5) step2() any {
	timeStart := time.Now()
	result := 0

	seeds, seedsToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation := getMappings(d.input1)

	var wg sync.WaitGroup

	type lockResult struct {
		mu     sync.Mutex
		result int
	}

	lr := lockResult{
		result: 0,
	}

	for i := 0; i < len(seeds); i += 2 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < seeds[i+1]; j++ {
				seed := seeds[i] + j

				// to soil
				d := getValueFromMapping(seed, seedsToSoil)

				// to fertilizer
				d = getValueFromMapping(d, soilToFertilizer)

				// to water
				d = getValueFromMapping(d, fertilizerToWater)

				// to light
				d = getValueFromMapping(d, waterToLight)

				// to temperature
				d = getValueFromMapping(d, lightToTemperature)

				// to humidity
				d = getValueFromMapping(d, temperatureToHumidity)

				// to location
				d = getValueFromMapping(d, humidityToLocation)

				lr.mu.Lock()
				if lr.result == 0 {
					lr.result = d
				} else if d < lr.result {
					lr.result = d
				}

				lr.mu.Unlock()
			}
		}(i)
	}

	wg.Wait()

	fmt.Println("time", time.Since(timeStart))

	result = lr.result
	return result
}

func getValueFromMapping(source int, mappings []mapCoordinate) int {
	for _, v := range mappings {
		if source >= v.source && source <= v.source+v.lenght {
			return v.destination + (source - v.source)
		}
	}

	return source
}

func getMappings(input []string) ([]int, []mapCoordinate, []mapCoordinate, []mapCoordinate, []mapCoordinate, []mapCoordinate, []mapCoordinate, []mapCoordinate) {
	seedsToSoil := make([]mapCoordinate, 0)
	soilToFertilizer := make([]mapCoordinate, 0)
	fertilizerToWater := make([]mapCoordinate, 0)
	waterToLight := make([]mapCoordinate, 0)
	lightToTemperature := make([]mapCoordinate, 0)
	temperatureToHumidity := make([]mapCoordinate, 0)
	humidityToLocation := make([]mapCoordinate, 0)

	idxToSoil := slices.Index(input, "seed-to-soil map:")
	idxToFertilizer := slices.Index(input, "soil-to-fertilizer map:")
	idxToWater := slices.Index(input, "fertilizer-to-water map:")
	idxToLight := slices.Index(input, "water-to-light map:")
	idxToTemperature := slices.Index(input, "light-to-temperature map:")
	idxToHumidity := slices.Index(input, "temperature-to-humidity map:")
	idxToLocation := slices.Index(input, "humidity-to-location map:")

	toSoilData := input[idxToSoil+1 : idxToFertilizer-1]
	toFertilizerData := input[idxToFertilizer+1 : idxToWater-1]
	toWaterData := input[idxToWater+1 : idxToLight-1]
	toLightData := input[idxToLight+1 : idxToTemperature-1]
	toTemperatureData := input[idxToTemperature+1 : idxToHumidity-1]
	toHumidityData := input[idxToHumidity+1 : idxToLocation-1]
	toLocationData := input[idxToLocation+1:]

	seedsData := strings.Split(input[0], "seeds: ")
	seedsN := strings.Split(seedsData[1], " ")
	seeds := make([]int, 0)

	for _, v := range seedsN {
		vInt, _ := strconv.Atoi(v)
		seeds = append(seeds, vInt)
	}

	for _, v := range toSoilData {
		var destination, source, lenght int
		fmt.Sscanf(v, "%d %d %d", &destination, &source, &lenght)
		coo := mapCoordinate{
			source:      source,
			destination: destination,
			lenght:      lenght,
		}

		seedsToSoil = append(seedsToSoil, coo)
	}

	for _, v := range toFertilizerData {
		var destination, source, lenght int
		fmt.Sscanf(v, "%d %d %d", &destination, &source, &lenght)
		coo := mapCoordinate{
			source:      source,
			destination: destination,
			lenght:      lenght,
		}

		soilToFertilizer = append(soilToFertilizer, coo)
	}

	for _, v := range toWaterData {
		var destination, source, lenght int
		fmt.Sscanf(v, "%d %d %d", &destination, &source, &lenght)
		coo := mapCoordinate{
			source:      source,
			destination: destination,
			lenght:      lenght,
		}

		fertilizerToWater = append(fertilizerToWater, coo)
	}

	for _, v := range toLightData {
		var destination, source, lenght int
		fmt.Sscanf(v, "%d %d %d", &destination, &source, &lenght)
		coo := mapCoordinate{
			source:      source,
			destination: destination,
			lenght:      lenght,
		}

		waterToLight = append(waterToLight, coo)
	}

	for _, v := range toTemperatureData {
		var destination, source, lenght int
		fmt.Sscanf(v, "%d %d %d", &destination, &source, &lenght)
		coo := mapCoordinate{
			source:      source,
			destination: destination,
			lenght:      lenght,
		}

		lightToTemperature = append(lightToTemperature, coo)
	}

	for _, v := range toHumidityData {
		var destination, source, lenght int
		fmt.Sscanf(v, "%d %d %d", &destination, &source, &lenght)
		coo := mapCoordinate{
			source:      source,
			destination: destination,
			lenght:      lenght,
		}

		temperatureToHumidity = append(temperatureToHumidity, coo)
	}

	for _, v := range toLocationData {
		var destination, source, lenght int
		fmt.Sscanf(v, "%d %d %d", &destination, &source, &lenght)
		coo := mapCoordinate{
			source:      source,
			destination: destination,
			lenght:      lenght,
		}

		humidityToLocation = append(humidityToLocation, coo)
	}

	return seeds, seedsToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation
}
