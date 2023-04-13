package main

import (
	"encoding/hex"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/oneLogin/utils"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Swagger Example API
// @version                     0.0.1
// @description                 This is a sample Server pets
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	PriKey := "MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBALSsLeIkMDoN0aWxoqK0sENGjNEC84nfHAFykEwWdJtXN+sJxAYkMo1A8e/hljt71YRL1QkYmZtek/Vi8XWhKD0Fme3rih50T3+Koe6YkLBKLEjM6MEmdV0z3lOmyCxbIYORuv5giiRUFSrAbKFFh0vUzezmyFefAAfqA316+eDhAgMBAAECgYAIt433LTvOcUA+OFXad9FRTaQZqYTKkCMvxrFDmonBvPGLu4rjqPdvbUS/CClRcWYZ3fbHW5J9tpB49G8l98KTKlTmbHWm2m6UBhqHxuNnhGK24muz99c1GzbXWgppxa2HeiA0yu2VT56RDfUrZGqvv+jD\nnpbuRP3YInXiMirF4QJBAOd7rH2ttqw6c1HHvL+9f3cWuUAXHcTqNhsWVlU2HslDCJX2VDW76lF+nZLp7xeNsHb8gxBOvcfd2YYs4tyiHa0CQQDHztosmZY/JDtJUtY+5xMMOWegCkeblCTLtc7v8zjLu2lcn6wEuShDXfxgVg7VPCTKbTiqYy+gEg9GsNKj\nxQ6FAkAMQ3EP93QGC9KwMnS9c7ydAoct7guVsxLKvJQ2T3eyEesShspPTnVLe/m9Hseb59XBd/85jfJf9FDh2t7p8WzBAkAKs4Nv3BH188TRGoSq/clBYFmycpp/NKH73xLkOwyRrMnp0gtufVQwt3nq1vEYbo4x4UOlrIZCdnUm/hVp/AXZAkAVhL9u6RGGM8TY4sDI5e9DZM/oyvM5R96uGF1Xb6LK7Znicb/P7eD1tWd0wCTI34e8iE7qxzfnNz02SMfMuWaE"
	publicKey := "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC0rC3iJDA6DdGlsaKitLBDRozRAvOJ3xwBcpBMFnSbVzfrCcQGJDKNQPHv4ZY7e9WES9UJGJmbXpP1YvF1oSg9BZnt64oedE9/iqHumJCwSixIzOjBJnVdM95TpsgsWyGDkbr+YIokVBUqwGyhRYdL1M3s5shXnwAH6gN9evng4QIDAQAB"

	//PriKey := "MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAIC36ljhxVqx1DgpUC8ZRhpScOs9X3FESzSjcvLVv2SXpwuMEHJIj7RpByk7rJ4FVHfXl52cotMdDVWi7BRQQz9xhIxyJG/H4zeeoDQHjg7PHYFDTVdsN/DRHDv1tipOBny0IrrGWW677Sid6Z/sMeUpKPL/8+eWh7JwO1arUn3hAgMBAAECgYAOzJZ+F58oOU/sERvt/lroBdiDw2+oxzBaYfyCXP7/YsxK8JSnfx4+oOC45eqH1JcMnFYLQgoaebmhwfSgtUW15+sdBReIaciZgS/p5GNk+j7pxgP4N9AJQiXGHxG235pibBAEVH92qnBHgIpN40Hza1CWFhgTw1GOaFMPcz8qvQJBAOHGkbh4cn7bNxhslz0nvARjAMyekZdU4QCyQCCi0J3SfvxvzDRrvQVqun5AxUHTUNWhlRY3yub7DnYsFkqx3mcCQQCR8ybrZGMqNzeaxq/Fs5HTzbJosZ5k7nDfs5kSrW2NgY3+SJyJk+StrYvS4hZOXwg0iyiMrB1iwWmlaZkb8YR3AkBTauBwPeBfynLizUxbxhCLtmCXOYclWLEBZtqWtFFL3ngYoN3cCGqAU9yvxRKcrYzSQa8p1FddXCkNtGBQHMPFAkEAgCyNSn6QBBwYDipdZX+tGthzzTPnyfYJVLwyO0/pfTOA0wdLyhsC4nAd8qaxNkSJPTPU+a2R5Q+8yxLw7rRtQwJAeSj6/m5EOh+dCad0hyOZnzIO4Xaf8O9vujm6DeZ/smt3eLZis3OrlOwiReYb48R5N4OtUcJXMfghBFTf7QBMSw=="
	//publicKey := "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCAt+pY4cVasdQ4KVAvGUYaUnDrPV9xREs0o3Ly1b9kl6cLjBBySI+0aQcpO6yeBVR315ednKLTHQ1VouwUUEM/cYSMciRvx+M3nqA0B44Ozx2BQ01XbDfw0Rw79bYqTgZ8tCK6xlluu+0onemf7DHlKSjy//PnloeycDtWq1J94QIDAQAB"
	//
	str := "123456"
	utils.Sha256WithRsa(str, PriKey)
	data := utils.RSA_Encrypt([]byte(str), publicKey)
	aaa, _ := hex.DecodeString(data)
	newStr := utils.RSA_Decrypt(aaa, PriKey)
	fmt.Sprintf("%s,%s", data, newStr)

	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
