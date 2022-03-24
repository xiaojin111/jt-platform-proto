#编译协议文件
.ONESHELL:
protos:
	@echo "编译协议文件:$(file).proto";
	protoc  --go_out=. \
		    --validate_out="lang=go:." \
	        --micro_out=. \
			--js_out=import_style=commonjs:./js \
    		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./js \
	        --doc_out=./doc \
	        --doc_opt="html,$(file).html" \
	        -I. "$(file).proto";
	@echo "success..."
