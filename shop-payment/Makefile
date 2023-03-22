


#SUB_PACKAGE  := $(subst $(shell git rev-parse --show-toplevel),,$(CURDIR))
SUB_PACKAGE := "shop-payment"
APP      := $(shell basename $(SUB_PACKAGE))


#export GO111MODULE=on
OUTPUT     = $(CURDIR)/output
CONF       = $(CURDIR)/conf
I18N       = $(CURDIR)/internationalization

.DEFAULT: all
all: build

build: clean prepare
	go build -o "$(OUTPUT)/bin/$(SUB_PACKAGE)" "$(CURDIR)/main.go"

clean:
	echo "====clean $(OUTPUT)"
	rm -rf "$(OUTPUT)"
	echo "====rm log"
	find . -name 'log' | xargs rm -rf

rmi:
	docker rm $(docker ps -a | grep "Exited" | awk '{print $1 }')
	docker rmi $(docker images | grep "none" | awk '{print $3}')

prepare:
	mkdir -p "$(OUTPUT)/log"
	cp -vr "$(CONF)" "$(OUTPUT)"
	cp -vr "$(I18N)" "$(OUTPUT)"
	cp -v "$(CURDIR)/control.sh" "$(OUTPUT)"

fmt:
	go fmt ./...

run:
	cd "$(OUTPUT)" && bin/$(APP)

swag:
	#swag init #只能是自己目錄下的
	swag init --parseDependency

i18n:
	goi18n extract -sourceLanguage zh -outdir internationalization/

i18n-tran:
	echo "do 之前 要有其他語系的 active.*.toml, touch internationalization/active.en.toml"
	cd internationalization && goi18n merge -sourceLanguage zh active.*.toml && cd -;
	echo "create translate.*.toml"

i18n-done:
	cd internationalization && goi18n merge active.*.toml translate.*.toml && cd -;
	echo ".... i18n-done.. merge translate to active"

i18n-clean:
	cd internationalization && rm translate.*.toml
	echo "remove translate.*.toml"

dao-model:
	cd shop-payment/dao_model && reverse -f gen_struct.yml && cd -;

.PHONY: all build fmt clean prepare run test init upgrade
$(VERBOSE).SILENT:
