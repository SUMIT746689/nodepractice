If you update db then do these steps for migration
You need to have atlas installed in your system

```
go generate ./ent
```

If you have docker then run this command
```
atlas migrate diff migration_name --dir "file://ent/migrate/migrations" --to "ent://ent/schema" --dev-url "docker://mysql/8/ent"
```

and if you don't have docker then run this command
```
atlas migrate diff changed_db --dir "file://ent/migrate/migrations" --to "ent://ent/schema" --dev-url "mysql://root:mypassword@127.0.0.1/pos_online_atlas?parseTime=true"
```

<!-- hash first then migration for mehedi  -->
# apply the migration
atlas migrate apply --dir "file://ent/migrate/migrations" --url "mysql://root:mypassword@127.0.0.1/pos_online?parseTime=true"

<!-- for create new scheme file -->
go run -mod=mod entgo.io/ent/cmd/ent new SchemaName

<!-- for seed -->
go run cmd/db_seed/main.go

<!-- for backend run -->
docker compose up --build