package routes
import(
	"github.com/gin-gonic/gin"
	controller "golang-restuarent_management/controllers"
)
func itemRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.get("/item",controller.GetItems())
	incomingRoutes.get("/item/:item_id",controller.GetItems())
	incomingRoutes.post("/item",controller.CreateItem())
	incomingRoutes.patch("/item/:item_id",controller.UpdateItem())

}