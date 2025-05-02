package routes
import(
	"github.com/gin-gonic/gin"
	controller "golang-restuarent_management/controllers"
)
func foodRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.get("/foods",controller.GetFoods())
	incomingRoutes.get("/foods/:food_id",controller.GetFoods())
	incomingRoutes.post("/foods",controller.CreateFood())
	incomingRoutes.patch("/foods/:food_id",controller.UpdateFood())

}
