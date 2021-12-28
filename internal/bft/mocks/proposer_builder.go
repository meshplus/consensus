// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	bft "github.com/meshplus/consensus/internal/bft"
	mock "github.com/stretchr/testify/mock"
)

// ProposerBuilder is an autogenerated mock type for the ProposerBuilder type
type ProposerBuilder struct {
	mock.Mock
}

// NewProposer provides a mock function with given fields: leader, proposalSequence, viewNum, decisionsInView, quorumSize
func (_m *ProposerBuilder) NewProposer(leader uint64, proposalSequence uint64, viewNum uint64, decisionsInView uint64, quorumSize int) bft.Proposer {
	ret := _m.Called(leader, proposalSequence, viewNum, decisionsInView, quorumSize)

	var r0 bft.Proposer
	if rf, ok := ret.Get(0).(func(uint64, uint64, uint64, uint64, int) bft.Proposer); ok {
		r0 = rf(leader, proposalSequence, viewNum, decisionsInView, quorumSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bft.Proposer)
		}
	}

	return r0
}
