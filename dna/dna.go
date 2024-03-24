package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[byte]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
// /
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	var h Histogram = make(Histogram)
	for _, c := range [4]byte{'A', 'C', 'G', 'T'} {
		h[c] = 0
	}
	for _, c := range []byte(d) {
		switch c {
		case 'A', 'C', 'G', 'T':
			h[c] += 1
		default:
			return h, errors.New(":(")
		}

	}
	return h, nil
}
