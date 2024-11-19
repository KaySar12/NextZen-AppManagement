GORELEASER=goreleaser
GORELEASERBUILD=$(GORELEASER) build
SERVICE = casaos-app-management
ARCHITECHTURE= amd64
OS=linux
VERSION=v1
BIN_PATH=build/sysroot/usr/bin
BUILD_PATH=build
CUR_DIR=$(PWD)
ARCHIVE_PATH=buildzip
ifeq ($(TAG),)
PACKAGE_NAME := $(OS)-$(ARCHITECHTURE)-nextzenos-app-management
else
PACKAGE_NAME := $(OS)-$(ARCHITECHTURE)-nextzenos-app-management-$(TAG)
endif
COMMIT_MESSAGE ?="update"
build_service:
	$(GORELEASERBUILD) --clean --snapshot -f .goreleaser.debug.yaml --id $(SERVICE)-$(ARCHITECHTURE)

package:
	 cp  $(CUR_DIR)/dist/$(SERVICE)-$(ARCHITECHTURE)_$(OS)_$(ARCHITECHTURE)_$(VERSION)/$(BIN_PATH)/$(SERVICE) $(CUR_DIR)/$(BIN_PATH) \
	 && tar -czvf $(PACKAGE_NAME).tar.gz $(BUILD_PATH)

archive_package:
	@mkdir -p $(CUR_DIR)/$(ARCHIVE_PATH)/$(CUR_TAG)
	@mv $(PACKAGE_NAME).tar.gz $(CUR_DIR)/$(ARCHIVE_PATH)/$(CUR_TAG)/
remove_package:
	rm $(PACKAGE_NAME).tar.gz
clear_archive:
	@rm -rf $(CUR_DIR)/$(ARCHIVE_PATH)
#make create_tag CUR_TAG=x.x TAG_MESSAGE="this is tag message"
create_tag:
	@${GIT} tag -a ${CUR_TAG} -m "${TAG_MESSAGE}" || { echo "Failed to create tag"; exit 1; }
	@${GIT} push ${GIT_REMOTE} ${CUR_TAG} ||  { echo "Failed to push tag"; exit 1; }
#make remove_tag CUR_TAG=x.x
remove_tag:
	@${GIT} tag -d ${CUR_TAG}
	@${GIT} push ${GIT_REMOTE} -d ${CUR_TAG}	
check_tag:
	@echo "Current tag: $(CUR_TAG)";  
push_release_all:
	${GORELEASER} release --clean  -f .goreleaser.yaml
push_release:
	${GORELEASER} release --single-target
push_git:
	@${GIT} pull ${GIT_REMOTE}
	@${GIT} add .
	@${GIT} commit -m "${COMMIT_MESSAGE}"
	@${GIT} push ${GIT_REMOTE}