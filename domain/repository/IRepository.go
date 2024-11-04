package repository

type Repository interface {
	CurrencyExchangeRateRepository
	CurrencyForPaymentRepository
	CAccountWithdrawalReasonRepository
	CAreaRepository
	CAgeRangeRepository
	CBillingContentRepository
	CBillingStatusRepository
	CBookingChanelRepository
	CBookingMethodRepository
	CBookingStatusRepository
	CCheckoutMethodRepository
	CCheckoutStatusRepository
	CContactStatusRepository
	CContentsLabelRepository
	CContentsCategoryRepository
	CCheckoutTimingRepository
	CContentsStatusRepository
	CShopContractPlanItemRepository
	CCountryRepository
	CCouponTargetAttrRepository
	CCurrencyRepository
	CGenderRepository
	CInvoiceStatusRepository
	CLanguageRepository
	CMailTemplateRepository
	CMenuLabelRepository
	COrderTypeRepository
	CPayoutItemCategoryRepository
	CPayoutMethodRepository
	CPayoutStatusRepository
	CRecommendLabelRepository
	CReviewCategoryRepository
	CReviewContentRepository
	CReviewStatusRepository
	CServiceRepository
	CShopEquipmentRepository
	CShopImageFilterCategoryRepository
	CShopMaintenanceLabelRepository
	CShopStatusRepository
	CShopPaymentMethodRepository
	CStateRepository
	CTellCountryRepository
	CTreatmentTimeRangeRepository
	CTypeOfContactRepository
	CWeekdayRepository
	DefaultFeeMasterRepository
	CConcernRepository
}
