// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// Id defines the a result id or result type for any instruction.
type Id uint32

// Verify returns an error if this is not a valid Id.
//
// TODO implement Id range checking.
func (i Id) Verify() error {
	return nil
}
