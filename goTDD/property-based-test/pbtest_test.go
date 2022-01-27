package main

import (
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic int
	Roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{100, "C"},
	{150, "CL"},
	{500, "D"},
	{755, "DCCLV"},
	{920, "CMXX"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
}

func TestConvertToRoman(t *testing.T) {
	for _, c := range cases {
		t.Run("arabic to roman", func(t *testing.T) {
			got := ConvertToRoman(c.Arabic)
			if got != c.Roman {
				t.Errorf("On: arabic to roman, got %v, want %v", got, c.Roman)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, c := range cases {
		t.Run("roman to arabic", func(t *testing.T) {
			got := ConvertToArabic(c.Roman)
			if got != c.Arabic {
				t.Errorf("On: roman to arabic, got %v, want %v", got, c.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log(arabic)
		roman := ConvertToRoman(int(arabic))
		fromRoman := ConvertToArabic(roman)
		return fromRoman == int(arabic)
	}

	if err := quick.Check(assertion, &quick.Config{MaxCount: 10}); err != nil {
		t.Error("failed checks", err)
	}
}
