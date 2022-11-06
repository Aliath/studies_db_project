# create backup
pg_dump postgresql://postgres:164197@127.0.0.1:4200/studies_db_project > backup_$(date +%F_%H-%M-%S).sql