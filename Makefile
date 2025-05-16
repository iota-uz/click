OPENAPI_GENERATOR_VERSION = 7.9.0
GENERATOR_JAR = openapi-generator-cli-$(OPENAPI_GENERATOR_VERSION).jar
GENERATOR_URL = https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/$(OPENAPI_GENERATOR_VERSION)/$(GENERATOR_JAR)

CONFIG_FILE = config.json
INPUT_SPEC = swagger.json
OUTPUT_DIR = .

GENERATED_FILES_LIST = $(OUTPUT_DIR)/.openapi-generator/FILES
EXCLUDE_FILES := .gitignore README.md .openapi-generator/FILES

.PHONY: all
all: $(GENERATOR_JAR)
	java -jar $(GENERATOR_JAR) generate \
		--skip-validate-spec \
		-i $(INPUT_SPEC) \
		-g go \
		-c $(CONFIG_FILE) \
		-o $(OUTPUT_DIR) \
		-t ./go-templates \
		--global-property=apiTests=false,modelTests=false
	@echo "✅ Golang client generated in $(OUTPUT_DIR)"

$(GENERATOR_JAR):
	@if [ ! -f "$(GENERATOR_JAR)" ]; then \
		echo "❗ Downloading $(GENERATOR_JAR)..."; \
		curl -fL -o $(GENERATOR_JAR) $(GENERATOR_URL); \
	else \
		echo "✅ Found $(GENERATOR_JAR)"; \
	fi

.PHONY: clean-generated
clean-generated:
	@if [ -f $(GENERATED_FILES_LIST) ]; then \
		while IFS= read -r file; do \
			file=$$(echo "$$file" | tr -d '\r'); \
			skip=0; \
			for exclude in $(EXCLUDE_FILES); do \
				if [ "$$file" = "$$exclude" ]; then \
					skip=1; break; \
				fi; \
			done; \
			if [ $$skip -eq 1 ]; then \
				echo "⏭  Skipping $$file"; \
			elif [ -e "$$file" ]; then \
				rm -f "$$file" && echo "🗑️  Deleted $$file"; \
			else \
				echo "❓ Not found $$file"; \
			fi; \
		done < $(GENERATED_FILES_LIST); \
	else \
		echo "⚠️  No file list found at $(GENERATED_FILES_LIST)"; \
	fi
