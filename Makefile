
.PHONY: build clear write build-all clear-all write-all

define build-pkg
	@cd $(1) && make
endef

define clear-pkg
	@cd $(1) && make clean
endef

define build-project
	@cd $(1) && make
endef

define clear-project
	@cd $(1) && make cleardist
endef

define write-project
	@cd $(1) && fpcmake -w
endef

define run-project
	@cd $(1) && $(2)
endef

build-all: \
	build-pkg \
	build-LogDataServer

clear-all: \
	clear-pkg \
	clear-LogDataServer

write-all: \
	write-LogDataServer

# ******************** pkg ********************
build-pkg:
	$(call build-pkg, .pkg/packager/registration)
	$(call build-pkg, .pkg/components/lazutils)
	$(call build-pkg, .pkg/components/freetype)
	$(call build-pkg, .pkg/lcl)

clear-pkg:
	$(call clear-pkg, .pkg/packager/registration)
	$(call clear-pkg, .pkg/components/lazutils)
	$(call clear-pkg, .pkg/components/freetype)
	$(call clear-pkg, .pkg/lcl)

# ******************** LogDataServer ********************
build-LogDataServer:
	$(call build-project, source/LogDataServer)

clear-LogDataServer:
	$(call clear-project, source/LogDataServer)

write-LogDataServer:
	$(call write-project, source/LogDataServer)

run-LogDataServer:
	$(call run-project, source/LogDataServer, LogDataServer)
