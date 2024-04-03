package productapi

import (
	"net/http"

	"github.com/ardanlabs/service/app/core/crud/productapp"
	"github.com/ardanlabs/service/business/api/auth"
	midhttp "github.com/ardanlabs/service/business/api/mid/http"
	"github.com/ardanlabs/service/business/core/crud/product"
	"github.com/ardanlabs/service/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	ProductCore *product.Core
	Auth        *auth.Auth
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := midhttp.Authenticate(cfg.Auth)
	ruleAny := midhttp.Authorize(cfg.Auth, auth.RuleAny)
	ruleUserOnly := midhttp.Authorize(cfg.Auth, auth.RuleUserOnly)
	ruleAuthorizeProduct := midhttp.AuthorizeProduct(cfg.Auth, cfg.ProductCore)

	api := newAPI(productapp.NewCore(cfg.ProductCore))
	app.Handle(http.MethodGet, version, "/products", api.query, authen, ruleAny)
	app.Handle(http.MethodGet, version, "/products/{product_id}", api.queryByID, authen, ruleAuthorizeProduct)
	app.Handle(http.MethodPost, version, "/products", api.create, authen, ruleUserOnly)
	app.Handle(http.MethodPut, version, "/products/{product_id}", api.update, authen, ruleAuthorizeProduct)
	app.Handle(http.MethodDelete, version, "/products/{product_id}", api.delete, authen, ruleAuthorizeProduct)
}