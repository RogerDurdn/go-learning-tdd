package main

import (
	"reflect"
	"testing"
)

type Case struct {
	name string
	input interface{}
	expectedCall []string
}
type Person struct {
	Name string
	Profile Profile
}
type Profile struct {
	Age int
	City string
}


func TestWalk(t *testing.T) {
	cases := make([]Case,0)

	cases = append(cases,
		createCase("getting expected",
		struct{Name string}{ "roger"},
		[]string{"roger"}))
	cases = append(cases, 
		createCase("two words expected",
		struct{ Name string
		Country string}{"juan", "mx"},
		[]string{"juan", "mx"}))
	cases = append(cases, 
		createCase("one numeric value",
		struct{ Name string
			Age int}{"older", 33},
		[]string{"older"}))
	cases = append(cases, 
		createCase("nested fields",
		Person{"roger",
			Profile{22, "MX"},
		},
		[]string{"roger", "MX"}))
	cases = append(cases, 
		createCase("pointers to things",
		&Person{"pepe",
			Profile{22, "EDO"},
		},
		[]string{"pepe", "EDO"}))
	cases = append(cases, 
		createCase("slices",
			[]Profile{
				{22, "EDO"},
				{10, "MX"},
			},
		[]string{"EDO", "MX"}))
	cases = append(cases, 
		createCase("arrays",
		[2]Profile{
			{22, "JP"},
			{10, "USA"},
		},
		[]string{"JP", "USA"}))

	for _, c := range cases {
		var got []string
		t.Run(c.name, func(t *testing.T) {
			walk(c.input, func(input string){
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, c.expectedCall) {
				t.Errorf("expected %v, got %v, on:%s", c.expectedCall, got, c.name)
			}
		})
	}
	caseMap := createCase("maps",
		map[string]string{
			"some": "one",
			"other": "two",
		},
		[]string{"one", "two"})
	t.Run(caseMap.name, func(t *testing.T) {
		var got []string
		walk(caseMap.input, func(s string) {
			got = append(got, s)
		})
		containsVal(t, got, "one")
		containsVal(t, got, "two")
	})
	t.Run("Channels", func(t *testing.T) {
		aChannel := make(chan Profile)
		go func() {
			aChannel <- Profile{10, "Berlin"}
			aChannel <- Profile{20, "NY"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "NY"}
		walk(aChannel, func(input string) {
			got = append(got,input)
		})
		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v, but want %v", got, want)
		}
	})
	t.Run("functions", func(t *testing.T) {
		aFunction := func()(Profile, Profile){
			return Profile{20,"KO"}, Profile{10,"CAN"}
		}
		var got []string
		want := []string{"KO", "CAN"}
		walk(aFunction, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v, but want %v", got, want)
		}
	})
}

func containsVal(t testing.TB, values []string, wanted string){
	t.Helper()
	present := false
	for _, value := range values {
		if value == wanted{
			present = true
		}
	}
	if !present {
		t.Errorf("Got %v, wanted nested %v",values, wanted)
	}
}

func createCase(name string, input interface{}, expected []string ) Case {
	return Case{name: name,
		input: input,
		expectedCall: expected,
		} 
}