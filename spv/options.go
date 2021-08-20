package spv

import "github.com/libsv/go-bc"

// ClientOpts can be implemented to provided functional options for an spv.Client.
type ClientOpts func(*spvclient)

// WithBlockHeaderChain will inject the provided BlockHeaderChain into the spv.Client
func WithBlockHeaderChain(bhc bc.BlockHeaderChain) ClientOpts {
	return func(s *spvclient) {
		s.bhc = bhc
	}
}

// WithTXGetter will inject the provided TXGetter into the spv.Client
func WithTXGetter(txg TXGetter) ClientOpts {
	return func(s *spvclient) {
		s.txg = txg
	}
}

// WithMerkleProofGetter will inject the provided MerkleProofGetter into the spv.Client
func WithMerkleProofGetter(mpg MerkleProofGetter) ClientOpts {
	return func(s *spvclient) {
		s.mpg = mpg
	}
}
