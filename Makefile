NAME = radio
GO = go

# Logging helpers
log_color = \033[34m
log_name = $(NAME)
log_no_color = \033[0m
m = printf "$(log_color)$(log_name)$(log_no_color) %s$(log_no_color)\n"

# Start pulseaudio
pulse:
	@$m "Starting pulse audio..."
	@pulseaudio --disallow-module-loading --disallow-exit --fail

# Build executable for mac
bin/mac/$(NAME):
	@$m "Building for Mac..."
	@mkdir -p bin/mac
	@GOOS=darwin GOARCH=amd64 $(GO) build -o bin/mac/$(NAME) cmd/main.go

# Build executable for raspberry pi
bin/rpi/$(NAME):
	@$m "Building for RPi..."
	@mkdir -p bin/rpi
	@env GOOS=linux GOARCH=arm GOARM=5 $(GO) build -o bin/rpi/$(NAME) cmd/main.go

# Clean up build artefacts
clean:
	@$m "Cleaning..."
	@rm -rf bin/*

# Make hex helper
hex: bin/hex

bin/hex: tools/hex/main.go
	@$m "Building hex..."
	@go build -o bin/hex tools/hex/main.go

# Pack the system sounds
sounds/bin.go: $(wildcard sounds/wav/*)
	@$m "Packing sounds/wav..."
	@go-bindata -nocompress -o sounds/bin.go $<
