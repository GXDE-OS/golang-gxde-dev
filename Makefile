build:
	echo Finish

install:
	mkdir -pv $(DESTDIR)/usr/share/gocode-gxde/src/
	cp -rv gocode/* $(DESTDIR)/usr/share/gocode-gxde/src/