package persistence

import (
	"context"
	"testing"
	env "wellbe-common/settings/env"
	commondb "wellbe-common/share/commondb"
	model "wellbe-common/share/commonmodel"
	log "wellbe-common/share/log"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	env.EnvLoad("./../../../")
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	defer db.Close()
	commonPersistence := NewCommonPersistence(db)
	ctx := context.Background()
	_, err := commonPersistence.CreateNumbering(&ctx, &model.Numbering{NumberingKey: "dummykey", 
																		InitialValue: 99,
																		CurrentValue: 98,
																		MaxValue: 97,
																		FixLength: 1,
																		CreateDatetime: "dummy-CreateDatetime",
																		CreateFunc: "dummy-CreateFunc",
																		UpdateDatetime: "dummy-UpdateDatetime",
																		UpdateFunc: "dummy-UpdateFunc"})
	numbering, _ := commonPersistence.GetOneNumbering(&ctx, "dummykey")
	commondb.Rollback(&ctx)
	assert.Equal(t, numbering.NumberingKey, "dummykey")
	assert.Equal(t, numbering.InitialValue, int64(99))
	assert.Equal(t, numbering.CurrentValue, int64(98))
	assert.Equal(t, numbering.MaxValue, int64(97))
	assert.Equal(t, numbering.FixLength, 1)
	assert.Equal(t, numbering.CreateDatetime, "dummy-CreateDatetime")
	assert.Equal(t, numbering.CreateFunc, "dummy-CreateFunc")
	assert.Equal(t, numbering.UpdateDatetime, "dummy-UpdateDatetime")
	assert.Equal(t, numbering.UpdateFunc, "dummy-UpdateFunc")
	assert.Nil(t, err)
}

func TestUpdate(t *testing.T) {
	env.EnvLoad("./../../../")
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	defer db.Close()
	commonPersistence := NewCommonPersistence(db)
	ctx := context.Background()
	commonPersistence.CreateNumbering(&ctx, &model.Numbering{NumberingKey: "dummykey", 
																InitialValue: 99,
																CurrentValue: 98,
																MaxValue: 97,
																FixLength: 1,
																CreateDatetime: "dummy-CreateDatetime",
																CreateFunc: "dummy-CreateFunc",
																UpdateDatetime: "dummy-UpdateDatetime",
																UpdateFunc: "dummy-UpdateFunc"})
	_, err := commonPersistence.UpdateNumbering(&ctx, &model.Numbering{NumberingKey: "dummykey", 
																		InitialValue: 89,
																		CurrentValue: 88,
																		MaxValue: 87,
																		FixLength: 2,
																		CreateDatetime: "dummy2-CreateDatetime",
																		CreateFunc: "dummy2-CreateFunc",
																		UpdateDatetime: "dummy2-UpdateDatetime",
																		UpdateFunc: "dummy2-UpdateFunc"})
	numbering, _ := commonPersistence.GetOneNumbering(&ctx, "dummykey")
	commondb.Rollback(&ctx)
	assert.Equal(t, numbering.NumberingKey, "dummykey")
	assert.Equal(t, numbering.InitialValue, int64(89))
	assert.Equal(t, numbering.CurrentValue, int64(88))
	assert.Equal(t, numbering.MaxValue, int64(87))
	assert.Equal(t, numbering.FixLength, 2)
	assert.Equal(t, numbering.CreateDatetime, "dummy2-CreateDatetime")
	assert.Equal(t, numbering.CreateFunc, "dummy2-CreateFunc")
	assert.Equal(t, numbering.UpdateDatetime, "dummy2-UpdateDatetime")
	assert.Equal(t, numbering.UpdateFunc, "dummy2-UpdateFunc")
	assert.Nil(t, err)
}