package snowflake_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"go.zenithar.org/kornflake/internal/services/v1/snowflake"
	snowflakev1 "go.zenithar.org/kornflake/pkg/gen/go/identifier/snowflake/v1"
)

func TestGet(t *testing.T) {
	g := NewGomegaWithT(t)

	s := snowflake.New(0)
	ctx := context.Background()

	res, err := s.Get(ctx, &snowflakev1.GetRequest{})
	g.Expect(err).To(BeNil(), "Error should be nil")
	g.Expect(res).ToNot(BeNil(), "Result should not be nil")
	g.Expect(res.Identifier).ToNot(BeEmpty(), "Identifier should not be empty")
}

// -----------------------------------------------------------------------------

func BenchmarkGet(b *testing.B) {
	s := snowflake.New(1)
	ctx := context.Background()
	req := &snowflakev1.GetRequest{}

	for i := 0; i < b.N; i++ {
		_, err := s.Get(ctx, req)
		if err != nil {
			b.Errorf("service returned an error")
		}
	}
}
