package routes
import(
	"github.com/gin-gonic/gin"
	controller"golang-restuarent_management/controllers"
)
func OrderRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.get("/order",controller.GetOrders())
	incomingRoutes.get("/order/:order_id",controller.GetOrders())
	incomingRoutes.post("/order",controller.CreateOrder())
	incomingRoutes.patch("/order/:order_id",controller.UpdateOrder())

}