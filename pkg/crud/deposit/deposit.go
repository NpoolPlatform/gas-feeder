package deposit

import (
	"context"
	"fmt"
	"time"

<<<<<<< HEAD
	constant "github.com/NpoolPlatform/gas-feeder/pkg/const"
=======
>>>>>>> api done
	"github.com/NpoolPlatform/gas-feeder/pkg/db"
	"github.com/NpoolPlatform/gas-feeder/pkg/db/ent"
	"github.com/NpoolPlatform/gas-feeder/pkg/db/ent/deposit"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/gasfeeder"

	"github.com/google/uuid"
)

type Deposit struct {
	*db.Entity
}

func New(ctx context.Context, tx *ent.Tx) (*Deposit, error) {
	e, err := db.NewEntity(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("fail create entity: %v", err)
	}

	return &Deposit{
		Entity: e,
	}, nil
}

func (s *Deposit) rowToObject(row *ent.Deposit) *npool.Deposit {
	return &npool.Deposit{
		ID:            row.ID.String(),
		AccountID:     row.AccountID.String(),
		DepositAmount: row.DepositAmount,
	}
}

func (s *Deposit) Create(ctx context.Context, in *npool.Deposit) (*npool.Deposit, error) {
	var info *ent.Deposit
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		info, err = s.Tx.Deposit.Create().
			SetAccountID(uuid.MustParse(in.AccountID)).
			SetDepositAmount(in.DepositAmount).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("fail create Deposit: %v", err)
	}

	return s.rowToObject(info), nil
}

func (s *Deposit) CreateBulk(ctx context.Context, in []*npool.Deposit) ([]*npool.Deposit, error) {
	rows := []*ent.Deposit{}
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		bulk := make([]*ent.DepositCreate, len(in))
		for i, info := range in {
			bulk[i] = s.Tx.Deposit.Create().
				SetAccountID(uuid.MustParse(info.AccountID)).
				SetDepositAmount(info.DepositAmount)
		}
		rows, err = s.Tx.Deposit.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("fail create Deposits: %v", err)
	}

	infos := []*npool.Deposit{}
	for _, row := range rows {
		infos = append(infos, s.rowToObject(row))
	}

	return infos, nil
}

func (s *Deposit) Row(ctx context.Context, id uuid.UUID) (*npool.Deposit, error) {
	var info *ent.Deposit
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		info, err = s.Tx.Deposit.Query().Where(deposit.ID(id)).Only(_ctx)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("fail get Deposit: %v", err)
	}

	return s.rowToObject(info), nil
}

func (s *Deposit) Update(ctx context.Context, in *npool.Deposit) (*npool.Deposit, error) {
	var info *ent.Deposit
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		info, err = s.Tx.Deposit.UpdateOneID(uuid.MustParse(in.GetID())).
			SetAccountID(uuid.MustParse(in.GetAccountID())).
			SetDepositAmount(in.GetDepositAmount()).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("fail update Deposit: %v", err)
	}

	return s.rowToObject(info), nil
}

func (s *Deposit) Rows(ctx context.Context, conds cruder.Conds, offset, limit int) ([]*npool.Deposit, int, error) {
	rows := []*ent.Deposit{}
	var total int

	err := db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		stm, err := s.queryFromConds(conds)
		if err != nil {
			return fmt.Errorf("fail construct stm: %v", err)
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return fmt.Errorf("fail count Deposit: %v", err)
		}

		rows, err = stm.Order(ent.Desc(deposit.FieldUpdatedAt)).Offset(offset).Limit(limit).All(_ctx)
		if err != nil {
			return fmt.Errorf("fail query Deposit: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get Deposit: %v", err)
	}

	infos := []*npool.Deposit{}
	for _, row := range rows {
		infos = append(infos, s.rowToObject(row))
	}

	return infos, total, nil
}

func (s *Deposit) queryFromConds(conds cruder.Conds) (*ent.DepositQuery, error) {
	stm := s.Tx.Deposit.Query()
	for k, v := range conds {
		switch k {
<<<<<<< HEAD
		case constant.FieldID:
=======
		case deposit.FieldID:
>>>>>>> api done
			id, err := cruder.AnyTypeUUID(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid ID: %v", err)
			}
			stm = stm.Where(deposit.ID(id))
<<<<<<< HEAD
		case constant.FieldAccountID:
=======
		case deposit.FieldAccountID:
>>>>>>> api done
			id, err := cruder.AnyTypeUUID(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid AccountID: %v", err)
			}
			stm = stm.Where(deposit.AccountID(id))
<<<<<<< HEAD
		case constant.FieldDepositAmount:
=======
		case deposit.FieldDepositAmount:
>>>>>>> api done
			depositThreshold, err := cruder.AnyTypeUint64(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid DepositAmount: %v", err)
			}
			switch v.Op {
			case cruder.EQ:
				stm = stm.Where(deposit.DepositAmountEQ(depositThreshold))
			case cruder.GT:
				stm = stm.Where(deposit.DepositAmountGT(depositThreshold))
			case cruder.LT:
				stm = stm.Where(deposit.DepositAmountLT(depositThreshold))
			}
		default:
			return nil, fmt.Errorf("invalid Deposit field")
		}
	}

	return stm, nil
}

func (s *Deposit) RowOnly(ctx context.Context, conds cruder.Conds) (*npool.Deposit, error) {
	var info *ent.Deposit

	err := db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		stm, err := s.queryFromConds(conds)
		if err != nil {
			return fmt.Errorf("fail construct stm: %v", err)
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			return fmt.Errorf("fail query Deposit: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get Deposit: %v", err)
	}

	return s.rowToObject(info), nil
}

func (s *Deposit) Count(ctx context.Context, conds cruder.Conds) (uint32, error) {
	var err error
	var total int

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		stm, err := s.queryFromConds(conds)
		if err != nil {
			return fmt.Errorf("fail construct stm: %v", err)
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return fmt.Errorf("fail check Deposits: %v", err)
		}

		return nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count Deposits: %v", err)
	}

	return uint32(total), nil
}

func (s *Deposit) Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	var err error
	exist := false

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		exist, err = s.Tx.Deposit.Query().Where(deposit.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, fmt.Errorf("fail check Deposit: %v", err)
	}

	return exist, nil
}

func (s *Deposit) ExistConds(ctx context.Context, conds cruder.Conds) (bool, error) {
	var err error
	exist := false

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		stm, err := s.queryFromConds(conds)
		if err != nil {
			return fmt.Errorf("fail construct stm: %v", err)
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return fmt.Errorf("fail check congases: %v", err)
		}

		return nil
	})
	if err != nil {
		return false, fmt.Errorf("fail check congases: %v", err)
	}

	return exist, nil
}

func (s *Deposit) Delete(ctx context.Context, id uuid.UUID) (*npool.Deposit, error) {
	var info *ent.Deposit
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		info, err = s.Tx.Deposit.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete Deposit: %v", err)
	}

	return s.rowToObject(info), nil
}