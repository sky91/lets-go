package dbx

import (
	"context"
	"database/sql/driver"
	stderr "errors"
	"github.com/pkg/errors"
)

type Tx interface {
	driver.Tx
	comparable
}

type EntClient[TX Tx] interface {
	Tx(context.Context) (TX, error)
	comparable
}

type TxOp[TX driver.Tx] func(context.Context, TX) error

func InTx[TX driver.Tx](ctx context.Context, tx TX, run TxOp[TX]) error {
	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()
	if err := run(ctx, tx); err != nil {
		if e := tx.Rollback(); e != nil {
			return stderr.Join(err, errors.Wrap(e, "tx.Rollback() error"))
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "tx.Commit() error")
	}
	return nil
}

type ContextClient[TX Tx, C EntClient[TX]] struct {
	TxFromContext     func(context.Context) TX
	ClientFromContext func(context.Context) C
}

func (thisV ContextClient[TX, C]) InTx(ctx context.Context, run TxOp[TX]) error {
	if thisV.TxFromContext != nil {
		var noTx TX
		if tx := thisV.TxFromContext(ctx); tx != noTx {
			return run(ctx, tx)
		}
	}
	if thisV.ClientFromContext != nil {
		var noClient C
		if client := thisV.ClientFromContext(ctx); client != noClient {
			tx, err := client.Tx(ctx)
			if err != nil {
				return errors.Wrap(err, "client.Tx() error")
			}
			return InTx(ctx, tx, run)
		}
	}
	return errors.New("no tx or client")
}

var RollbackErr = errors.New("rollback")
