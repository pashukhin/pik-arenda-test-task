create table "schedule" (
    "id" serial unique,
    "worker_id" integer not null references "worker"("id") on delete cascade,
    "start"  timestamp not null,
    "end"  timestamp not null
);