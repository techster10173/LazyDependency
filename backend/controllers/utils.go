package controllers

import (
	"fmt"
	"html"
	"math/rand"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
)

var sanitizerPolicy *bluemonday.Policy

func Init() {
	sanitizerPolicy = bluemonday.UGCPolicy()
}

func structToDbMap(item interface{}) map[string]interface{} {
	res := map[string]interface{}{}

	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	if reflectValue.IsZero() {
		return res
	}
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("db")
		field := reflectValue.Field(i).Interface()
		if tag != "" && tag != "-" {
			if t, ok := field.(time.Time); ok {
				res[tag] = t.UTC().Format(time.RFC3339)
			} else if v.Field(i).Type.Kind() == reflect.Ptr || v.Field(i).Type.Kind() == reflect.Struct {
				res[tag] = structToDbMap(field)
			} else {
				if reflect.TypeOf(field).Kind() == reflect.String {
					res[tag] = html.UnescapeString(sanitizerPolicy.Sanitize(field.(string)))
				} else {
					res[tag] = field
				}
			}
		}
	}
	return res
}

func ValueExtractor(data interface{}, exists bool) (ret interface{}) {
	if exists {
		return data
	} else {
		return nil
	}
}
func isNilFixed(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func searchForLatestTilda(version string, array []string) string {
	current := version
	currentPatch := strings.Split(version, ".")[2]
	currentMinor := strings.Split(version, ".")[1]
	currentMajor := strings.Split(version, ".")[0]
	currentMaxPatch, _ := strconv.Atoi(currentPatch)
	for _, element := range array {
		newPa := strings.Split(element, ".")[2]
		newPaInt, _ := strconv.Atoi(newPa)
		newMi := strings.Split(element, ".")[1]
		newMa := strings.Split(element, ".")[0]
		if newMi == currentMinor && newMa == currentMajor && newPaInt > currentMaxPatch {
			currentMaxPatch = newPaInt
			current = element
		}
	}
	return current
}

func searchForLatestCarat(version string, array []string) string {
	current := version
	currentPatch := strings.Split(version, ".")[2]
	currentMinor := strings.Split(version, ".")[1]
	currentMajor := strings.Split(version, ".")[0]
	currentMaxPatch, _ := strconv.Atoi(currentPatch)
	currentMaxMinor, _ := strconv.Atoi(currentMinor)
	for _, element := range array {
		newPa := strings.Split(element, ".")[2]
		newPaInt, _ := strconv.Atoi(newPa)
		newMi := strings.Split(element, ".")[1]
		newMiInt, _ := strconv.Atoi(newMi)
		newMa := strings.Split(element, ".")[0]
		if newMa == currentMajor && (newMiInt > currentMaxMinor || (newMiInt == currentMaxMinor && newPaInt > currentMaxPatch)) {
			currentMaxPatch = newPaInt
			currentMaxMinor = newMiInt
			current = element
		}
	}
	return current
}

func newVersion(pkg string, version string) string {
	pkgString := pkg + "@*"
	cmd := exec.Command("npm", "view", pkgString, "version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return ""
	}
	arr2 := strings.Split(string(output)+"\n", "\n")
	array := make([]string, 0)
	for _, element := range arr2 {
		if element != "" {
			array = append(array, strings.Split(element, "'")[1])
		}
	}
	// fmt.Printf("%v\n\n", array)
	final := ""
	command := version
	if command == "" {
		final = array[len(array)-1]
	} else if command[0] == '*' {
		final = array[len(array)-1]
	} else if strings.Contains(command, "-") {
		final = strings.Split(command, " - ")[0]
	} else if command[0] == '~' {
		final = searchForLatestTilda(string(command[1:]), array[:])
	} else if command[0] == '^' {
		final = searchForLatestCarat(string(command[1:]), array[:])
	} else if command[0] == '>' {
		final = array[len(array)-1]
	} else if command[0] == '<' {
		final = array[0]
	} else if strings.Contains(command, "x") {
		if strings.Count(command, "x") == 3 {
			final = array[len(array)-1]
		} else if strings.Count(command, "x") == 2 {
			final = searchForLatestCarat(strings.Replace(command, "x", "0", -1), array[:])
		} else if strings.Count(command, "x") == 1 {
			final = searchForLatestTilda(strings.Replace(command, "x", "0", -1), array[:])
		}
	} else {
		final = command
	}

	return final

}
