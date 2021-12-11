NAME=vcf

.PHONY: clean

windows-amd64:
	GOARCH=amd64 GOOS=windows go build -trimpath -ldflags '-w -s' -o build/$@/$(NAME).exe ${MAIN_ENTRY}

clean:
	rm -rf ./build