REPO:=github.com/pei0804/goa-code-reading

gen: clean generate imports

bootstrap:
	@goagen bootstrap --force -d $(REPO)/design

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf tool
	@rm -rf swagger

generate:
	@goagen app     -d $(REPO)/design
	@goagen swagger -d $(REPO)/design
	@goagen client  -d $(REPO)/design

imports:
	@which gorep || go get -v github.com/novalagung/gorep
	@gorep -path="./tool/cli" \
          -from="../client" \
          -to="$(REPO)/client"
	@gorep -path="./" \
          -from="../app" \
          -to="$(REPO)/app"
	@gorep -path="./models" \
          -from="../app" \
          -to="$(REPO)/app"
	@gorep -path="./app" \
          -from="../app" \
          -to="$(REPO)/app"
	@gorep -path="./" \
          -from="../client" \
          -to="$(REPO)/client"
	@gorep -path="./" \
          -from="../tool/cli" \
          -to="$(REPO)/tool/cli"
