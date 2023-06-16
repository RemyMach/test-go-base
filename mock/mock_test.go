package mock

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type Datastore interface {
	Store(data string) error
}

func ProcessData(ds Datastore, data string) error {
	return ds.Store(data)
}

type DatastoreMock struct {
	mock.Mock
}

func (m *DatastoreMock) Store(data string) error {
	args := m.Called(data)
	return args.Error(0)
}

func TestProcessData(t *testing.T) {
	datastoreMock := new(DatastoreMock)

	datastoreMock.On("Store", "test data").Return(nil)

	err := ProcessData(datastoreMock, "test data")

	require.NoError(t, err)

	datastoreMock.AssertCalled(t, "Store", "test data")
}
