// Code generated by protoc-gen-grpc-mock. DO NOT EDIT.
// source: components/automate-gateway/api/compliance/reporting/stats/stats.proto

package stats

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// verify that the mock satisfies the StatsServiceServer interface (at compile time)
var _ StatsServiceServer = &StatsServiceServerMock{}

// NewStatsServiceServerMock gives you a fresh instance of StatsServiceServerMock.
func NewStatsServiceServerMock() *StatsServiceServerMock {
	return &StatsServiceServerMock{validateRequests: true}
}

// NewStatsServiceServerMockWithoutValidation gives you a fresh instance of
// StatsServiceServerMock which does not attempt to validate requests before passing
// them to their respective '*Func'.
func NewStatsServiceServerMockWithoutValidation() *StatsServiceServerMock {
	return &StatsServiceServerMock{}
}

// StatsServiceServerMock is the mock-what-you-want struct that stubs all not-overridden
// methods with "not implemented" returns
type StatsServiceServerMock struct {
	validateRequests bool
	ReadSummaryFunc  func(context.Context, *Query) (*Summary, error)
	ReadTrendFunc    func(context.Context, *Query) (*Trends, error)
	ReadProfilesFunc func(context.Context, *Query) (*Profile, error)
	ReadFailuresFunc func(context.Context, *Query) (*Failures, error)
}

func (m *StatsServiceServerMock) ReadSummary(ctx context.Context, req *Query) (*Summary, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ReadSummaryFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'ReadSummary' not implemented")
}

func (m *StatsServiceServerMock) ReadTrend(ctx context.Context, req *Query) (*Trends, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ReadTrendFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'ReadTrend' not implemented")
}

func (m *StatsServiceServerMock) ReadProfiles(ctx context.Context, req *Query) (*Profile, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ReadProfilesFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'ReadProfiles' not implemented")
}

func (m *StatsServiceServerMock) ReadFailures(ctx context.Context, req *Query) (*Failures, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ReadFailuresFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'ReadFailures' not implemented")
}

// Reset resets all overridden functions
func (m *StatsServiceServerMock) Reset() {
	m.ReadSummaryFunc = nil
	m.ReadTrendFunc = nil
	m.ReadProfilesFunc = nil
	m.ReadFailuresFunc = nil
}
