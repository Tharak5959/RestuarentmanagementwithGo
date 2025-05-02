package routes
import(
	"github.com/gin-gonic/gin"
	controller "golang-restuarent_management/controllers"
)
func invoiceRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.get("/invoice",controller.GetInvoices())
	incomingRoutes.get("/invoice/:invoice_id",controller.GetInvoices())
	incomingRoutes.post("/invoice",controller.CreateInvoice())
	incomingRoutes.patch("/invoice/:invoice_id",controller.UpdateInvoice())

}