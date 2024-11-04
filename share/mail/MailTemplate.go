package mail

import (
	"context"
	"fmt"
	"strconv"
	"wellbe-common/share/commondb"
	"wellbe-common/share/commondb/persistence"
	"wellbe-common/share/commonmodel"
	commonconstants "wellbe-common/share/commonsettings/constants"
	"wellbe-common/share/commonsettings/constants/code/cCheckoutTiming"
	"wellbe-common/share/commonsettings/constants/code/cMailTemplate"
	errordef "wellbe-common/share/errordef"
	log "wellbe-common/share/log"
	commonmessages "wellbe-common/share/messages"
)

func GetMailTemplateBookingAccept(languageCd string, bookingNo string, shopName string, address string, bookingDateTime string, amount string, currencyCdIso string, checkoutDueDate string, checkoutTimingCd int) (*commonmodel.CMailTemplate, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	mailTemplate, err := commonPersistence.GetMailTemplate(&ctx, strconv.Itoa(cMailTemplate.BOOKING_ACCEPT), languageCd)
    if err != nil {
        return nil, err
    }
	mailTemplatePayment, err := commonPersistence.GetMailTemplate(&ctx, strconv.Itoa(cMailTemplate.CHECKOUT_ANNOUNCEMENT), languageCd)
    if err != nil {
        return nil, err
    }
	paymentAnnouceStr := ""
	if checkoutTimingCd != cCheckoutTiming.LOCAL_CHECKOUT {
		paymentAnnouceStr = fmt.Sprintf(mailTemplatePayment.Body, amount, currencyCdIso, checkoutDueDate)
	}
	mailTemplate.Subject = fmt.Sprintf(mailTemplate.Subject, bookingNo)
    mailTemplate.Body = fmt.Sprintf(mailTemplate.Body, bookingNo, shopName, address, bookingDateTime, paymentAnnouceStr)
	return mailTemplate, nil
}

func GetMailTemplateCancelRequest(languageCd string, bookingNo string, bookingDateTime string, refundAmount string, paidAmount string, cancelAmount string, currencyCdIso string, shopName string) (*commonmodel.CMailTemplate, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	mailTemplate, err := commonPersistence.GetMailTemplate(&ctx, strconv.Itoa(cMailTemplate.CANCEL_REQUEST), languageCd)
    if err != nil {
        return nil, err
    }
	mailTemplate.Subject = fmt.Sprintf(mailTemplate.Subject, shopName, bookingNo)
    mailTemplate.Body = fmt.Sprintf(mailTemplate.Body, bookingNo, bookingDateTime, shopName, refundAmount, currencyCdIso, paidAmount, currencyCdIso, cancelAmount, currencyCdIso)
	return mailTemplate, nil
}

func GetMailTemplateAutoCancelDuedateExpired(languageCd string, bookingNo string, bookingDateTime string, shopName string) (*commonmodel.CMailTemplate, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	mailTemplate, err := commonPersistence.GetMailTemplate(&ctx, strconv.Itoa(cMailTemplate.AUTO_CANCEL), languageCd)
    if err != nil {
        return nil, err
    }
	mailTemplate.Subject = fmt.Sprintf(mailTemplate.Subject, shopName, bookingNo)
    mailTemplate.Body = fmt.Sprintf(mailTemplate.Body, bookingNo, bookingDateTime, shopName)
	return mailTemplate, nil
}

func GetMailTemplatePayoutNotification(languageCd string, closingDate string) (*commonmodel.CMailTemplate, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	mailTemplate, err := commonPersistence.GetMailTemplate(&ctx, strconv.Itoa(cMailTemplate.PAYOUT_NOTIFY), languageCd)
    if err != nil {
        return nil, err
    }
    mailTemplate.Body = fmt.Sprintf(mailTemplate.Body, closingDate)
	return mailTemplate, nil
}

func GetMailTemplatePaymentFailed(languageCd string, bookingNo string, bookingShopName string, amount string, currencyCdIso string) (*commonmodel.CMailTemplate, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	mailTemplate, err := commonPersistence.GetMailTemplate(&ctx, strconv.Itoa(cMailTemplate.PAYMENT_FAILED), languageCd)
    if err != nil {
        return nil, err
    }
	mailTemplate.Subject = fmt.Sprintf(mailTemplate.Subject, bookingNo)
    mailTemplate.Body = fmt.Sprintf(mailTemplate.Body, bookingNo, bookingShopName, amount, currencyCdIso)
	return mailTemplate, nil
}

func GetMailTemplateRequestBookingAccept(languageCd string, bookingNo string, shopName string, address string, bookingDateTime string, amount string, currencyCdIso string, checkoutDueDate string, checkoutTimingCd int) (*commonmodel.CMailTemplate, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	mailTemplate, err := commonPersistence.GetMailTemplate(&ctx, strconv.Itoa(cMailTemplate.REQUEST_BOOKING_ACCEPT), languageCd)
    if err != nil {
        return nil, err
    }
	mailTemplatePayment, err := commonPersistence.GetMailTemplate(&ctx, strconv.Itoa(cMailTemplate.CHECKOUT_ANNOUNCEMENT_REQUEST), languageCd)
    if err != nil {
        return nil, err
    }
	paymentAnnouceStr := ""
	if checkoutTimingCd != cCheckoutTiming.LOCAL_CHECKOUT {
		paymentAnnouceStr = fmt.Sprintf(mailTemplatePayment.Body, amount, currencyCdIso, checkoutDueDate)
	}
	mailTemplate.Subject = fmt.Sprintf(mailTemplate.Subject, bookingNo)
    mailTemplate.Body = fmt.Sprintf(mailTemplate.Body, bookingNo, shopName, address, bookingDateTime, paymentAnnouceStr)
	return mailTemplate, nil
}

func GetMailTemplateCompleteRegistrationAccount(languageCd string) (*commonmodel.CMailTemplate, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	mailTemplate, err := commonPersistence.GetMailTemplate(&ctx, strconv.Itoa(cMailTemplate.COMPLETE_ACCOUNT_REGISTERATION), languageCd)
    if err != nil {
        return nil, err
    }
	return mailTemplate, nil
}

func GetMailTemplateCompleteWithdrawn(languageCd string) (*commonmodel.CMailTemplate, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	mailTemplate, err := commonPersistence.GetMailTemplate(&ctx, strconv.Itoa(cMailTemplate.WITHDRAWN_COMPLETE), languageCd)
    if err != nil {
        return nil, err
    }
	return mailTemplate, nil
}

func GetMailTemplateCancelRequestByShop(languageCd string, bookingNo string, bookingDateTime string, refundAmount string, paidAmount string, cancelAmount string, currencyCdIso string, cancelReason string, shopName string) (*commonmodel.CMailTemplate, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	mailTemplate, err := commonPersistence.GetMailTemplate(&ctx, strconv.Itoa(cMailTemplate.CANCEL_REQUEST_BY_SHOP), languageCd)
    if err != nil {
        return nil, err
    }
	mailTemplate.Subject = fmt.Sprintf(mailTemplate.Subject, shopName, bookingNo)
    mailTemplate.Body = fmt.Sprintf(mailTemplate.Body, bookingNo, bookingDateTime, shopName, refundAmount, currencyCdIso, paidAmount, currencyCdIso, cancelAmount, currencyCdIso, cancelReason)
	return mailTemplate, nil
}

func GetMailTemplateAutoCancelRequestExpire(languageCd string, bookingNo string, bookingDateTime string, shopName string) (*commonmodel.CMailTemplate, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	mailTemplate, err := commonPersistence.GetMailTemplate(&ctx, strconv.Itoa(cMailTemplate.AUTO_CANCEL_REQUEST_EXPIRE), languageCd)
    if err != nil {
        return nil, err
    }
	mailTemplate.Subject = fmt.Sprintf(mailTemplate.Subject, shopName, bookingNo)
    mailTemplate.Body = fmt.Sprintf(mailTemplate.Body, bookingNo, bookingDateTime, shopName)
	return mailTemplate, nil
}

func GetMailTemplateShopBookingAcceptRequest(languageCd string, bookingNo string, bookingDateTime string, bookingMenu string, numberOfPerson string, bookingAmount string, basicalCurrencyCdIso string) (*commonmodel.CMailTemplate, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	mailTemplate, err := commonPersistence.GetMailTemplate(&ctx, strconv.Itoa(cMailTemplate.SHOP_BOOKING_ACCEPT_REQUEST), languageCd)
    if err != nil {
        return nil, err
    }
	mailTemplate.Subject = fmt.Sprintf(mailTemplate.Subject, bookingNo)
    mailTemplate.Body = fmt.Sprintf(mailTemplate.Body, bookingNo, bookingDateTime, bookingMenu, numberOfPerson, bookingAmount, basicalCurrencyCdIso)
	return mailTemplate, nil
}

func GetMailTemplateShopBookingAccept(languageCd string, bookingNo string, bookingDateTime string, bookingMenu string, numberOfPerson string, bookingAmount string, basicalCurrencyCdIso string) (*commonmodel.CMailTemplate, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	mailTemplate, err := commonPersistence.GetMailTemplate(&ctx, strconv.Itoa(cMailTemplate.SHOP_BOOKING_ACCEPT), languageCd)
    if err != nil {
        return nil, err
    }
	mailTemplate.Subject = fmt.Sprintf(mailTemplate.Subject, bookingNo)
    mailTemplate.Body = fmt.Sprintf(mailTemplate.Body, bookingNo, bookingDateTime, bookingMenu, numberOfPerson, bookingAmount, basicalCurrencyCdIso)
	return mailTemplate, nil
}

func GetMailTemplateShopBookingApproved(languageCd string, bookingId string, bookingNo string, shopName string, address string, bookingDateTime string, amount string, currencyCdIso string, checkoutDueDate string) (*commonmodel.CMailTemplate, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	mailTemplate, err := commonPersistence.GetMailTemplate(&ctx, strconv.Itoa(cMailTemplate.SHOP_BOOKING_APPROVED), languageCd)
    if err != nil {
        return nil, err
    }
	mailTemplate.Subject = fmt.Sprintf(mailTemplate.Subject, bookingNo)
    mailTemplate.Body = fmt.Sprintf(mailTemplate.Body, bookingId, bookingNo, shopName, address, bookingDateTime, amount, currencyCdIso, checkoutDueDate)
	return mailTemplate, nil
}

func GetMailTemplateUserContactToWellbe(languageCd string, name string, address string, typeOfContact string, bookingNumber string, subject string, body string) (*commonmodel.CMailTemplate, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	mailTemplate, err := commonPersistence.GetMailTemplate(&ctx, strconv.Itoa(cMailTemplate.USER_CONTACT_TO_WELLBE), languageCd)
    if err != nil {
        return nil, err
    }

	typeOfContacts, err := commonPersistence.GetCTypeOfContactWithKey(&ctx, typeOfContact, languageCd)
    if err != nil {
        return nil, err
    }

	if len(typeOfContacts) == 0 {
		return nil, &errordef.LogicError{Code: commonconstants.LOGIC_ERROR_CODE_MASTER_DATA_IS_UNSETUP, Msg: fmt.Sprintf(commonmessages.MESSAGE_EN_MASTER_DATA_IS_UNSETUP, "c_type_of_contact", "no data")}
	}

	mailTemplate.Subject = fmt.Sprintf(mailTemplate.Subject, address)
    mailTemplate.Body = fmt.Sprintf(mailTemplate.Body, name, address, bookingNumber, typeOfContacts[0].TypeOfContactName, subject, body)
	return mailTemplate, nil
}

func GetMailTemplateReviewRequest(languageCd string, shop_name string) (*commonmodel.CMailTemplate, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	mailTemplate, err := commonPersistence.GetMailTemplate(&ctx, strconv.Itoa(cMailTemplate.REVIEW_REQUEST), languageCd)
    if err != nil {
        return nil, err
    }

	mailTemplate.Subject = fmt.Sprintf(mailTemplate.Subject, shop_name)
    mailTemplate.Body = fmt.Sprintf(mailTemplate.Body)
	return mailTemplate, nil
}

func GetMailTemplateCheckoutRequestResent(languageCd string, bookingId string, bookingNo string, shopName string, address string, bookingDateTime string, amount string, currencyCdIso string, checkoutDueDate string) (*commonmodel.CMailTemplate, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	mailTemplate, err := commonPersistence.GetMailTemplate(&ctx, strconv.Itoa(cMailTemplate.CHECKOUT_REQUEST_RESENT), languageCd)
    if err != nil {
        return nil, err
    }
	mailTemplate.Subject = fmt.Sprintf(mailTemplate.Subject, bookingNo)
    mailTemplate.Body = fmt.Sprintf(mailTemplate.Body, bookingId, bookingNo, shopName, address, bookingDateTime, amount, currencyCdIso, checkoutDueDate)
	return mailTemplate, nil
}

func GetMailTemplateCancelMailToShop(languageCd string, bookingNo string, bookingDateTime string, bookingMenu string, numberOfPerson string, bookingAmount string, currencyCdIso string, cancelDateTime string) (*commonmodel.CMailTemplate, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	mailTemplate, err := commonPersistence.GetMailTemplate(&ctx, strconv.Itoa(cMailTemplate.BOOKING_CANCEL_TO_SHOP), languageCd)
    if err != nil {
        return nil, err
    }
	mailTemplate.Subject = fmt.Sprintf(mailTemplate.Subject, bookingNo)
    mailTemplate.Body = fmt.Sprintf(mailTemplate.Body, bookingNo, bookingDateTime, bookingMenu, numberOfPerson, bookingAmount, currencyCdIso, cancelDateTime)
	return mailTemplate, nil
}