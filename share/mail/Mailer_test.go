package mail

import (
	"context"
	"testing"
	env "wellbe-common/settings/env"
	"wellbe-common/share/commondb"
	"wellbe-common/share/commondb/persistence"
	"wellbe-common/share/commonmodel"

	"github.com/stretchr/testify/assert"
)

func TestRequestEmailUpdate(t *testing.T) {
    env.EnvLoad("./../../")
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	commonPersistence := persistence.NewCommonPersistence(db)
	mailTemplate, _ := commonPersistence.GetMailTemplate(&ctx, "1", "2")
	err := Send(&commonmodel.Mail{
		ToAddress: "k.sasaki.tky@gmail.com",
		SubjectText: mailTemplate.Subject,
		BodyText: mailTemplate.Body,
	})
    assert.Nil(t, err)
}