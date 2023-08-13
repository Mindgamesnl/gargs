package gargs

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var NotEnoughParams = errors.New("Not enough positional parameters")

var paramsMap map[string]string
var motd = `
Welcome to gargs!
`

func SetMOTD(message string) {
	motd = message
}

func Handle(args []string, opts interface{}, positionalParams ...string) (error, bool) {
	v := reflect.ValueOf(opts)
	t := v.Type()

	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		fmt.Println("Error: opts should be a pointer to a struct")
		os.Exit(1)
	}

	v = v.Elem()
	t = t.Elem()

	set := flag.NewFlagSet("os.Args[0]", flag.ContinueOnError)

	helpTriggered := false
	set.Usage = func() {
		helpTriggered = true

		// loop over motd lines
		for _, line := range strings.Split(motd, "\n") {
			fmt.Println(line)
		}

		fmt.Printf("Usage: %s [options..] %s\n", os.Args[0], strings.Join(positionalParams, " "))
		fmt.Println()
		set.PrintDefaults()
	}

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		short := field.Tag.Get("short")
		desc := field.Tag.Get("desc")
		defaultVal := field.Tag.Get("defaultValue") // <-- Extract default value
		name := strings.ToLower(field.Name)

		switch field.Type.Kind() {
		case reflect.String:
			defStrVal := ""
			if defaultVal != "" {
				defStrVal = defaultVal
			}
			defaultVal := field.Tag.Get("defaultValue")
			set.StringVar(v.Field(i).Addr().Interface().(*string), name, defaultVal, desc)
			if short != "" {
				set.StringVar(v.Field(i).Addr().Interface().(*string), short, defStrVal, fmt.Sprintf("%s (shorthand)", desc))
			}
		case reflect.Bool:
			defBoolVal := false
			if defaultVal != "" {
				defBool, err := strconv.ParseBool(defaultVal)
				if err == nil {
					defBoolVal = defBool
				}
			}
			set.BoolVar(v.Field(i).Addr().Interface().(*bool), name, defBoolVal, desc)
			if short != "" {
				set.BoolVar(v.Field(i).Addr().Interface().(*bool), short, defBoolVal, fmt.Sprintf("%s (shorthand)", desc))
			}
		case reflect.Int:
			defIntVal := 0
			if defaultVal != "" {
				defInt, err := strconv.Atoi(defaultVal)
				if err == nil {
					defIntVal = defInt
				}
			}
			defaultIntVal, _ := strconv.Atoi(field.Tag.Get("defaultValue"))
			set.IntVar(v.Field(i).Addr().Interface().(*int), name, defaultIntVal, desc)

			if short != "" {
				set.IntVar(v.Field(i).Addr().Interface().(*int), short, defIntVal, fmt.Sprintf("%s (shorthand)", desc))
			}
		}
	}

	if err := set.Parse(args[1:]); err != nil {
		// check if the error is an expected one
		if errors.Is(err, flag.ErrHelp) {
			return nil, true
		} else {
			set.Usage()
			os.Exit(1)
		}
	}

	// Parse positional parameters
	unparsedArgs := set.Args()
	if len(unparsedArgs) < countRequiredParams(positionalParams) {
		fmt.Printf("Not enough positional parameters. Expected format: %s\n", strings.Join(positionalParams, " "))
		set.Usage()
		return NotEnoughParams, false
	}

	paramsList := unparsedArgs[:min(len(unparsedArgs), len(positionalParams))]
	paramsMap = make(map[string]string)

	for i, param := range positionalParams {
		if i < len(paramsList) {
			paramName := strings.Trim(param, "<>[]")
			paramsMap[paramName] = paramsList[i]
		}
	}

	return nil, helpTriggered
}

func Params() []string {
	params := make([]string, len(paramsMap))
	i := 0
	for _, param := range paramsMap {
		params[i] = param
		i++
	}
	return params
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GetParam(name string) string {
	if val, ok := paramsMap[name]; ok {
		return val
	}
	return ""
}

func countRequiredParams(params []string) int {
	count := 0
	for _, param := range params {
		if strings.HasPrefix(param, "<") && strings.HasSuffix(param, ">") {
			count++
		}
	}
	return count
}
