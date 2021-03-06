package titan_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCoreLiveOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupCoreTestComponents(t, ctrl)
	ok := cmp.core.Live(context.Background())
	if !ok {
		t.Fail()
	}
}
