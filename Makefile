NAME=gpctl
PREFIX := ~/.local/bin
BUILD_DIR := ./build

BUILD_TARGET_PATH := ${BUILD_DIR}/${NAME}
INSTALL_TARGET_PATH := ${PREFIX}/${NAME}

${BUILD_DIR}/:
	mkdir -p ${BUILD_DIR}
${BUILD_TARGET_PATH}: ${BUILD_DIR}/
	go build -o ${BUILD_TARGET_PATH} main.go

build: ${BUILD_TARGET_PATH}

${PREFIX}/:
	mkdir -p ${PREFIX}
${INSTALL_TARGET_PATH}: ${PREFIX}/
	cp ${BUILD_TARGET_PATH} ${INSTALL_TARGET_PATH}

install: ${INSTALL_TARGET_PATH}
	@echo "✅ Successfully installed ${NAME} to ${INSTALL_TARGET_PATH}"
uninstall:
	rm ${INSTALL_TARGET_PATH}

test:
	go test ./...

clean:
	rm -r ${BUILD_DIR}
