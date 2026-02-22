#
.PHONY: dummy

GOCMD:=$(shell command -v go)
GOLINT:=$(shell command -v revive)


GOLINT_OPTS=-formatter stylish

tidy:	dummy
	$(RM) *~

gofumpt:	tidy
	$(GOCMD) fmt

clean:	tidy
	$(GOCMD) clean

lint:	tidy
	if [ -n $(GOLINT) ]; then $(GOLINT) $(GOLINT_OPTS); fi
