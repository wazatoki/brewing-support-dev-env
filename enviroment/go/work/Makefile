ROOT_DIR = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: clean
clean:
	rm -rf build

.PHONY: test
test:
	cd brewing_support && \
	go test ./... && \
	cd $(ROOT_DIR)

.PHONY: build
build: clean
	mkdir build && \
	cd brewing_support && \
	export GOOS=linux && \
	go build -a -tags netgo -installsuffix netgo -extldflags "-static" -o ../build/brewing_support_linux_amd64.bin . && \
	export GOOS=windows && \
	go build -a -tags netgo -installsuffix netgo -extldflags "-static" -o ../build/brewing_support_windows_amd64.exe . && \
	cd $(ROOT_DIR) && \
	cp -r resources build/

.PHONY: migrateUP
migrateUP:
	migrate -path resources/migrations -database postgres://brewing_support:tmp_project_name@tmp_project_name_db:5432/tmp_project_namedb?sslmode=disable up

.PHONY: migrateDOWN
migrateDOWN:
	migrate -path resources/migrations -database postgres://brewing_support:tmp_project_name@tmp_project_name_db:5432/tmp_project_namedb?sslmode=disable down

.PHONY: migrateCreate
migrateCreate:
	ARG=""
	migrate create -ext sql -dir migrations -seq ${ARG}