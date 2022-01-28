PROJECT_NAME := Pulumi Kind Package

PACK             := kind
ORG              := pawelprazak
PROJECT          := github.com/${ORG}/pulumi-${PACK}
NODE_MODULE_NAME := @${ORG}/pulumi-${PACK}
TF_NAME          := ${PACK}
PROVIDER_PATH    := provider
VERSION_PATH     := ${PROVIDER_PATH}/pkg/version.Version

TFGEN           := pulumi-tfgen-${PACK}
PROVIDER        := pulumi-resource-${PACK}
VERSION         := $(shell pulumictl get version)

TESTPARALLELISM := 4

WORKING_DIR     := $(shell pwd)
PULUMI_ROOT		:= /opt/pulumi

OS := $(shell uname)
EMPTY_TO_AVOID_SED :=

prepare::
	@if test -z "${NAME}"; then echo "NAME not set"; exit 1; fi
	@if test -z "${REPOSITORY}"; then echo "REPOSITORY not set"; exit 1; fi
	@if test ! -d "provider/cmd/pulumi-tfgen-x${EMPTY_TO_AVOID_SED}yz"; then "Project already prepared"; exit 1; fi
	NAME_UPPER=$$(echo ${NAME} | tr '[:lower:]' '[:upper:]');

	mv -v "provider/cmd/pulumi-tfgen-x${EMPTY_TO_AVOID_SED}yz" provider/cmd/pulumi-tfgen-${NAME}
	mv -v "provider/cmd/pulumi-resource-x${EMPTY_TO_AVOID_SED}yz" provider/cmd/pulumi-resource-${NAME}

	if [[ "${OS}" != "Darwin" ]]; then \
		sed -i 's,github.com/pulumi/pulumi-x${EMPTY_TO_AVOID_SED}yz,${REPOSITORY},g' provider/go.mod; \
		find ./ ! -path './.git/*' -type f -exec sed -i 's/[x]yz/${NAME}/g' {} \; &> /dev/null; \
		find ./ ! -path './.git/*' -type f -exec sed -i 's/[X]YZ/$${NAME_UPPER}/g' {} \; &> /dev/null; \
	fi

	# In MacOS the -i parameter needs an empty string to execute in place.
	if [[ "${OS}" == "Darwin" ]]; then \
		sed -i '' 's,github.com/pulumi/pulumi-x${EMPTY_TO_AVOID_SED}yz,${REPOSITORY},g' provider/go.mod; \
		find ./ ! -path './.git/*' -type f -exec sed -i '' 's/[x]yz/${NAME}/g' {} \; &> /dev/null; \
		find ./ ! -path './.git/*' -type f -exec sed -i '' 's/[X]YZ/$${NAME_UPPER}/g' {} \; &> /dev/null; \
	fi

.PHONY: development provider build_sdks build_nodejs build_dotnet build_go build_python cleanup

ensure::
	cd provider && go mod tidy
	cd sdk && go mod tidy

development:: install_plugins provider lint_provider build_sdks install_sdks cleanup # Build the provider & SDKs for a development environment

# Required for the codegen action that runs in pulumi/pulumi and pulumi/pulumi-terraform-bridge
build:: install_plugins provider build_sdks install_sdks
only_build:: build

tfgen:: install_plugins
	(cd provider && go build -a -o $(WORKING_DIR)/bin/${TFGEN} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${TFGEN})
	$(WORKING_DIR)/bin/${TFGEN} schema --out provider/cmd/${PROVIDER}
	(cd provider && VERSION=$(VERSION) go generate cmd/${PROVIDER}/main.go)

provider:: tfgen install_plugins # build the provider binary
	(cd provider && go build -a -o $(WORKING_DIR)/bin/${PROVIDER} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${PROVIDER})

build_sdks:: install_plugins provider build_nodejs build_python build_go build_jvm build_dotnet # build all the sdks

build_nodejs:: NODEJS_VERSION := $(shell pulumictl get version --language javascript)
build_nodejs:: install_plugins tfgen # build the node sdk
	$(WORKING_DIR)/bin/$(TFGEN) nodejs --overlays provider/overlays/nodejs --out sdk/nodejs/
	cd sdk/nodejs/ && \
        yarn install && \
        yarn run tsc && \
        cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/ && \
    	sed -i.bak -e "s/\$${VERSION}/$(NODEJS_VERSION)/g" ./bin/package.json

build_python:: PYPI_VERSION := $(shell pulumictl get version --language python)
build_python:: install_plugins tfgen # build the python sdk
	$(WORKING_DIR)/bin/$(TFGEN) python --overlays provider/overlays/python --out sdk/python/
	cd sdk/python/ && \
        cp ../../README.md . && \
        python3 setup.py clean --all 2>/dev/null && \
        rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
        sed -i.bak -e "s/\$${VERSION}/$(PYPI_VERSION)/g" -e "s/\$${PLUGIN_VERSION}/$(VERSION)/g" ./bin/setup.py && \
        rm ./bin/setup.py.bak && \
        cd ./bin && python3 setup.py build sdist

build_dotnet:: DOTNET_VERSION := $(shell pulumictl get version --language dotnet)
build_dotnet:: install_plugins tfgen # build the dotnet sdk
	pulumictl get version --language dotnet
	$(WORKING_DIR)/bin/$(TFGEN) dotnet --overlays provider/overlays/dotnet --out sdk/dotnet/
	cd sdk/dotnet/ && \
		echo "${DOTNET_VERSION}" >version.txt && \
        dotnet build /p:Version=${DOTNET_VERSION}

build_jvm:: install_plugins tfgen # build the jvm sdk
	$(WORKING_DIR)/bin/$(TFGEN) jvm --overlays provider/overlays/jvm --out sdk/jvm/
	cd $(WORKING_DIR)/sdk/jvm && \
		mkdir -p src/main/resources/pl/pawelprazak/pulumi/$(PACK) && \
		echo "${VERSION}" > src/main/resources/pl/pawelprazak/pulumi/$(PACK)/version.txt && \
		gradle -Pversion=${VERSION} build

build_go:: install_plugins tfgen # build the go sdk
	$(WORKING_DIR)/bin/$(TFGEN) go --overlays provider/overlays/go --out sdk/go/

lint_provider:: provider # lint the provider code
	cd provider && golangci-lint run -c ../.golangci.yml

cleanup:: # cleans up the temporary directories and files
	rm -rf bin
	rm -f provider/cmd/${PROVIDER}/schema.go
	rm -rf dist

help::
	@grep '^[^.#]\+:\s\+.*#' Makefile | \
 	sed "s/\(.\+\):\s*\(.*\) #\s*\(.*\)/`printf "\033[93m"`\1`printf "\033[0m"`	\3 [\2]/" | \
 	expand -t20

clean::
	rm -rf sdk/{dotnet,nodejs,go,python,jvm}

install_plugins::

install_dotnet_sdk::
	mkdir -p $(PULUMI_ROOT)/nuget
	find . -name '*.nupkg' -print -exec cp -p {} $(PULUMI_ROOT)/nuget \;

install_python_sdk::

install_go_sdk::

install_jvm_sdk::
	cd $(WORKING_DIR)/sdk/jvm && gradle -Pversion=${VERSION} publishToMavenLocal

install_nodejs_sdk::
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin

install_sdks:: install_dotnet_sdk install_python_sdk install_nodejs_sdk install_jvm_sdk

manual_provider_install::
	pulumi plugin install resource kind $(shell bin/pulumi-resource-kind -version) --server "$(PROJECT)/releases/download"

install_local::
	mkdir -p $(HOME)/.pulumi/plugins/resource-kind-v$(shell bin/pulumi-resource-kind -version)
	touch $(HOME)/.pulumi/plugins/resource-kind-v$(shell bin/pulumi-resource-kind -version).lock
	cp -v bin/pulumi-resource-kind $(HOME)/.pulumi/plugins/resource-kind-v$(shell bin/pulumi-resource-kind -version)/pulumi-resource-kind
	pulumi plugin ls | grep kind

cleanup_local::
	pulumi plugin rm resource kind $(shell bin/pulumi-resource-kind -version) --yes

test::
	pulumi login --local
	cd examples && go test -v -tags=all -parallel ${TESTPARALLELISM} -timeout 2h

tag::
	git tag -a v$(VERSION)
	git push origin v$(VERSION)

prerelease_snapshot::
	goreleaser -p 3 --rm-dist --snapshot --config .goreleaser.prerelease.yml

release_snapshot::
	goreleaser -p 3 --rm-dist --snapshot

release::
	goreleaser -p 3 --rm-dist

hack_local_deps::
	cd provider && go mod edit -replace github.com/pulumi/pulumi/pkg/v3=${HOME}/repos/vl/pulumi/pkg
	cd provider && go mod edit -replace github.com/pulumi/pulumi/sdk/v3=${HOME}/repos/vl/pulumi/sdk
	cd provider && go mod edit -replace github.com/pulumi/pulumi-terraform-bridge/v3=${HOME}/repos/vl/pulumi-terraform-bridge