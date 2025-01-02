package constants

const BOT_COMMAND_PREFIX = "!"

const (
	COMMAND_HELP                  = "help"
	COMMAND_REMIND                = "remind"
	COMMAND_RESET_CART            = "resetcart"
	COMMAND_RESET_CART_BY_USER_ID = "reset_cart_by_user_id"
	COMMAND_RESET_CART_BY_EMAIL   = "reset_cart_by_email"
	COMMAND_READY_PICK            = "readypick"
	COMMAND_SHOW_WAREHOUSE        = "showwarehouse"
	COMMAND_RESET_SHOW_WAREHOUSE  = "resetshowwarehouse"
	COMMAND_PICK                  = "pick"
	COMMAND_COUNT_KPI             = "kpi"
	COMMAND_COUNT_PROD_KPI        = "prod kpi"
	COMMAND_PICK_PACK             = "runpickpack"
)

const USER_ID_REGEX = `<@(\d+)>`
