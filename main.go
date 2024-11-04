package main

import (
	"os"
	"time"
	"wellbe-common/adapter/apiclient"
	repository "wellbe-common/adapter/repository"
	persistence "wellbe-common/adapter/repository/persistence"
	"wellbe-common/adapter/repository/query"
	webapi "wellbe-common/adapter/webapi"
	application "wellbe-common/application"
	service "wellbe-common/domain/service"
	migrate "wellbe-common/migrate"
	settings "wellbe-common/settings"
	env "wellbe-common/settings/env"
	constants "wellbe-common/share/commonsettings/constants"
	number "wellbe-common/share/number"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){
	err := env.EnvLoad("./")
	if err != nil {
		return
	}
	db, err := repository.DbOpen()
	if err != nil {
		return
	}

	if len(os.Args) > 1 && os.Args[1] == constants.ARGS_KEYWORD_MIGRATE {
		migrate.Migrate(db, "")
		return
	} else if len(os.Args) > 1 && os.Args[1] == constants.ARGS_KEYWORD_MIGRATE_DEV {
		migrate.Migrate(db, "dev")
		return
	}
	tr := repository.NewTransaction(db)
  
	r := gin.New()
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/"},
	}))
	r.Use(webapi.RecordUaAndTime,
		cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"*",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"PUT",
			"GET",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			constants.API_KEY_REUQEST_HEADER_NAME,
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	  }))
	numberUtil := number.NewNumberUtil()
	persistence := persistence.NewPersistence(tr)
  api := apiclient.NewApiclient()
	healthCheckWebApi := webapi.NewHealthCheckWebApi()
	r = healthCheckWebApi.CreateAccessPoint(r)

    currencyExchangeRateService := service.NewCurrencyExchangeRateService(persistence, numberUtil)
    currencyExchangeRateApplication := application.NewCurrencyExchangeRateApplication(currencyExchangeRateService, tr)
    currencyExchangeRateWebApi := webapi.NewCurrencyExchangeRateWebApi(currencyExchangeRateApplication)
    r = currencyExchangeRateWebApi.CreateAccessPoint(r)

    currencyForPaymentService := service.NewCurrencyForPaymentService(persistence, numberUtil)
    currencyForPaymentApplication := application.NewCurrencyForPaymentApplication(currencyForPaymentService, tr)
    currencyForPaymentWebApi := webapi.NewCurrencyForPaymentWebApi(currencyForPaymentApplication)
    r = currencyForPaymentWebApi.CreateAccessPoint(r)

    cAccountWithdrawalReasonService := service.NewCAccountWithdrawalReasonService(persistence, numberUtil)
    cAccountWithdrawalReasonApplication := application.NewCAccountWithdrawalReasonApplication(cAccountWithdrawalReasonService, tr)
    cAccountWithdrawalReasonWebApi := webapi.NewCAccountWithdrawalReasonWebApi(cAccountWithdrawalReasonApplication)
    r = cAccountWithdrawalReasonWebApi.CreateAccessPoint(r)

    cAreaService := service.NewCAreaService(persistence, numberUtil)
    cAreaApplication := application.NewCAreaApplication(cAreaService, tr)
    cAreaWebApi := webapi.NewCAreaWebApi(cAreaApplication)
    r = cAreaWebApi.CreateAccessPoint(r)

    cAgeRangeService := service.NewCAgeRangeService(persistence, numberUtil)
    cAgeRangeApplication := application.NewCAgeRangeApplication(cAgeRangeService, tr)
    cAgeRangeWebApi := webapi.NewCAgeRangeWebApi(cAgeRangeApplication)
    r = cAgeRangeWebApi.CreateAccessPoint(r)

    cBillingContentService := service.NewCBillingContentService(persistence, numberUtil)
    cBillingContentApplication := application.NewCBillingContentApplication(cBillingContentService, tr)
    cBillingContentWebApi := webapi.NewCBillingContentWebApi(cBillingContentApplication)
    r = cBillingContentWebApi.CreateAccessPoint(r)

    cBillingStatusService := service.NewCBillingStatusService(persistence, numberUtil)
    cBillingStatusApplication := application.NewCBillingStatusApplication(cBillingStatusService, tr)
    cBillingStatusWebApi := webapi.NewCBillingStatusWebApi(cBillingStatusApplication)
    r = cBillingStatusWebApi.CreateAccessPoint(r)

    cBookingChanelService := service.NewCBookingChanelService(persistence, numberUtil)
    cBookingChanelApplication := application.NewCBookingChanelApplication(cBookingChanelService, tr)
    cBookingChanelWebApi := webapi.NewCBookingChanelWebApi(cBookingChanelApplication)
    r = cBookingChanelWebApi.CreateAccessPoint(r)

    cBookingMethodService := service.NewCBookingMethodService(persistence, numberUtil)
    cBookingMethodApplication := application.NewCBookingMethodApplication(cBookingMethodService, tr)
    cBookingMethodWebApi := webapi.NewCBookingMethodWebApi(cBookingMethodApplication)
    r = cBookingMethodWebApi.CreateAccessPoint(r)

    cBookingStatusService := service.NewCBookingStatusService(persistence, numberUtil)
    cBookingStatusApplication := application.NewCBookingStatusApplication(cBookingStatusService, tr)
    cBookingStatusWebApi := webapi.NewCBookingStatusWebApi(cBookingStatusApplication)
    r = cBookingStatusWebApi.CreateAccessPoint(r)

    cCheckoutMethodService := service.NewCCheckoutMethodService(persistence, numberUtil)
    cCheckoutMethodApplication := application.NewCCheckoutMethodApplication(cCheckoutMethodService, tr)
    cCheckoutMethodWebApi := webapi.NewCCheckoutMethodWebApi(cCheckoutMethodApplication)
    r = cCheckoutMethodWebApi.CreateAccessPoint(r)

    cCheckoutTimingService := service.NewCCheckoutTimingService(persistence, numberUtil)
    cCheckoutTimingApplication := application.NewCCheckoutTimingApplication(cCheckoutTimingService, tr)
    cCheckoutTimingWebApi := webapi.NewCCheckoutTimingWebApi(cCheckoutTimingApplication)
    r = cCheckoutTimingWebApi.CreateAccessPoint(r)

    cCheckoutStatusService := service.NewCCheckoutStatusService(persistence, numberUtil)
    cCheckoutStatusApplication := application.NewCCheckoutStatusApplication(cCheckoutStatusService, tr)
    cCheckoutStatusWebApi := webapi.NewCCheckoutStatusWebApi(cCheckoutStatusApplication)
    r = cCheckoutStatusWebApi.CreateAccessPoint(r)

    cContentsLabelService := service.NewCContentsLabelService(persistence, numberUtil)
    cContentsLabelApplication := application.NewCContentsLabelApplication(cContentsLabelService, tr)
    cContentsLabelWebApi := webapi.NewCContentsLabelWebApi(cContentsLabelApplication)
    r = cContentsLabelWebApi.CreateAccessPoint(r)

    cContentsCategoryService := service.NewCContentsCategoryService(persistence, numberUtil)
    cContentsCategoryApplication := application.NewCContentsCategoryApplication(cContentsCategoryService, tr)
    cContentsCategoryWebApi := webapi.NewCContentsCategoryWebApi(cContentsCategoryApplication)
    r = cContentsCategoryWebApi.CreateAccessPoint(r)

    cContentsStatusService := service.NewCContentsStatusService(persistence, numberUtil)
    cContentsStatusApplication := application.NewCContentsStatusApplication(cContentsStatusService, tr)
    cContentsStatusWebApi := webapi.NewCContentsStatusWebApi(cContentsStatusApplication)
    r = cContentsStatusWebApi.CreateAccessPoint(r)

    cContactStatusService := service.NewCContactStatusService(persistence, numberUtil)
    cContactStatusApplication := application.NewCContactStatusApplication(cContactStatusService, tr)
    cContactStatusWebApi := webapi.NewCContactStatusWebApi(cContactStatusApplication)
    r = cContactStatusWebApi.CreateAccessPoint(r)

    cShopContractPlanItemService := service.NewCShopContractPlanItemService(persistence, numberUtil)
    cShopContractPlanItemApplication := application.NewCShopContractPlanItemApplication(cShopContractPlanItemService, tr)
    cShopContractPlanItemWebApi := webapi.NewCShopContractPlanItemWebApi(cShopContractPlanItemApplication)
    r = cShopContractPlanItemWebApi.CreateAccessPoint(r)

    cCountryService := service.NewCCountryService(persistence, numberUtil)
    cCountryApplication := application.NewCCountryApplication(cCountryService, tr)
    cCountryWebApi := webapi.NewCCountryWebApi(cCountryApplication)
    r = cCountryWebApi.CreateAccessPoint(r)

    cCouponTargetAttrService := service.NewCCouponTargetAttrService(persistence, numberUtil)
    cCouponTargetAttrApplication := application.NewCCouponTargetAttrApplication(cCouponTargetAttrService, tr)
    cCouponTargetAttrWebApi := webapi.NewCCouponTargetAttrWebApi(cCouponTargetAttrApplication)
    r = cCouponTargetAttrWebApi.CreateAccessPoint(r)

    cCurrencyService := service.NewCCurrencyService(persistence, numberUtil)
    cCurrencyApplication := application.NewCCurrencyApplication(cCurrencyService, tr)
    cCurrencyWebApi := webapi.NewCCurrencyWebApi(cCurrencyApplication)
    r = cCurrencyWebApi.CreateAccessPoint(r)

    cGenderService := service.NewCGenderService(persistence, numberUtil)
    cGenderApplication := application.NewCGenderApplication(cGenderService, tr)
    cGenderWebApi := webapi.NewCGenderWebApi(cGenderApplication)
    r = cGenderWebApi.CreateAccessPoint(r)

    cInvoiceStatusService := service.NewCInvoiceStatusService(persistence, numberUtil)
    cInvoiceStatusApplication := application.NewCInvoiceStatusApplication(cInvoiceStatusService, tr)
    cInvoiceStatusWebApi := webapi.NewCInvoiceStatusWebApi(cInvoiceStatusApplication)
    r = cInvoiceStatusWebApi.CreateAccessPoint(r)

    cLanguageService := service.NewCLanguageService(persistence, numberUtil)
    cLanguageApplication := application.NewCLanguageApplication(cLanguageService, tr)
    cLanguageWebApi := webapi.NewCLanguageWebApi(cLanguageApplication)
    r = cLanguageWebApi.CreateAccessPoint(r)

    cMailTemplateService := service.NewCMailTemplateService(persistence, numberUtil)
    cMailTemplateApplication := application.NewCMailTemplateApplication(cMailTemplateService, tr)
    cMailTemplateWebApi := webapi.NewCMailTemplateWebApi(cMailTemplateApplication)
    r = cMailTemplateWebApi.CreateAccessPoint(r)

    cMenuLabelService := service.NewCMenuLabelService(persistence, numberUtil)
    cMenuLabelApplication := application.NewCMenuLabelApplication(cMenuLabelService, tr)
    cMenuLabelWebApi := webapi.NewCMenuLabelWebApi(cMenuLabelApplication)
    r = cMenuLabelWebApi.CreateAccessPoint(r)

    cOrderTypeService := service.NewCOrderTypeService(persistence, numberUtil)
    cOrderTypeApplication := application.NewCOrderTypeApplication(cOrderTypeService, tr)
    cOrderTypeWebApi := webapi.NewCOrderTypeWebApi(cOrderTypeApplication)
    r = cOrderTypeWebApi.CreateAccessPoint(r)

    cPayoutItemCategoryService := service.NewCPayoutItemCategoryService(persistence, numberUtil)
    cPayoutItemCategoryApplication := application.NewCPayoutItemCategoryApplication(cPayoutItemCategoryService, tr)
    cPayoutItemCategoryWebApi := webapi.NewCPayoutItemCategoryWebApi(cPayoutItemCategoryApplication)
    r = cPayoutItemCategoryWebApi.CreateAccessPoint(r)

    cPayoutMethodService := service.NewCPayoutMethodService(persistence, numberUtil)
    cPayoutMethodApplication := application.NewCPayoutMethodApplication(cPayoutMethodService, tr)
    cPayoutMethodWebApi := webapi.NewCPayoutMethodWebApi(cPayoutMethodApplication)
    r = cPayoutMethodWebApi.CreateAccessPoint(r)

    cPayoutStatusService := service.NewCPayoutStatusService(persistence, numberUtil)
    cPayoutStatusApplication := application.NewCPayoutStatusApplication(cPayoutStatusService, tr)
    cPayoutStatusWebApi := webapi.NewCPayoutStatusWebApi(cPayoutStatusApplication)
    r = cPayoutStatusWebApi.CreateAccessPoint(r)

    cRecommendLabelService := service.NewCRecommendLabelService(persistence, numberUtil)
    cRecommendLabelApplication := application.NewCRecommendLabelApplication(cRecommendLabelService, tr)
    cRecommendLabelWebApi := webapi.NewCRecommendLabelWebApi(cRecommendLabelApplication)
    r = cRecommendLabelWebApi.CreateAccessPoint(r)

    cReviewCategoryService := service.NewCReviewCategoryService(persistence, numberUtil)
    cReviewCategoryApplication := application.NewCReviewCategoryApplication(cReviewCategoryService, tr)
    cReviewCategoryWebApi := webapi.NewCReviewCategoryWebApi(cReviewCategoryApplication)
    r = cReviewCategoryWebApi.CreateAccessPoint(r)

    cReviewContentService := service.NewCReviewContentService(persistence, numberUtil)
    cReviewContentApplication := application.NewCReviewContentApplication(cReviewContentService, tr)
    cReviewContentWebApi := webapi.NewCReviewContentWebApi(cReviewContentApplication)
    r = cReviewContentWebApi.CreateAccessPoint(r)

    cReviewStatusService := service.NewCReviewStatusService(persistence, numberUtil)
    cReviewStatusApplication := application.NewCReviewStatusApplication(cReviewStatusService, tr)
    cReviewStatusWebApi := webapi.NewCReviewStatusWebApi(cReviewStatusApplication)
    r = cReviewStatusWebApi.CreateAccessPoint(r)

    cServiceService := service.NewCServiceService(persistence, numberUtil)
    cServiceApplication := application.NewCServiceApplication(cServiceService, tr)
    cServiceWebApi := webapi.NewCServiceWebApi(cServiceApplication)
    r = cServiceWebApi.CreateAccessPoint(r)

    cShopEquipmentService := service.NewCShopEquipmentService(persistence, numberUtil)
    cShopEquipmentApplication := application.NewCShopEquipmentApplication(cShopEquipmentService, tr)
    cShopEquipmentWebApi := webapi.NewCShopEquipmentWebApi(cShopEquipmentApplication)
    r = cShopEquipmentWebApi.CreateAccessPoint(r)

    cShopImageFilterCategoryService := service.NewCShopImageFilterCategoryService(persistence, numberUtil)
    cShopImageFilterCategoryApplication := application.NewCShopImageFilterCategoryApplication(cShopImageFilterCategoryService, tr)
    cShopImageFilterCategoryWebApi := webapi.NewCShopImageFilterCategoryWebApi(cShopImageFilterCategoryApplication)
    r = cShopImageFilterCategoryWebApi.CreateAccessPoint(r)
    
    cShopMaintenanceLabelService := service.NewCShopMaintenanceLabelService(persistence, numberUtil)
    cShopMaintenanceLabelApplication := application.NewCShopMaintenanceLabelApplication(cShopMaintenanceLabelService, tr)
    cShopMaintenanceLabelWebApi := webapi.NewCShopMaintenanceLabelWebApi(cShopMaintenanceLabelApplication)
    r = cShopMaintenanceLabelWebApi.CreateAccessPoint(r)

    cShopStatusService := service.NewCShopStatusService(persistence, numberUtil)
    cShopStatusApplication := application.NewCShopStatusApplication(cShopStatusService, tr)
    cShopStatusWebApi := webapi.NewCShopStatusWebApi(cShopStatusApplication)
    r = cShopStatusWebApi.CreateAccessPoint(r)

    cShopPaymentMethodService := service.NewCShopPaymentMethodService(persistence, numberUtil)
    cShopPaymentMethodApplication := application.NewCShopPaymentMethodApplication(cShopPaymentMethodService, tr)
    cShopPaymentMethodWebApi := webapi.NewCShopPaymentMethodWebApi(cShopPaymentMethodApplication)
    r = cShopPaymentMethodWebApi.CreateAccessPoint(r)

    cStateService := service.NewCStateService(persistence, numberUtil)
    cStateApplication := application.NewCStateApplication(cStateService, tr)
    cStateWebApi := webapi.NewCStateWebApi(cStateApplication)
    r = cStateWebApi.CreateAccessPoint(r)

    cTellCountryService := service.NewCTellCountryService(persistence, numberUtil)
    cTellCountryApplication := application.NewCTellCountryApplication(cTellCountryService, tr)
    cTellCountryWebApi := webapi.NewCTellCountryWebApi(cTellCountryApplication)
    r = cTellCountryWebApi.CreateAccessPoint(r)

    cTreatmentTimeRangeService := service.NewCTreatmentTimeRangeService(persistence, numberUtil)
    cTreatmentTimeRangeApplication := application.NewCTreatmentTimeRangeApplication(cTreatmentTimeRangeService, tr)
    cTreatmentTimeRangeWebApi := webapi.NewCTreatmentTimeRangeWebApi(cTreatmentTimeRangeApplication)
    r = cTreatmentTimeRangeWebApi.CreateAccessPoint(r)

    cTypeOfContactService := service.NewCTypeOfContactService(persistence, numberUtil)
    cTypeOfContactApplication := application.NewCTypeOfContactApplication(cTypeOfContactService, tr)
    cTypeOfContactWebApi := webapi.NewCTypeOfContactWebApi(cTypeOfContactApplication)
    r = cTypeOfContactWebApi.CreateAccessPoint(r)

    cWeekdayService := service.NewCWeekdayService(persistence, numberUtil)
    cWeekdayApplication := application.NewCWeekdayApplication(cWeekdayService, tr)
    cWeekdayWebApi := webapi.NewCWeekdayWebApi(cWeekdayApplication)
    r = cWeekdayWebApi.CreateAccessPoint(r)

    defaultFeeMasterService := service.NewDefaultFeeMasterService(persistence, numberUtil)
    defaultFeeMasterApplication := application.NewDefaultFeeMasterApplication(defaultFeeMasterService, tr)
    defaultFeeMasterWebApi := webapi.NewDefaultFeeMasterWebApi(defaultFeeMasterApplication)
    r = defaultFeeMasterWebApi.CreateAccessPoint(r)

    cConcernService := service.NewCConcernService(persistence, numberUtil)
    cConcernApplication := application.NewCConcernApplication(cConcernService, tr)
    cConcernWebApi := webapi.NewCConcernWebApi(cConcernApplication)
    r = cConcernWebApi.CreateAccessPoint(r)

    currencyExchangeRateBatchService := service.NewCurrencyExchangeRateBatchService(persistence, numberUtil, api)
    currencyExchangeRateBatchApplication := application.NewCurrencyExchangeRateBatchApplication(currencyExchangeRateBatchService, tr)
    currencyExchangeRateBatchWebApi := webapi.NewCurrencyExchangeRateBatchWebApi(currencyExchangeRateBatchApplication)
    r = currencyExchangeRateBatchWebApi.CreateAccessPoint(r)

    q := query.NewQuery(tr)
    queryCurrencyForPayment := webapi.NewQueryCurrencyForPaymentWebApi(q, persistence, tr)
    r = queryCurrencyForPayment.CreateAccessPoint(r)
    
    queryCAgeRange := webapi.NewQueryCAgeRangeWebApi(q, persistence, tr)
    r = queryCAgeRange.CreateAccessPoint(r)
    
    queryCArea := webapi.NewQueryCAreaWebApi(q, persistence, tr)
    r = queryCArea.CreateAccessPoint(r)
    
    queryCState := webapi.NewQueryCStateWebApi(q, persistence, tr)
    r = queryCState.CreateAccessPoint(r)


	port := settings.GetApiSettings()
	r.Run(port)
}