package system

type ServiceGroup struct {
	CacheSmsService
	JwtService
	ApiService
	MenuService
	UserService
	CasbinService
	InitDBService
	AutoCodeService
	BaseMenuService
	AuthorityService
	DictionaryService
	SystemConfigService
	AutoCodeHistoryService
	OperationRecordService
	DictionaryDetailService
	AuthorityBtnService
}
