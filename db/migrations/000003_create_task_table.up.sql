create table "task" (
    "id" serial unique,
    "worker_id" integer references "worker"("id") on delete set null,
    "start"  timestamp not null,
    "end"  timestamp not null,
    "cancelled" boolean default FALSE
);