APP_NAME := sun
BINARY_PATH := ./$(APP_NAME)
INSTALL_PATH := /usr/local/bin/$(APP_NAME)

build:
	@echo "Building the Go project..."
	go build -o $(BINARY_PATH) .
	@echo "Installing $(APP_NAME) to $(INSTALL_PATH)..."
	@sudo mv $(BINARY_PATH) $(INSTALL_PATH)
	@echo "$(APP_NAME) installed successfully."

clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_PATH)
	@echo "Clean up completed."

.PHONY: build clean
