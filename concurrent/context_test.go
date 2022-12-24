package concurrent

import (
	"context"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	ctx, cf := context.WithTimeout(context.Background(), time.Second*4)
	defer cf()
	ctx.Done()
}