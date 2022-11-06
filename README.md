# studies-db-project

```bash
# provision production database
docker-compose up --detach

# provision development database
docker-compose -f docker-compose.development.yaml up --detach

# create a backup
chmod 777 .tools/shell/backup.sh && .tools/shell/backup.sh

# restore db
chmod 777 .tools/shell/restore.sh && .tools/shell/restore.sh backup_2022-11-06_18-21-27.sql

# run seed script
go run ./cmd/seed

# run script to change user fullname
go run ./cmd/change_name

```
