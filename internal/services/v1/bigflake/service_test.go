package bigflake_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"go.zenithar.org/kornflake/internal/services/v1/bigflake"
	bigflakev1 "go.zenithar.org/kornflake/pkg/gen/go/identifier/bigflake/v1"
)

func TestGet(t *testing.T) {
	g := NewGomegaWithT(t)

	s := bigflake.New(0)
	ctx := context.Background()

	res, err := s.Get(ctx, &bigflakev1.GetRequest{})
	g.Expect(err).To(BeNil(), "Error should be nil")
	g.Expect(res).ToNot(BeNil(), "Result should not be nil")
	g.Expect(res.Identifier).ToNot(BeEmpty(), "Identifier should not be empty")
}

// -----------------------------------------------------------------------------

func BenchmarkGet(b *testing.B) {
	s := bigflake.New(1)
	ctx := context.Background()
	req := &bigflakev1.GetRequest{}

	for i := 0; i < b.N; i++ {
		_, err := s.Get(ctx, req)
		if err != nil {
			b.Errorf("service returned an error")
		}
	}
}
