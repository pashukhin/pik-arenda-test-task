create table "free_schedule" (
    "id" serial unique,
    "start"  timestamp not null,
    "end"  timestamp not null,
    "value" integer not null
);