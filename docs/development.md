# Development


## DB Migration
Use [golang-migrate/migrate](https://github.com/golang-migrate/migrate) for DB migration.

Upgrade schema version.
```bash
make migrate-up limit=n #migrate up n
```
Downgrade schema version.
```bash
make migrate-down limit=n #migrate down n
```
Print current schema version.
```
make migrate-version
```
