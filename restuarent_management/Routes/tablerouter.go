package routes
import(
	"github.com/gin-gonic/gin"
	controller "golang-restuarent_management/controllers"
)
func tableRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.get("/table",controller.GetTables())
	incomingRoutes.get("/table/:table_id",controller.GetTables())
	incomingRoutes.post("/table",controller.CreateTable())
	incomingRoutes.patch("/table/:table_id",controller.UpdateTable())

}
