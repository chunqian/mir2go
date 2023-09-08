
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
	build-LogDataServer

clear-all: \
	clear-LogDataServer

# ******************** LogDataServer ********************
build-LogDataServer:
	$(call build-project, server/LogServer, source/LogDataServer, LogDataServer)

clear-LogDataServer:
	$(call clear-project, server/LogServer, LogDataServer)

run-LogDataServer:
	$(call run-project, server/LogServer, LogDataServer)

# ******************** LoginGate ********************
build-LoginGate:
	$(call build-project, server/LoginGate, source/LoginGate, LoginGate)

clear-LoginGate:
	$(call clear-project, server/LoginGate, LoginGate)

run-LoginGate:
	$(call run-project, server/LoginGate, LoginGate)
