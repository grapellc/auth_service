# Auth service migrations

- **001_auth_tables.sql** – Initial schema (integer IDs). Required.
- **002_auth_int_to_uuid.sql** – Migrates existing data from integer to UUID primary keys.

Do **not** add a file named `001_auth_tables_uuid.sql`. It would conflict with `001_auth_tables.sql` (duplicate goose version 1) and cause startup panic. If you see that file in a built image, rebuild from a commit that only has 001 and 002 above.
