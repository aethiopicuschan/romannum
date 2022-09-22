package romannum_test

import (
	"testing"

	"github.com/aethiopicuschan/romannum"
)

func TestToRoman(t *testing.T) {
	// Statue of Liberty
	result, err := romannum.ToRoman(1776)
	if err != nil {
		t.Errorf("%s", err)
	}
	if result != "MDCCLXXVI" {
		t.Errorf("want \"MDCCLXXVI\", got %s", result)
	}
	// Out of range
	_, err = romannum.ToRoman(0)
	if err == nil {
		t.Errorf("want out of range error, got nil")
	}
	_, err = romannum.ToRoman(4000)
	if err == nil {
		t.Errorf("want out of range error, got nil")
	}
}

func TestToInt(t *testing.T) {
	// Statue of Liberty
	result, err := romannum.ToInt("MDCCLXXVI")
	if err != nil {
		t.Errorf("%s", err)
	}
	if result != 1776 {
		t.Errorf("want 1776, got %d", result)
	}
	// Invalid source
	_, err = romannum.ToInt("ABCDEF")
	if err == nil {
		t.Errorf("want invalid source error, got nil")
	}
}
