// Code generated by fastssz. DO NOT EDIT.
// Hash: b05b1d4bd1edc06c10690d812bb9c15ef308ef91a59c5c38b2a8215f12440e48
// Version: 0.1.3
package testcases

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Vec object
func (v *Vec) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(v)
}

// MarshalSSZTo ssz marshals the Vec object to a target array
func (v *Vec) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Values'
	if size := len(v.Values); size != 6 {
		err = ssz.ErrVectorLengthFn("Vec.Values", size, 6)
		return
	}
	for ii := 0; ii < 6; ii++ {
		dst = ssz.MarshalUint64(dst, v.Values[ii])
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Vec object
func (v *Vec) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 48 {
		return ssz.ErrSize
	}

	// Field (0) 'Values'
	v.Values = ssz.ExtendUint64(v.Values, 6)
	for ii := 0; ii < 6; ii++ {
		v.Values[ii] = ssz.UnmarshallUint64(buf[0:48][ii*8 : (ii+1)*8])
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Vec object
func (v *Vec) SizeSSZ() (size int) {
	size = 48
	return
}

// HashTreeRoot ssz hashes the Vec object
func (v *Vec) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(v)
}

// HashTreeRootWith ssz hashes the Vec object with a hasher
func (v *Vec) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Values'
	{
		if size := len(v.Values); size != 6 {
			err = ssz.ErrVectorLengthFn("Vec.Values", size, 6)
			return
		}
		subIndx := hh.Index()
		for _, i := range v.Values {
			hh.AppendUint64(i)
		}
		hh.Merkleize(subIndx)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Vec object
func (v *Vec) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(v)
}

// MarshalSSZ ssz marshals the Vec2 object
func (v *Vec2) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(v)
}

// MarshalSSZTo ssz marshals the Vec2 object to a target array
func (v *Vec2) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(4)

	// Offset (0) 'Values2'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(v.Values2) * 4

	// Field (0) 'Values2'
	if size := len(v.Values2); size > 100 {
		err = ssz.ErrListTooBigFn("Vec2.Values2", size, 100)
		return
	}
	for ii := 0; ii < len(v.Values2); ii++ {
		dst = ssz.MarshalUint32(dst, v.Values2[ii])
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Vec2 object
func (v *Vec2) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Values2'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 4 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (0) 'Values2'
	{
		buf = tail[o0:]
		num, err := ssz.DivideInt2(len(buf), 4, 100)
		if err != nil {
			return err
		}
		v.Values2 = ssz.ExtendUint32(v.Values2, num)
		for ii := 0; ii < num; ii++ {
			v.Values2[ii] = ssz.UnmarshallUint32(buf[ii*4 : (ii+1)*4])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Vec2 object
func (v *Vec2) SizeSSZ() (size int) {
	size = 4

	// Field (0) 'Values2'
	size += len(v.Values2) * 4

	return
}

// HashTreeRoot ssz hashes the Vec2 object
func (v *Vec2) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(v)
}

// HashTreeRootWith ssz hashes the Vec2 object with a hasher
func (v *Vec2) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Values2'
	{
		if size := len(v.Values2); size > 100 {
			err = ssz.ErrListTooBigFn("Vec2.Values2", size, 100)
			return
		}
		subIndx := hh.Index()
		for _, i := range v.Values2 {
			hh.AppendUint32(i)
		}
		hh.FillUpTo32()
		numItems := uint64(len(v.Values2))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(100, numItems, 4))
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Vec2 object
func (v *Vec2) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(v)
}