NAME = radio
GO = go

# Start pulseaudio
pulse:
	pulseaudio --disallow-module-loading --disallow-exit --fail

# Build executable for mac
bin/mac/$(NAME):
	@echo "Building for Mac..."
	@mkdir -p bin/mac
	@GOOS=darwin GOARCH=amd64 $(GO) build -o bin/mac/$(NAME) cmd/main.go

# Build executable for raspberry pi
bin/rpi/$(NAME):
	@echo "Building for RPi..."
	@mkdir -p bin/rpi
	@GOOS=linux GOARCH=amd64 $(GO) build -o bin/rpi/$(NAME) cmd/main.go

# Clean up build artefacts
clean:
	@echo "Cleaning..."
	@rm -rf bin/*
