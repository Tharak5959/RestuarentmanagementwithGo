package routes
import(
	"github.com/gin-gonic/gin"
	controller "golang-restuarent_management/controllers"
)
func menuRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.get("/menu",controller.GetMenus())
	incomingRoutes.get("/menu/:menu_id",controller.GetMenus())
	incomingRoutes.post("/menu",controller.CreateMenu())
	incomingRoutes.patch("/menu/:menu_id",controller.UpdateMenu())

}