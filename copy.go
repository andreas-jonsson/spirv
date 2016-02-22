// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// Copy copies the given slice as-is.
func Copy(v []uint32) []uint32 {
	if len(v) == 0 {
		// `nil` is not the same as `make([]T, 0)`. So be explicit, lest
		// our unit tests complain when comparing the return value of
		// this function with a nil slice in the comparison target.
		//
		// Ref: http://play.golang.org/p/5vO9v00gl7
		return nil
	}

	new := make([]uint32, len(v))
	copy(new, v)
	return v
}
