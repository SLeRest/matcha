package main

import (
	"fmt"
)

func checkFirstName(FirstName string) (string, bool) {
	if len(FirstName) <= 0 || len(FirstName) >= 255 {
		msg := "Error: Wrong len FirstName: len must be beetwen " +
				"1 and 254 characteres. Given FirstName len: %d"
		return fmt.Sprintf(msg , len(FirstName)), false
	}
	return "", true
}

func checkLastName(LastName string) (string, bool) {
	if len(LastName) == 0 || len(LastName) >= 255 {
		msg := "Error: Wrong len LastName: len must be beetwen " +
				"1 and 254 characteres. Given LastName len: %d"
		return fmt.Sprintf(msg, len(LastName)), false
	}
	return "", true
}

// TODO faire le bonus de la geolocalisation ? FUN BUT HARD
func checkCity(City string) (string, bool) {
	if len(City) == 0 || len(City) >= 255 {
		msg := "Error: Wrong len City: len must be beetwen " +
				"1 and 254 characteres. Given City len: %d"
		return fmt.Sprintf(msg, len(City)), false
	}
	return "", true
}

// TODO faire le bonus de la geolocalisation ? FUN BUT HARD
func checkCountry(Country string) (string, bool) {
	if len(Country) == 0 || len(Country) >= 255 {
		msg := "Error: Wrong len City: len must be beetwen " +
				"1 and 254 characteres. Given Country len: %d"
		return fmt.Sprintf(msg, len(Country)), false
	}
	return "", true
}

func checkAge(Age int) (string, bool) {
	if Age < 18 || Age >= 120 {
		msg := "Error: Wrong Age: Age must be beetwen " +
				"18 and 120 year old. Given Age: %d"
		return fmt.Sprintf(msg, Age), false
	}
	return "", true
}

// TODO faire le bonus de non-binaire  ? PAS FUN MAIS EASY
func checkGender(Gender string) (string, bool) {
	if Gender != "Male" && Gender != "Female" {
		msg := "Error: Gender must be Male or Female: Given Gender %v"
		return fmt.Sprintf(msg, Gender), false
	}
	return "", true
}

// TODO faire le bonus de non-binaire  ? PAS FUN MAIS EASY
func checkSexualOrientation(SexualOrientation string) (string, bool) {
	if (SexualOrientation != "Male" &&
	   SexualOrientation != "Female" &&
	   SexualOrientation != "Both") {
		   msg := "Error: SexualOrientation must be Male or Female or Both: " +
				"Given len SexualOrientation %v"
		   return fmt.Sprintf(msg, len(SexualOrientation)), false
	}
	return "", true
}

// TODO check max datatype postgres
func checkDescription(Description string) (string, bool) {
	return "", true
}

func checkTags(Tags []string) (string, bool) {
	for _, Tag := range Tags {
		if len(Tag) <= 0 || len(Tag) >= 255 {
			return fmt.Sprintf("Error: Wrong Tag: %v", Tag), false
		}
	}
	return "", true
}

// TODO si il existe un delire de pointeur sur fonction dans Go ca serait pas mal
func validateUser(User user) ([]string, bool) {
	var error []string
	if msg, ret := checkFirstName(User.FirstName); ret == false {
		error = append(error, msg)
	}
	if msg, ret := checkLastName(User.LastName); ret == false {
		error = append(error, msg)
	}
	if msg, ret := checkCity(User.City); ret == false {
		error = append(error, msg)
	}
	if msg, ret := checkCountry(User.Country); ret == false {
		error = append(error, msg)
	}
	if msg, ret := checkAge(User.Age); ret == false {
		error = append(error, msg)
	}
	if msg, ret := checkGender(User.Gender); ret == false {
		error = append(error, msg)
	}
	if msg, ret := checkSexualOrientation(User.SexualOrientation); ret == false {
		error = append(error, msg)
	}
	if msg, ret := checkDescription(User.Description); ret == false {
		error = append(error, msg)
	}
	if msg, ret := checkTags(User.Tags); ret == false {
		error = append(error, msg)
	}
	if len(error) > 0 {
		return error, false
	}
	return error, true
}
