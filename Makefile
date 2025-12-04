SHELL=/bin/bash

username=yourname1
password=yoursecret
token=
dummy_token=
HOST_PORT=8008
DB_HOST=mysql
DB_NAME=first

comp=docker compose
exec=${comp} exec
logs=${comp} logs
go_service=go
db_service=mysql
db_command=${db_service} ${DB_HOST} --defaults-extra-file=/run/secrets/mysql_secret ${DB_NAME}
db_exec=${exec} -T ${db_command}
fqdn_host=http://localhost:${HOST_PORT}
curl_prefix=curl -i -X POST ${fqdn_host}/api/
curl_suffix=-H "Content-Type: application/json" -d
register=${curl_prefix}register ${curl_suffix}
login=${curl_prefix}login ${curl_suffix}
api=curl -i ${fqdn_host}/api/admin/user

.PHONY: $(shell egrep -oh ^[a-zA-Z0-9][a-zA-Z0-9_-]+: $(MAKEFILE_LIST) | sed 's/://')

up:
	@${comp} up -d --build

down:
	@${comp} down -v

restart-go:
	@${comp} restart ${go_service}

clean-restart:
	@make down && make up

test-root:
	@curl ${fqdn_host} && echo

test-register:
	@${register} '{"username": "${username}", "password": "${password}"}' && echo

err-register-1:
	@${register} '{}' && echo

err-register-2:
	@${register} '{"username": "${username}"}' && echo

test-login:
	@${login} '{"username": "${username}", "password": "${password}"}' && echo

err-login-1:
	@${login} '{"username": "${username}", "password": "wrongPassword"}' && echo

err-login-2:
	@${login} '{"username": "wrongUserName", "password": "${password}"}' && echo

test-api:
	@${api} -H "Authorization: Bearer ${token}" && echo

err-api-1:
	@${api} && echo

err-api-2:
	@${api} -H "Authorization: Bearer ${dummy_token}" && echo

into-go:
	@${exec} ${go_service} bash

logs-go:
	@${logs} -f ${go_service}

into-db:
	@${exec} ${db_command}

tables:
	@echo "show tables;" | ${db_exec}
# TODO: wait for migration

table-detail:
	@echo "describe users;" | ${db_exec}
# TODO: wait for migration

table-rows:
	@echo "SELECT * FROM users;" | ${db_exec}
# TODO: wait for migration

test-db-insert:
	@echo "INSERT INTO users (username, password) VALUES ('test', 'test');" | ${db_exec}
# TODO: wait for migration
