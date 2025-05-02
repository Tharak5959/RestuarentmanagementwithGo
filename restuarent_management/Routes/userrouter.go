package routes
import(
	"github.com/gin-gonic/gin"
	controller "golang-restuarent_management/controllers"
)
func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.get("/users",controller.Getusers())
	incomingRoutes.get("/users/:user_id",controller.Getusers())
	incomingRoutes.post("users/signup",controller.signup())
	incomingRoutes.post("users/login",controller.login())

}