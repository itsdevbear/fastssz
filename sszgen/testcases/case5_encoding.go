// Code generated by fastssz. DO NOT EDIT.
// Hash: b89ee134b519595b8276174178242203f6bfb9e0594654a23e1c35e5fb8b5369
// Version: 0.1.3-dev
package testcases

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Case5A object
func (c *Case5A) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(c)
}

// MarshalSSZTo ssz marshals the Case5A object to a target array
func (c *Case5A) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'A'
	if size := len(c.A); size != 2 {
		err = ssz.ErrVectorLengthFn("Case5A.A", size, 2)
		return
	}
	for ii := 0; ii < 2; ii++ {
		if size := len(c.A[ii]); size != 2 {
			err = ssz.ErrBytesLengthFn("Case5A.A[ii]", size, 2)
			return
		}
		dst = append(dst, c.A[ii]...)
	}

	// Field (1) 'B'
	if size := len(c.B); size != 2 {
		err = ssz.ErrVectorLengthFn("Case5A.B", size, 2)
		return
	}
	for ii := 0; ii < 2; ii++ {
		if size := len(c.B[ii]); size != 2 {
			err = ssz.ErrBytesLengthFn("Case5A.B[ii]", size, 2)
			return
		}
		dst = append(dst, c.B[ii]...)
	}

	// Field (2) 'C'
	if size := len(c.C); size != 2 {
		err = ssz.ErrVectorLengthFn("Case5A.C", size, 2)
		return
	}
	for ii := 0; ii < 2; ii++ {
		if size := len(c.C[ii]); size != 2 {
			err = ssz.ErrBytesLengthFn("Case5A.C[ii]", size, 2)
			return
		}
		dst = append(dst, c.C[ii]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Case5A object
func (c *Case5A) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 12 {
		return ssz.ErrSize
	}

	// Field (0) 'A'
	c.A = make([][]byte, 2)
	for ii := 0; ii < 2; ii++ {
		if cap(c.A[ii]) == 0 {
			c.A[ii] = make([]byte, 0, len(buf[0:4][ii*2:(ii+1)*2]))
		}
		c.A[ii] = append(c.A[ii], buf[0:4][ii*2:(ii+1)*2]...)
	}

	// Field (1) 'B'
	c.B = make([]Case5Bytes, 2)
	for ii := 0; ii < 2; ii++ {
		if cap(c.B[ii]) == 0 {
			c.B[ii] = make([]byte, 0, len(buf[4:8][ii*2:(ii+1)*2]))
		}
		c.B[ii] = append(c.B[ii], buf[4:8][ii*2:(ii+1)*2]...)
	}

	// Field (2) 'C'
	c.C = make([][]byte, 2)
	for ii := 0; ii < 2; ii++ {
		if cap(c.C[ii]) == 0 {
			c.C[ii] = make([]byte, 0, len(buf[8:12][ii*2:(ii+1)*2]))
		}
		c.C[ii] = append(c.C[ii], buf[8:12][ii*2:(ii+1)*2]...)
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Case5A object
func (c *Case5A) SizeSSZ() (size int) {
	size = 12
	return
}

// HashTreeRoot ssz hashes the Case5A object
func (c *Case5A) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(c)
}

// HashTreeRootWith ssz hashes the Case5A object with a hasher
func (c *Case5A) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'A'
	{
		if size := len(c.A); size != 2 {
			err = ssz.ErrVectorLengthFn("Case5A.A", size, 2)
			return
		}
		subIndx := hh.Index()
		for _, i := range c.A {
			if len(i) != 2 {
				err = ssz.ErrBytesLength
				return
			}
			hh.PutBytes(i)
		}
		hh.Merkleize(subIndx)
	}

	// Field (1) 'B'
	{
		if size := len(c.B); size != 2 {
			err = ssz.ErrVectorLengthFn("Case5A.B", size, 2)
			return
		}
		subIndx := hh.Index()
		for _, i := range c.B {
			if len(i) != 2 {
				err = ssz.ErrBytesLength
				return
			}
			hh.PutBytes(i)
		}
		hh.Merkleize(subIndx)
	}

	// Field (2) 'C'
	{
		if size := len(c.C); size != 2 {
			err = ssz.ErrVectorLengthFn("Case5A.C", size, 2)
			return
		}
		subIndx := hh.Index()
		for _, i := range c.C {
			if len(i) != 2 {
				err = ssz.ErrBytesLength
				return
			}
			hh.PutBytes(i)
		}
		hh.Merkleize(subIndx)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Case5A object
func (c *Case5A) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(c)
}
