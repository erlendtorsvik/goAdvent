package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	reHcl    = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	rePid    = regexp.MustCompile(`^\d{9}$`)
	validEcl = map[string]bool{
		"amb": true, "blu": true, "brn": true, "gry": true,
		"grn": true, "hzl": true, "oth": true,
	}

	validators = map[string]func(string) bool{
		"byr": func(s string) bool { return validInt(1920, 2002, s) },
		"iyr": func(s string) bool { return validInt(2010, 2020, s) },
		"eyr": func(s string) bool { return validInt(2020, 2030, s) },
		"hgt": func(s string) bool { return validHeight(s) },
		"hcl": func(s string) bool { return reHcl.MatchString(s) },
		"ecl": func(s string) bool { return validEcl[s] },
		"pid": func(s string) bool { return rePid.MatchString(s) },
	}
)

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	passports := [][]string{}

	for v := range strings.SplitSeq(strings.TrimSpace(string(f)), "\n\n") {
		passports = append(passports, strings.Fields(v))
	}

	fmt.Println(part1ValidCount(passports))
	fmt.Println(part2ValidCount(passports))
	// alternativ måte med å bruke en validerins map (definert lenger oppe som validators)
	fmt.Println(part2bValidCount(passports))
}

func part1ValidCount(passports [][]string) int {
	validCounter := 0

	for _, v := range passports {
		validPassport := parseLine1(v)

		if validPassport {
			validCounter++
		}
	}

	return validCounter
}

func part2ValidCount(passports [][]string) int {
	validCounter := 0

	for _, v := range passports {
		validPassport := parseLine2(v)

		if validPassport {
			validCounter++
		}
	}

	return validCounter
}

func parseLine1(line []string) bool {
	found := make(map[string]bool)

	for _, f := range line {
		parts := strings.SplitN(f, ":", 2)

		if len(parts) != 2 {
			continue
		}
		key := parts[0]
		found[key] = true

	}

	for v := range validators {
		if !found[v] {
			return false
		}
	}
	return true
}

func parseLine2(line []string) bool {
	foundValid := make(map[string]bool)

	for _, f := range line {
		parts := strings.SplitN(f, ":", 2)

		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		value := parts[1]

		switch key {
		case "byr":
			if validInt(1920, 2002, value) {
				foundValid[key] = true
			}
		case "iyr":
			if validInt(2010, 2020, value) {
				foundValid[key] = true
			}
		case "eyr":
			if validInt(2020, 2030, value) {
				foundValid[key] = true
			}
		case "hgt":
			if validHeight(value) {
				foundValid[key] = true
			}
		case "hcl":
			if strings.HasPrefix(value, "#") {
				match := reHcl.MatchString(value)
				if match {
					foundValid[key] = true
				}
			}
		case "ecl":
			if validEcl[value] {
				foundValid[key] = true
			}
		case "pid":
			if rePid.MatchString(value) {
				foundValid[key] = true
			}
		}

	}

	for key := range validators {
		if !foundValid[key] {
			return false
		}
	}
	return true
}

func validInt(min int, max int, s string) bool {
	number, err := strconv.Atoi(s)

	return err == nil && number >= min && number <= max
}

func validHeight(s string) bool {
	if strings.HasSuffix(s, "cm") {
		parsedString := strings.TrimSuffix(s, "cm")
		return validInt(150, 193, parsedString)
	}
	if strings.HasSuffix(s, "in") {
		parsedString := strings.TrimSuffix(s, "in")
		return validInt(59, 76, parsedString)
	}
	return false
}

func part2bValidCount(passports [][]string) int {
	valid := 0

	for _, line := range passports {
		p := parsePassport(line)
		if validatePassport(p) {
			valid++
		}
	}
	return valid
}

func validatePassport(p map[string]string) bool {
	for key, f := range validators {
		// her returneres ok == true om keyen finnes i mappen
		value, ok := p[key]

		// keyen er ikke i mappen i det hele tatt, returner early uten å validere videre
		if !ok {
			return false
		}

		// kjør funksjonen for validering basert på key
		if !f(value) {
			return false
		}
	}
	return true
}

func parsePassport(fields []string) map[string]string {
	m := make(map[string]string)
	for _, f := range fields {
		parts := strings.SplitN(f, ":", 2)

		if len(parts) != 2 {
			continue
		}

		m[parts[0]] = parts[1]
	}

	return m
}
