package image

import (
	"testing"
)

func TestEbayImagesToJPG(t *testing.T) {
	tests := []struct {
		comment string
		input   []string
		want    []string
		isError bool
	}{
		{
			comment: "Correct test",
			input: []string{"https://i.ebayimg.com/images/g/B8EAAOSwVCJelqmr/s-l1600.jpg",
				"https://i.ebayimg.com/images/g/oR8AAOSwqqZiQfID/s-l1600.jpg",
				"https://i.ebayimg.com/images/g/uYoAAOSwgH5jl~ST/s-l1600.png",
				"https://i.ebayimg.com/images/g/oR8AAOSwqqZiQfID/s-l1600.jpeg",
			},
			want: []string{"https://i.ebayimg.com/images/g/B8EAAOSwVCJelqmr/s-l1600.jpg",
				"https://i.ebayimg.com/images/g/oR8AAOSwqqZiQfID/s-l1600.jpg",
				"https://i.ebayimg.com/images/g/uYoAAOSwgH5jl~ST/s-l1600.jpg",
				"https://i.ebayimg.com/images/g/oR8AAOSwqqZiQfID/s-l1600.jpg",
			},
			isError: false,
		},
		{
			comment: "Test with error in ebay domain",
			input: []string{
				"https://i.ebay.com/images/g/oR8AAOSwqqZiQfID/s-l1600.jpg",
			},
			want:    nil,
			isError: true,
		},
		{
			comment: "Test with error in extention",
			input: []string{
				"https://i.ebay.com/images/g/oR8AAOSwqqZiQfID/s-l1600jpg",
			},
			want:    nil,
			isError: true,
		},
	}
	for _, test := range tests {
		result, err := EbayImagesToJPG(test.input)
		if err != nil && !test.isError {
			t.Errorf("Test %q must be error: %v", test.comment, err)
		}
		if err == nil && test.isError {
			t.Errorf("Test %q must NOT be error", test.comment)
		}
		if len(result) != len(test.want) {
			t.Errorf("Test %q got result with len %d, but want len %d", test.comment, len(result), len(test.want))
		}
		for i := range result {
			if result[i] != test.want[i] {
				t.Errorf("Test %q got url %q, but want url %q", test.comment, result[i], test.want[i])
			}
		}
	}
}
