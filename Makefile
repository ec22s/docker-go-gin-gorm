SHELL=/bin/bash

comp=docker compose
exec=${comp} exec
go_service=go
db_service=mysql
db_command=${db_service} mysql --defaults-extra-file=/run/secrets/mysql_secret first
db_exec=${exec} -T ${db_command}
fqdn_host=http://localhost:8008
curl_prefix=curl -X POST ${fqdn_host}/api/
curl_suffix=-H "Content-Type: application/json" -d
register=${curl_prefix}register ${curl_suffix}
login=${curl_prefix}login ${curl_suffix}
username=yourname
password=yoursecret

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

into-go:
	@${exec} ${go_service} bash

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
