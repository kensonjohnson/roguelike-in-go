set quiet := true

MAIN_PACKAGE_PATH := "."
BINARY_NAME := "rougelike-demo"

[private]
help:
    just --list --unsorted

_confirm:
    echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# Run dev server
dev:
    go run . -debug

# Run all Go test files
test:
    go test -v -race -buildvcs ./...

# Prepare and audit code, then push to Git Remote
push: tidy audit no-dirty
    git push

# Verify and Vet all Go files in project
audit:
    go mod verify
    go vet ./...

[private]
no-dirty:
    git diff --exit-code

# Run formatter and tidy over all Go files in project
tidy:
    go fmt ./...
    go mod tidy -v

# Build for current OS/Arch
build:
    # Include additional build steps, like TypeScript, SCSS or Tailwind compilation here...
    go build -o=/tmp/bin/{{ BINARY_NAME }} {{ MAIN_PACKAGE_PATH }}

# Build for current OS/Arch and run the resulting binary
run: build
    /tmp/bin/{{ BINARY_NAME }}

# Matrix build for all OS/Architectures
production-deploy: _confirm tidy audit
    GOOS=darwin GOARCH=arm64 go build -ldflags='-s' -o=/tmp/bin/macos_arm64/${BINARY_NAME} ${MAIN_PACKAGE_PATH}
    GOOS=windows GOARCH=amd64 go build -ldflags='-s' -o=/tmp/bin/windows_amd64/${BINARY_NAME} ${MAIN_PACKAGE_PATH}
    GOOS=windows GOARCH=arm go build -ldflags='-s' -o=/tmp/bin/windows_arm/${BINARY_NAME} ${MAIN_PACKAGE_PATH}
    GOOS=windows GOARCH=arm64 go build -ldflags='-s' -o=/tmp/bin/windows_arm64/${BINARY_NAME} ${MAIN_PACKAGE_PATH}
    # Include additional deployment steps here...
