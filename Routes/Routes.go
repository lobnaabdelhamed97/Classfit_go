package Routes

import (
	"CLASSFIT_GO/Controllers"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/get-game")
	{
		grp1.POST("game", Controllers.GetGames)
		grp1.GET("game/:gm_id", Controllers.GetGameByID)

	}
	grp2 := r.Group("/view-absence")
	{
		grp2.GET("absence", Controllers.GetAbsences)
		grp2.GET("absence/:absence_id", Controllers.GetAbsenceByID)
		grp2.POST("absence", Controllers.CreateAbsence)

	}
	r.POST("/view-game", Controllers.ViewGame)
    r.POST("/mem-info", Controllers.Mem_info)

	return r
}
