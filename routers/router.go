package routers

import (
	"PC/api"
	"PC/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func InitRouter() *gin.Engine {

	// 初始化默认路由
	router := gin.Default()
	router.LoadHTMLFiles("./view/index.html")
	router.Static("/static", "./static")
	router.StaticFS("/myPackge", http.Dir("./myPackge"))
	router.StaticFS("/uploadImg", http.Dir("./uploadImg"))
	router.GET("/admin", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	//websocket
	router.GET("/v1/ping", api.SocketOrder)
	router.GET("/v1/adminSocket", api.NewSocketClient)
	//支持跨域
	router.Use(middlewares.Cors())
	v1 := router.Group("api")

	{
		//首页拼单列表
		v1.POST("/v1/orders", api.GetOrderList)

		//	获取订单详情 提交订单
		v1.POST("/v1/orders/info", api.GetClassInfo)

		//支付订单
		v1.POST("/v1/payorder", api.PayCourseOrder)

		//生成订单列表
		v1.POST("/v1/orders/createOrder", api.CreateWeChatOrder)

		//支付订单后填写地址信息完成订单
		v1.POST("/v1/address", api.GetAddress)
		//获取地址列表
		v1.POST("/v1/getAddress", api.GetAddressList)

		//根据选择的地址完成地址信息 传递地址id更新订单的address_id
		v1.POST("/v1/upAddressId", api.UpOrderAddress)

		//后台获取用户订单 参数 发货和未发货
		v1.POST("/v1/getOrders", api.GetSuccessOrders)
		//发货订单
		v1.POST("/v1/order/upSendType", api.UpOrderSendType)

		//添加分类 创建课程
		v1.POST("/v1/addcat", api.AddCatControl)
		//获取课程分类
		v1.POST("/v1/getCourseCat", api.GetCourseCat)

		//图片上传
		v1.POST("/v1/upload", api.UploadImg)

		//上传课程书籍
		v1.POST("/v1/createCourseBook", api.CreateCourseBook)
		//获取书籍
		v1.POST("/v1/book/getbooks", api.GetCourseBook)
		//删除书籍
		v1.POST("/v1/book/delbooks", api.DelCourseBook)



	}

	user := router.Group("user")
	{
		//用户信息 创建微信小程序用户
		user.POST("/v1/userInfo", api.GetUserInfo)
		//注册admin用户
		user.POST("/v1/createAdmin", api.CreateAdminUser)
		//登录admin用户
		user.POST("/v1/loginAdmin", api.LoginAdminUser)
		//获取登录admin用户信息
		user.POST("/v1/getAdminInfo", api.GetAdminUserInfo)
		//获取登录首页订单信息
		user.POST("/v1/adminHomeInfo", api.GetAdminHomeInfo)
	}

	admin := router.Group("admin")
	{
		//获取后台用户信息
		admin.POST("/getUsers", api.GetAdminUser)
		//	获取课程排行榜
		admin.POST("/courseRank", api.GetCourseRank)
		//	获取课程详情
		admin.POST("/courseDetail", api.GetCourseDetail)
		//	上传微信首页主题信息
		admin.POST("/weChatIndex", api.WeChatIndex)
		//	获取首页信息
		admin.POST("/weChatIndex/list", api.WeChatIndexList)
		//	上传首页图片单图
		admin.POST("/weChatIndex/uploadImg", api.WeChatIndexUploadImg)

		//	文章增加
		admin.POST("/help/createArticle", api.CreateArticle)
		//	文章查询
		admin.POST("/help/getArticle", api.GetArticle)

		//	文章查询一个根据id
		admin.POST("/help/getArticleone", api.GetArticleOne)

		//	用户添加订阅通知
		admin.POST("/user/addSubmit", api.UpdataUserInfo)

	}

	video := router.Group("video")
	{
		//创建room
		video.POST("/getRoomId", api.GetRoomID)
		//获取直播间列表
		video.POST("/getRoomList", api.GetRoomList)

	}

	//文章分类
	classify := router.Group("classify")
	{
		//创建CMS分类
		classify.POST("/createClassify", api.CreatClassify)
		//获取CMS列表
		classify.POST("/getClassify", api.GetClassify)
		//更新CMS
		classify.POST("/upClassify", api.UpdataClassify)

		//	查询分类对应的文章

		classify.POST("/getClassifyArticle", api.GetClassifyArticle)

	}

	banner := router.Group("banner")
	{
		//创建banner
		banner.POST("/createbanner", api.CreatBanner)
		//获取banner列表
		banner.POST("/getbanner", api.GetBanner)

		//获取banner一个 获取content
		banner.POST("/getbannerOne", api.GetBannerONE)

		//更新banner
		banner.POST("/upbanner", api.UpdataBanner)
	}

	studio := router.Group("studio")
	{
		//创建Studio
		studio.POST("/createstudio", api.CreatStudio)
		//获取Studio列表
		studio.POST("/getstudio", api.GetStudio)

		//获取Studio详情
		studio.POST("/getstudioOne", api.GetStudioOne)
		//更新Studio
		studio.POST("/upstudio", api.UpdataStudio)
	}

	// use ginSwagger middleware to
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router

}
