TEMPLATE_DATA = internal/templatedata
TEST_DATA = internal/testdata

TEMPLATES = \
	$(TEMPLATE_DATA)/base.tmpl \
	$(TEMPLATE_DATA)/client.tmpl \
	$(TEMPLATE_DATA)/client_impl.tmpl \
	$(TEMPLATE_DATA)/client_stream.tmpl \
	$(TEMPLATE_DATA)/fx.tmpl \
	$(TEMPLATE_DATA)/parameters.tmpl \
	$(TEMPLATE_DATA)/server.tmpl \
	$(TEMPLATE_DATA)/server_impl.tmpl \
	$(TEMPLATE_DATA)/server_stream.tmpl

.PHONY: generate
generate:
	@prototool generate $(TEST_DATA)

.PHONY: templates
templates: $(TEMPLATES)
	@go get -u github.com/jteeuwen/go-bindata/go-bindata
	@go-bindata -pkg templatedata -o ./$(TEMPLATE_DATA)/bindata.go ./$(TEMPLATE_DATA)
