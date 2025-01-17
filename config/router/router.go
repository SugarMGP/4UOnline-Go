package router

import (
	"4u-go/app/controllers/activityController"
	"4u-go/app/controllers/adminController"
	"4u-go/app/controllers/announcementController"
	"4u-go/app/controllers/collegeController"
	"4u-go/app/controllers/lostAndFoundController"
	"4u-go/app/controllers/objectController"
	"4u-go/app/controllers/qrcodeController"
	"4u-go/app/controllers/userController"
	"4u-go/app/controllers/websiteController"
	"4u-go/app/midwares"
	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init(r *gin.Engine) {
	const pre = "/api"

	api := r.Group(pre)
	{
		user := api.Group("/user")
		{
			user.POST("/create/student/wechat", userController.BindOrCreateStudentUserFromWechat)
			user.POST("/create/student", userController.CreateStudentUser)

			user.POST("/login/wechat", userController.WeChatLogin)
			user.POST("/login", userController.AuthByPassword)
			user.POST("/login/session", userController.AuthBySession)

			user.POST("/upload", objectController.UploadFile)

			user.POST("/repass", midwares.CheckLogin, userController.ChangePassword)
			user.DELETE("/delete", midwares.CheckLogin, userController.DeleteAccount)
		}

		admin := api.Group("/admin")
		{
			admin.POST("/create/key", adminController.CreateAdminByKey)

			adminLostAndFound := admin.Group("/lost-and-found", midwares.CheckAdmin)
			{
				adminLostAndFound.PUT("", lostAndFoundController.ReviewLostAndFound)
				adminLostAndFound.PUT("/update", lostAndFoundController.UpdateLostAndFound)
			}

			adminActivity := admin.Group("/activity", midwares.CheckAdmin)
			{
				adminActivity.POST("", activityController.CreateActivity)
				adminActivity.PUT("", activityController.UpdateActivity)
				adminActivity.DELETE("", activityController.DeleteActivity)
			}

			adminAnnouncement := admin.Group("/announcement", midwares.CheckAdmin)
			{
				adminAnnouncement.POST("", announcementController.CreateAnnouncement)
				adminAnnouncement.PUT("", announcementController.UpdateAnnouncement)
				adminAnnouncement.DELETE("", announcementController.DeleteAnnouncement)
			}

			adminCollege := admin.Group("/college", midwares.CheckSuperAdmin)
			{
				adminCollege.POST("", collegeController.CreateCollege)
				adminCollege.PUT("", collegeController.UpdateCollege)
				adminCollege.DELETE("", collegeController.DeleteCollege)
			}

			adminWebsite := admin.Group("/website", midwares.CheckAdmin)
			{
				adminWebsite.POST("", websiteController.CreateWebsite)
				adminWebsite.DELETE("", websiteController.DeleteWebsite)
				adminWebsite.PUT("", websiteController.UpdateWebsite)
				adminWebsite.GET("/list", websiteController.GetEditableWebsites)
			}

			adminQrcode := admin.Group("/qrcode", midwares.CheckAdmin)
			{
				adminQrcode.POST("", qrcodeController.CreateQrcode)
				adminQrcode.DELETE("", midwares.CheckSuperAdmin, qrcodeController.DeleteQrcode)
				adminQrcode.GET("", qrcodeController.GetQrcode)
				adminQrcode.POST("/list", qrcodeController.GetList)
				adminQrcode.PUT("", qrcodeController.UpdateQrcode)
			}
		}

		activity := api.Group("/activity")
		{
			activity.GET("/list", activityController.GetActivityList)
			activity.GET("", activityController.GetActivity)
		}

		announcement := api.Group("/announcement")
		{
			announcement.GET("/list", announcementController.GetAnnouncementList)
			announcement.GET("", announcementController.GetAnnouncement)
		}

		college := api.Group("/college")
		{
			college.GET("/list", collegeController.GetCollegeList)
		}

		website := api.Group("/website")
		{
			website.GET("/list", websiteController.GetWebsiteList)
		}

		lostAndFound := api.Group("/lost-and-found")
		{
			lostAndFound.POST("", midwares.CheckLogin, lostAndFoundController.CreateLostAndFound)
			lostAndFound.DELETE("", midwares.CheckLogin, lostAndFoundController.DeleteLostAndFound)
			lostAndFound.GET("/list", lostAndFoundController.GetLostAndFoundList)
			lostAndFound.GET("", midwares.CheckLogin, lostAndFoundController.GetLostAndFoundContact)
			lostAndFound.GET("/latest", lostAndFoundController.GetLatestLostAndFound)
			lostAndFound.GET("/user", midwares.CheckLogin, lostAndFoundController.GetUserLostAndFoundStatus)
			lostAndFound.PUT("/user", midwares.CheckLogin, lostAndFoundController.UpdateLostAndFoundStatus)
		}

		track := api.Group("/track")
		{
			track.GET("/qrcode/scan_count", qrcodeController.ScanCount)
		}
	}
}
