
.PHONY: build clear build-all clear-all

define build-project
	CGO_CFLAGS="-mmacosx-version-min=10.10" \
	CGO_LDFLAGS="-mmacosx-version-min=10.10" \
	go build -o ./$(strip $(1))/$(strip $(3)) ./$(strip $(2))/
endef

define clear-project
	cd $(strip $(1)) && rm -f $(strip $(2)) $(strip $(2)).exe
endef

define run-project
    @cd $(strip $(1)) && $(strip $(2))
endef

build-all: \
	build-logdataserver \
	build-logingate \
	build-rungate

clear-all: \
	clear-logdataserver \
	clear-logingate \
	clear-rungate

# ******************** logdataserver ********************
build-logdataserver:
	$(call build-project, bin/logdataserver, server/logdataserver, logdataserver)

clear-logdataserver:
	$(call clear-project, bin/logdataserver, logdataserver)

run-logdataserver:
	$(call run-project, bin/logdataserver, logdataserver)

# ******************** logingate ********************
build-logingate:
	$(call build-project, bin/logingate, gate/logingate, logingate)

clear-logingate:
	$(call clear-project, bin/logingate, logingate)

run-logingate:
	$(call run-project, bin/logingate, logingate)

# ******************** rungate ********************
build-rungate:
	$(call build-project, bin/rungate, gate/rungate, rungate)

clear-rungate:
	$(call clear-project, bin/rungate, rungate)

run-rungate:
	$(call run-project, bin/rungate, rungate)
