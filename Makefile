run:
	cd ./bot-backend && app_env=dev go run .

dev:
	cd ./bot-backend && app_env=dev app_version=1.0.0 app_version=1.0.0 air

build: clean build-backend

build-backend:
	cd ./bot-backend && GOOS=linux GOARCH=amd64 go build -v -a -o build/dev/bin/app . && \
		cp -r static build/dev/bin/ && \
		cp -r templates build/dev/bin/

init:
	terraform -chdir=infra/dev init

plan:
	terraform -chdir=infra/dev plan -var-file=variables.tfvars -out=plan_outfile

apply:
	terraform -chdir=infra/dev apply --auto-approve "plan_outfile"
	go run github.com/EdgeJay/psg-navi-bot/bot-utils \
		-url="`terraform -chdir=infra/dev output -raw api_url`/init-bot" \
		-version="`terraform -chdir=infra/dev output -raw app_version`" \
		-secret="`terraform -chdir=infra/dev output -raw init_token_secret`"
	@printf "\n"

destroy:
	terraform -chdir=infra/dev destroy -var-file=variables.tfvars

clean:
	rm -rf ./bot-backend/build/dev
