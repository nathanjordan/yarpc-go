TEMPLATES = \
	internal/template/decorator.tmpl \
	internal/template/gluefx.tmpl \
	internal/template/integration.tmpl \
	internal/template/method.tmpl \
	internal/template/wrapper.tmpl

templates: $(TEMPLATES)
	go get -u github.com/jteeuwen/go-bindata/go-bindata
	go-bindata -pkg templatedata -o ./internal/templatedata/bindata.go ./internal/templatedata/
