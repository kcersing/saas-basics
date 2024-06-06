package admin

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"saas/pkg/db/ent"
)

// rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func rollback(tx *ent.Tx, err error) error {

	hlog.Error("err")
	hlog.Error(err)
	hlog.Error("err")
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
