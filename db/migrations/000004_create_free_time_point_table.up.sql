create table "free_time_point" (
    "id" serial unique,
    "schedule_id" integer references "schedule"("id") on delete cascade,
    "task_id" integer references "task"("id") on delete cascade,
    "point" timestamp not null,
    "value" int not null,
    check (("schedule_id" IS NULL) != ("task_id" IS NULL))
);