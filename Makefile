OUTPUT_DIR=./bin
CMD_DIR=./cmd

build: battery debug

debug:
	go build -o ${OUTPUT_DIR}/debug ${CMD_DIR}/debug/*

battery:
	go build -o ${OUTPUT_DIR}/battery ${CMD_DIR}/battery/*