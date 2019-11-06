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
		env GOOS=linux GOARCH=arm GOARM=6 $(GO) build -o $@ $(GO_OPTS) $<; \
	fi

# All the executable names
exec = bin/_/radio bin/_/hex bin/_/rctl bin/_/vol bin/_/sel

# Build the whole mac bundle
mac: $(subst _,mac,$(exec))

# Build the whole rpi bundle
rpi: $(subst _,rpi,$(exec))

radio: bin/$(DEFAULT)/radio
hex: bin/$(DEFAULT)/hex
rctl: bin/$(DEFAULT)/rctl
vol: bin/$(DEFAULT)/vol
sel: bin/$(DEFAULT)/sel

# Build radio executable
bin/%/$(NAME): cmd/main.go $(GO_FILES) sounds/bin.go
	@$(GO_BUILD)

# Build the hex helper tool
bin/%/hex: tools/hex/main.go
	@$(GO_BUILD)

# Build the rctl helper tool
bin/%/rctl: tools/rctl/main.go $(GO_FILES)
	@$(GO_BUILD)

# Build the gpio pin tester tool
bin/%/pin: tools/pin/main.go $(GO_FILES)
	@$(GO_BUILD)

# Build the volume tester tool
bin/%/vol: tools/vol/main.go
	@$(GO_BUILD)

# Build the selector tester tool
bin/%/sel: tools/sel/main.go
	@$(GO_BUILD)

# Pack the system sounds
sounds/bin.go: $(wildcard sounds/wav/*)
	@$m "Packing sounds/wav..."
	@go-bindata -nocompress -o sounds/bin.go $<

# Clean up build artefacts
clean:
	@$m "Cleaning..."
	@rm -rf bin/*

# Vet all code
vet:
	@$m "Go vetting..."
	@$(GO) vet ./...

# Test all code
test:
	@$m "Testing..."
	@$(GO) test ./...
