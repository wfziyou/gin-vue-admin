package system

import (
	"context"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initApi struct{}

const initOrderApi = system.InitOrderSystem + 1

// auto run
func init() {
	system.RegisterInit(initOrderApi, &initApi{})
}

func (i initApi) InitializerName() string {
	return sysModel.SysApi{}.TableName()
}

func (i *initApi) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysApi{})
}

func (i *initApi) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysApi{})
}

func (i *initApi) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	var base = sysModel.SysApiGroup{Name: "base", Description: "base", Sort: 1}
	var jwt = sysModel.SysApiGroup{Name: "base", Description: "base", Sort: 1}
	var sysUser = sysModel.SysApiGroup{Name: "系统用户", Description: "系统用户", Sort: 1}
	var api = sysModel.SysApiGroup{Name: "api", Description: "api", Sort: 1}
	var authority = sysModel.SysApiGroup{Name: "角色", Description: "角色", Sort: 1}
	var casbin = sysModel.SysApiGroup{Name: "casbin", Description: "casbin", Sort: 1}
	var menu = sysModel.SysApiGroup{Name: "菜单", Description: "菜单", Sort: 1}
	var file = sysModel.SysApiGroup{Name: "文件上传与下载", Description: "文件上传与下载", Sort: 1}
	var sysSevice = sysModel.SysApiGroup{Name: "系统服务", Description: "系统服务", Sort: 1}
	var customer = sysModel.SysApiGroup{Name: "客户", Description: "客户", Sort: 1}
	var autoCode = sysModel.SysApiGroup{Name: "代码生成器", Description: "代码生成器", Sort: 1}
	var autoCodeHistory = sysModel.SysApiGroup{Name: "代码生成器历史", Description: "代码生成器历史", Sort: 1}
	var sysDictionaryDetail = sysModel.SysApiGroup{Name: "系统字典详情", Description: "系统字典详情", Sort: 1}
	var sysDictionary = sysModel.SysApiGroup{Name: "系统字典", Description: "系统字典", Sort: 1}
	var sysOperationRecord = sysModel.SysApiGroup{Name: "操作记录", Description: "操作记录", Sort: 1}
	var simpleUploader = sysModel.SysApiGroup{Name: "断点续传(插件版)", Description: "断点续传(插件版)", Sort: 1}
	var email = sysModel.SysApiGroup{Name: "email", Description: "email", Sort: 1}
	var authorityBtn = sysModel.SysApiGroup{Name: "按钮权限", Description: "按钮权限", Sort: 1}
	groups := []sysModel.SysApiGroup{
		base,
		jwt,
		sysUser,
		api,
		authority,
		casbin,
		menu,
		file,
		sysSevice,
		customer,
		autoCode,
		autoCodeHistory,
		sysDictionaryDetail,
		sysDictionary,
		sysOperationRecord,
		simpleUploader,
		email,
		authorityBtn,
	}

	if err := db.Create(&groups).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysApiGroup{}.TableName()+"表数据初始化失败!")
	}

	entities := []sysModel.SysApi{
		{GroupId: base.ID, Method: "POST", Path: "/base/login", Description: "用户登录(必选)"},

		{GroupId: jwt.ID, Method: "POST", Path: "/jwt/jsonInBlacklist", Description: "jwt加入黑名单(退出，必选)"},

		{GroupId: sysUser.ID, Method: "DELETE", Path: "/user/deleteUser", Description: "删除用户"},
		{GroupId: sysUser.ID, Method: "POST", Path: "/user/admin_register", Description: "用户注册"},
		{GroupId: sysUser.ID, Method: "POST", Path: "/user/getUserList", Description: "获取用户列表"},
		{GroupId: sysUser.ID, Method: "PUT", Path: "/user/setUserInfo", Description: "设置用户信息"},
		{GroupId: sysUser.ID, Method: "PUT", Path: "/user/setSelfInfo", Description: "设置自身信息(必选)"},
		{GroupId: sysUser.ID, Method: "GET", Path: "/user/getUserInfo", Description: "获取自身信息(必选)"},
		{GroupId: sysUser.ID, Method: "POST", Path: "/user/setUserAuthorities", Description: "设置权限组"},
		{GroupId: sysUser.ID, Method: "POST", Path: "/user/changePassword", Description: "修改密码（建议选择)"},
		{GroupId: sysUser.ID, Method: "POST", Path: "/user/setUserAuthority", Description: "修改用户角色(必选)"},
		{GroupId: sysUser.ID, Method: "POST", Path: "/user/resetPassword", Description: "重置用户密码"},

		{GroupId: api.ID, Method: "POST", Path: "/api/createApi", Description: "创建api"},
		{GroupId: api.ID, Method: "POST", Path: "/api/deleteApi", Description: "删除Api"},
		{GroupId: api.ID, Method: "POST", Path: "/api/updateApi", Description: "更新Api"},
		{GroupId: api.ID, Method: "POST", Path: "/api/getApiList", Description: "获取api列表"},
		{GroupId: api.ID, Method: "POST", Path: "/api/getAllApis", Description: "获取所有api"},
		{GroupId: api.ID, Method: "POST", Path: "/api/getApiById", Description: "获取api详细信息"},
		{GroupId: api.ID, Method: "DELETE", Path: "/api/deleteApisByIds", Description: "批量删除api"},

		{GroupId: authority.ID, Method: "POST", Path: "/authority/copyAuthority", Description: "拷贝角色"},
		{GroupId: authority.ID, Method: "POST", Path: "/authority/createAuthority", Description: "创建角色"},
		{GroupId: authority.ID, Method: "POST", Path: "/authority/deleteAuthority", Description: "删除角色"},
		{GroupId: authority.ID, Method: "PUT", Path: "/authority/updateAuthority", Description: "更新角色信息"},
		{GroupId: authority.ID, Method: "POST", Path: "/authority/getAuthorityList", Description: "获取角色列表"},
		{GroupId: authority.ID, Method: "POST", Path: "/authority/setDataAuthority", Description: "设置角色资源权限"},

		{GroupId: casbin.ID, Method: "POST", Path: "/casbin/updateCasbin", Description: "更改角色api权限"},
		{GroupId: casbin.ID, Method: "POST", Path: "/casbin/getPolicyPathByAuthorityId", Description: "获取权限列表"},

		{GroupId: menu.ID, Method: "POST", Path: "/menu/addBaseMenu", Description: "新增菜单"},
		{GroupId: menu.ID, Method: "POST", Path: "/menu/getMenu", Description: "获取菜单树(必选)"},
		{GroupId: menu.ID, Method: "POST", Path: "/menu/deleteBaseMenu", Description: "删除菜单"},
		{GroupId: menu.ID, Method: "POST", Path: "/menu/updateBaseMenu", Description: "更新菜单"},
		{GroupId: menu.ID, Method: "POST", Path: "/menu/getBaseMenuById", Description: "根据id获取菜单"},
		{GroupId: menu.ID, Method: "POST", Path: "/menu/getMenuList", Description: "分页获取基础menu列表"},
		{GroupId: menu.ID, Method: "POST", Path: "/menu/getBaseMenuTree", Description: "获取用户动态路由"},
		{GroupId: menu.ID, Method: "POST", Path: "/menu/getMenuAuthority", Description: "获取指定角色menu"},
		{GroupId: menu.ID, Method: "POST", Path: "/menu/addMenuAuthority", Description: "增加menu和角色关联关系"},

		{GroupId: file.ID, Method: "GET", Path: "/fileUploadAndDownload/findFile", Description: "寻找目标文件（秒传）"},
		{GroupId: file.ID, Method: "POST", Path: "/fileUploadAndDownload/breakpointContinue", Description: "断点续传"},
		{GroupId: file.ID, Method: "POST", Path: "/fileUploadAndDownload/breakpointContinueFinish", Description: "断点续传完成"},
		{GroupId: file.ID, Method: "POST", Path: "/fileUploadAndDownload/removeChunk", Description: "上传完成移除文件"},

		{GroupId: file.ID, Method: "POST", Path: "/fileUploadAndDownload/upload", Description: "文件上传示例"},
		{GroupId: file.ID, Method: "POST", Path: "/fileUploadAndDownload/deleteFile", Description: "删除文件"},
		{GroupId: file.ID, Method: "POST", Path: "/fileUploadAndDownload/editFileName", Description: "文件名或者备注编辑"},
		{GroupId: file.ID, Method: "POST", Path: "/fileUploadAndDownload/getFileList", Description: "获取上传文件列表"},

		{GroupId: sysSevice.ID, Method: "POST", Path: "/system/getServerInfo", Description: "获取服务器信息"},
		{GroupId: sysSevice.ID, Method: "POST", Path: "/system/getSystemConfig", Description: "获取配置文件内容"},
		{GroupId: sysSevice.ID, Method: "POST", Path: "/system/setSystemConfig", Description: "设置配置文件内容"},

		{GroupId: customer.ID, Method: "PUT", Path: "/customer/customer", Description: "更新客户"},
		{GroupId: customer.ID, Method: "POST", Path: "/customer/customer", Description: "创建客户"},
		{GroupId: customer.ID, Method: "DELETE", Path: "/customer/customer", Description: "删除客户"},
		{GroupId: customer.ID, Method: "GET", Path: "/customer/customer", Description: "获取单一客户"},
		{GroupId: customer.ID, Method: "GET", Path: "/customer/customerList", Description: "获取客户列表"},

		{GroupId: autoCode.ID, Method: "GET", Path: "/autoCode/getDB", Description: "获取所有数据库"},
		{GroupId: autoCode.ID, Method: "GET", Path: "/autoCode/getTables", Description: "获取数据库表"},
		{GroupId: autoCode.ID, Method: "POST", Path: "/autoCode/createTemp", Description: "自动化代码"},
		{GroupId: autoCode.ID, Method: "POST", Path: "/autoCode/preview", Description: "预览自动化代码"},
		{GroupId: autoCode.ID, Method: "GET", Path: "/autoCode/getColumn", Description: "获取所选table的所有字段"},
		{GroupId: autoCode.ID, Method: "POST", Path: "/autoCode/createPlug", Description: "自动创建插件包"},
		{GroupId: autoCode.ID, Method: "POST", Path: "/autoCode/installPlugin", Description: "安装插件"},

		{GroupId: autoCode.ID, Method: "POST", Path: "/autoCode/createPackage", Description: "生成包(package)"},
		{GroupId: autoCode.ID, Method: "POST", Path: "/autoCode/getPackage", Description: "获取所有包(package)"},
		{GroupId: autoCode.ID, Method: "POST", Path: "/autoCode/delPackage", Description: "删除包(package)"},

		{GroupId: autoCodeHistory.ID, Method: "POST", Path: "/autoCode/getMeta", Description: "获取meta信息"},
		{GroupId: autoCodeHistory.ID, Method: "POST", Path: "/autoCode/rollback", Description: "回滚自动生成代码"},
		{GroupId: autoCodeHistory.ID, Method: "POST", Path: "/autoCode/getSysHistory", Description: "查询回滚记录"},
		{GroupId: autoCodeHistory.ID, Method: "POST", Path: "/autoCode/delSysHistory", Description: "删除回滚记录"},

		{GroupId: sysDictionaryDetail.ID, Method: "PUT", Path: "/sysDictionaryDetail/updateSysDictionaryDetail", Description: "更新字典内容"},
		{GroupId: sysDictionaryDetail.ID, Method: "POST", Path: "/sysDictionaryDetail/createSysDictionaryDetail", Description: "新增字典内容"},
		{GroupId: sysDictionaryDetail.ID, Method: "DELETE", Path: "/sysDictionaryDetail/deleteSysDictionaryDetail", Description: "删除字典内容"},
		{GroupId: sysDictionaryDetail.ID, Method: "GET", Path: "/sysDictionaryDetail/findSysDictionaryDetail", Description: "根据ID获取字典内容"},
		{GroupId: sysDictionaryDetail.ID, Method: "GET", Path: "/sysDictionaryDetail/getSysDictionaryDetailList", Description: "获取字典内容列表"},

		{GroupId: sysDictionary.ID, Method: "POST", Path: "/sysDictionary/createSysDictionary", Description: "新增字典"},
		{GroupId: sysDictionary.ID, Method: "DELETE", Path: "/sysDictionary/deleteSysDictionary", Description: "删除字典"},
		{GroupId: sysDictionary.ID, Method: "PUT", Path: "/sysDictionary/updateSysDictionary", Description: "更新字典"},
		{GroupId: sysDictionary.ID, Method: "GET", Path: "/sysDictionary/findSysDictionary", Description: "根据ID获取字典"},
		{GroupId: sysDictionary.ID, Method: "GET", Path: "/sysDictionary/getSysDictionaryList", Description: "获取字典列表"},

		{GroupId: sysOperationRecord.ID, Method: "POST", Path: "/sysOperationRecord/createSysOperationRecord", Description: "新增操作记录"},
		{GroupId: sysOperationRecord.ID, Method: "GET", Path: "/sysOperationRecord/findSysOperationRecord", Description: "根据ID获取操作记录"},
		{GroupId: sysOperationRecord.ID, Method: "GET", Path: "/sysOperationRecord/getSysOperationRecordList", Description: "获取操作记录列表"},
		{GroupId: sysOperationRecord.ID, Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecord", Description: "删除操作记录"},
		{GroupId: sysOperationRecord.ID, Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecordByIds", Description: "批量删除操作历史"},

		{GroupId: simpleUploader.ID, Method: "POST", Path: "/simpleUploader/upload", Description: "插件版分片上传"},
		{GroupId: simpleUploader.ID, Method: "GET", Path: "/simpleUploader/checkFileMd5", Description: "文件完整度验证"},
		{GroupId: simpleUploader.ID, Method: "GET", Path: "/simpleUploader/mergeFileMd5", Description: "上传完成合并文件"},

		{GroupId: email.ID, Method: "POST", Path: "/email/emailTest", Description: "发送测试邮件"},
		{GroupId: email.ID, Method: "POST", Path: "/email/emailSend", Description: "发送邮件示例"},

		{GroupId: authorityBtn.ID, Method: "POST", Path: "/authorityBtn/setAuthorityBtn", Description: "设置按钮权限"},
		{GroupId: authorityBtn.ID, Method: "POST", Path: "/authorityBtn/getAuthorityBtn", Description: "获取已有按钮权限"},
		{GroupId: authorityBtn.ID, Method: "POST", Path: "/authorityBtn/canRemoveAuthorityBtn", Description: "删除按钮"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysApi{}.TableName()+"表数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initApi) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("path = ? AND method = ?", "/authorityBtn/canRemoveAuthorityBtn", "POST").
		First(&sysModel.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
