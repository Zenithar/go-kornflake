// Code generated by protoc-gen-defaults. DO NOT EDIT.

package bigflakev1

import (
	"context"
	"github.com/bxcodec/faker"
)

// MockBigflakeAPIServer is the mock implementation of the BigflakeAPIServer. Use this to create mock services that
// return random data. Useful in UI Testing.
type MockBigflakeAPIServer struct{}

// Get is mock implementation of the method Get
func (MockBigflakeAPIServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	var res GetResponse
	if err := faker.FakeData(&res); err != nil {
		return nil, err
	}
	return &res, nil
}
