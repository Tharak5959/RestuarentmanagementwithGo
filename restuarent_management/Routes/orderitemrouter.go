package routes
import(
	"github.com/gin-gonic/gin"
	controller "golang-restuarent_management/controllers"
)
func orderitemRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.get("/orderitem",controller.GetOrderitems())
	incomingRoutes.get("/orderitem/:orderitem_id",controller.GetOrderitems())
	incomingRoutes.post("/orderitem",controller.CreateOrderitem())
	incomingRoutes.patch("/orderitem/:orderitem_id",controller.UpdateOrderitem())

}