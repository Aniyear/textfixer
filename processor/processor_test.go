package processor

import "testing"

func TestConvertHexAndBin(t *testing.T) {
    input := "Add 1E (hex) and 101 (bin)"
    expected := "Add 30 and 5"
    result := ProcessText(input)

    if result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}

func TestApplyTransformations(t *testing.T) {
    cases := []struct {
        input, expected string
    }{
        {"go (up)", "GO"},
        {"stop SHOUTING (low)", "stop shouting"},
        {"this is amazing (up, 2)", "this is AMAZING AMAZING"},
        {"visit the brooklyn bridge (cap)", "visit the brooklyn Bridge"},
        {"it was the age of foolishness (cap, 4)", "it was the age Of Foolishness"},
    }

    for _, c := range cases {
        res := ProcessText(c.input)
        if res != c.expected {
            t.Errorf("Expected '%s', got '%s'", c.expected, res)
        }
    }
}
