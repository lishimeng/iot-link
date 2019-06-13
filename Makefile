
build: ui-requirements ui/build
        go build main.go

clean:
        @echo "Cleaning up workspace"
        @rm -rf iot-link
        @rm -rf iot-link.exe
        @rm -rf static/*

ui-requirements:
        @echo "Installing UI requirements"
        @cd ui && npm install

ui/build:
        @echo "Build UI"
        @cd ui && npm run build:prod
        @mv ui/dist/* static