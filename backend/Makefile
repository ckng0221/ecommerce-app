include .env
export


make_migration:
	atlas migrate diff --env dev --var dev_url=$(ATLAS_DEV_DB_URL) 

# Migrate based on version	
migrate_db:
	atlas migrate apply --env dev --var url=${DB_CONNECTION_STRING} --allow-dirty

# Migrate based on gorm models
# Will ignore migration files
# For local, could use this or auto sync
migrate_db_gorm:
	atlas schema apply --env local --auto-approve --var url=${DB_CONNECTION_STRING}

clear_migrations:
	rm -r migrations/