EXAMPLES := \
	basic \
	cellar \
	cookies \
	encodings/cbor \
	encodings/text \
	error \
	files \
	multipart \
	security/hierarchy \
	security/multiauth \
	streaming \
	tus \
	upload_download \
	# -- end

GEN_TARGETS := ${EXAMPLES:%=gen-%}

GOOS = $(shell go env GOOS)
EXE =

ifeq (${GOOS},windows)
	EXE = .exe
endif

.PHONY: all
all: ${EXAMPLES}

.PHONY: ${EXAMPLES}
${EXAMPLES}: deps
	@echo "run $@..."
	@go run ../cmd/goalint goa.design/examples/$@/design

.PHONY: gen
gen: ${GEN_TARGETS}

.PHONY: ${GEN_TARGETS}
${GEN_TARGETS}: deps
	@echo "generate ${@:gen-%=%}..."
	@PATH="${PWD}/bin:$${PATH}" go run goa.design/goa/v3/cmd/goa gen goa.design/examples/${@:gen-%=%}/design -o ${@:gen-%=%}

.PHONY: deps
deps: protoc
	go mod tidy

.PHONY: protoc
protoc: bin/protoc

# Install protoc
# https://github.com/goadesign/examples/blob/d560275d3ab0560dc7943841d329fda46efcd18c/Makefile#L39
PROTOC_VERSION = 23.4
UNZIP = unzip
ifeq (${GOOS},linux)
	PROTOC = protoc-${PROTOC_VERSION}-linux-x86_64
endif
ifeq (${GOOS},darwin)
	PROTOC = protoc-${PROTOC_VERSION}-osx-x86_64
endif
ifeq (${GOOS},windows)
	PROTOC = protoc-${PROTOC_VERSION}-win32
endif

bin/protoc:
	@echo "installing protoc..."
	@mkdir -p _protoc;
	@cd _protoc; \
	curl -O -L https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC}.zip; \
	${UNZIP} -o ${PROTOC}.zip bin/protoc${EXE}
	@cp _protoc/bin/protoc${EXE} $@
	@rm -rf _protoc
