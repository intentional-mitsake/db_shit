# db-shit 
*Database backup & restore CLI*

`db-shit` wraps native database tools (`pg_dump` / `pg_restore`) to backup and restore databases including automatic database creation if it doesn’t exist.


This was made purely for learning purposes; so no re-inventing the wheel.

---

## Features

- Backup & restore PostgreSQL databases
- Auto-create database if missing
- YAML config(needs work) + CLI flags
- Structured logging and error handling
- Encryption: will work on it
- Planned support for MySQL & MongoDB 

---

## Usage

### Backup
```bash
db-shit backup -u postgres -p password -H localhost -P 5432 -d mydb -D backups
