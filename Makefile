# ----------------------------------
# Configuration
# ----------------------------------
SWAGGER_FILE ?= ./a_record_regex.yaml
OUTPUT_DIR ?= ./output/go-sdk
GENERATOR ?= go
OPENAPI_GENERATOR_CLI ?= openapi-generator

# ----------------------------------
# Safety checks
# ----------------------------------

ifndef SWAGGER_FILE
$(error SWAGGER_FILE is not set)
endif

ifndef OUTPUT_DIR
$(error OUTPUT_DIR is not set)
endif

# ----------------------------------
# Targets
# ----------------------------------

# Default target
all: help

# Help - Show all targets
help:
	@echo "Usage:"
	@echo "  make generate      - Generate Go SDK from Swagger"
	@echo "  make clean         - Clean generated SDK (including subfolders)"
	@echo "Variables:"
	@echo "  SWAGGER_FILE       - Path to Swagger or OpenAPI spec (default: ./a_record_regex.yaml)"
	@echo "  OUTPUT_DIR         - Output directory for generated SDK (default: ./output/go-sdk)"
	@echo "Example:"
	@echo "  make generate SWAGGER_FILE=./apispec.yaml OUTPUT_DIR=./sdk"

# Generate Go SDK
generate: check-generator
	@echo "Generating Go SDK from $(SWAGGER_FILE) to $(OUTPUT_DIR)..."
	@mkdir -p $(OUTPUT_DIR)
	$(OPENAPI_GENERATOR_CLI) generate \
		-i $(SWAGGER_FILE) \
		-g $(GENERATOR) \
		-o $(OUTPUT_DIR)
	@echo "✅ SDK generated successfully at $(OUTPUT_DIR)"

# Clean generated files (including subfolders, hidden files like .travis, etc.)
clean:
	@echo "Cleaning contents inside $(OUTPUT_DIR)..."
	@if [ -d "$(OUTPUT_DIR)" ]; then \
		if [ "$$(ls -A $(OUTPUT_DIR))" ]; then \
			rm -rf $(OUTPUT_DIR)/* $(OUTPUT_DIR)/.*; \
			echo "🧹 Cleaned all contents inside $(OUTPUT_DIR)!"; \
		else \
			echo "⚠️  $(OUTPUT_DIR) is already empty."; \
		fi; \
	else \
		echo "⚠️  $(OUTPUT_DIR) does not exist. Nothing to clean."; \
	fi

# Validate if openapi-generator-cli is installed
check-generator:
	@command -v $(OPENAPI_GENERATOR_CLI) >/dev/null 2>&1 || { \
		echo >&2 "❌ openapi-generator-cli is not installed. Please install it first."; exit 1; \
	}

.PHONY: all help generate clean check-generator
