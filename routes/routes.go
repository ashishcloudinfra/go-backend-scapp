package routes

import (
	"github.com/gorilla/mux"

	"server.simplifycontrol.com/routes/account"
	"server.simplifycontrol.com/routes/assets"
	"server.simplifycontrol.com/routes/auth"
	"server.simplifycontrol.com/routes/budget"
	"server.simplifycontrol.com/routes/company"
	"server.simplifycontrol.com/routes/companyType"
	"server.simplifycontrol.com/routes/event"
	"server.simplifycontrol.com/routes/eventRoom"
	"server.simplifycontrol.com/routes/health"
	"server.simplifycontrol.com/routes/inventory"
	"server.simplifycontrol.com/routes/membershipPlan"
	"server.simplifycontrol.com/routes/menuItem"
	"server.simplifycontrol.com/routes/notification"
	"server.simplifycontrol.com/routes/orders"
	"server.simplifycontrol.com/routes/payment"
	"server.simplifycontrol.com/routes/personalDetail"
	personalfinance "server.simplifycontrol.com/routes/personalFinance"
	"server.simplifycontrol.com/routes/scripts"
)

func InitializeRoutes() *mux.Router {
	mux := mux.NewRouter()

	companyType.RegisterCompanyTypeRoutes(mux)
	auth.RegisterAuthRoutes(mux)
	company.RegisterCompanyRoutes(mux)
	account.RegisterAccountRoutes(mux)
	health.RegisterHealthRoutes(mux)
	personalDetail.RegisterPersonalDetailsRoutes(mux)
	membershipPlan.RegisterMembershipPlanRoutes(mux)
	eventRoom.RegisterEventRoomRoutes(mux)
	event.RegisterEventRoutes(mux)
	menuItem.RegisterMenuItemRoutes(mux)
	orders.RegisterOrdersRoutes(mux)
	inventory.RegisterInventoryRoutes(mux)
	budget.RegisterBudgetRoutes(mux)
	notification.RegisterNotificationRoutes(mux)
	assets.RegisterAssetsRoutes(mux)
	personalfinance.RegisterPersonalFinanceRoutes(mux)
	scripts.RegisterScriptsRoutes(mux)
	payment.RegisterPaymentRoutes(mux)

	return mux
}
