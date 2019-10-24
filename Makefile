NAME = radio
GO = go
GO_OPTS =
DEFAULT = mac

# Logging helpers
log_color = \033[34m
log_name = $(NAME)
log_no_color = \033[0m
m = printf "$(log_color)$(log_name)$(log_no_color) %s$(log_no_color)\n"

GO_FILES = $(shell find . -name "*.go" ! -path './tools/*')

# Helper for building go executables
GO_BUILD = \
	if [ "$$(basename $$(dirname $@))" = "mac" ]; then \
		$m "Building "`basename $@`" for Mac..."; \
	 	mkdir -p `dirname $@`; \
	 	env GOOS=darwin GOARCH=amd64 $(GO) build -o $@ $(GO_OPTS) $<; \
	else \
		$m "Building "`basename $@`" for RPi..."; \
	 	mkdir -p `dirname $@`; \
		env GOOS=linux GOARCH=arm GOARM=5 $(GO) build -o $@ $(GO_OPTS) $<; \
	fi

# All the executable names
exec = bin/_/radio bin/_/hex

# Build the whole mac bundle
mac: $(subst _,mac,$(exec))

# Build the whole rpi bundle
rpi: $(subst _,rpi,$(exec))

radio: bin/$(DEFAULT)/radio
hex: bin/$(DEFAULT)/hex

# Build radio executable
bin/%/$(NAME): cmd/main.go $(GO_FILES) sounds/bin.go
	@$(GO_BUILD)

# Build the hex helper tool
bin/%/hex: tools/hex/main.go
	@$(GO_BUILD)

# Pack the system sounds
sounds/bin.go: $(wildcard sounds/wav/*)
	@$m "Packing sounds/wav..."
	@go-bindata -nocompress -o sounds/bin.go $<

# Clean up build artefacts
clean:
	@$m "Cleaning..."
	@rm -rf bin/*
