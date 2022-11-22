OUTPUT_DIR=./bin
CMD_DIR=./cmd

build: network battery debug

debug:
	go build -o ${OUTPUT_DIR}/debug ${CMD_DIR}/debug/*

battery:
	go build -o ${OUTPUT_DIR}/battery ${CMD_DIR}/battery/*

network:
	go build -o ${OUTPUT_DIR}/network ${CMD_DIR}/network/*