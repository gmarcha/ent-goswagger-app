INSERT INTO "users" ("login", "id")
VALUES
    ('student', '489d7b59-41f6-431a-8e59-93983bed06b9'),
    ('tutor', '315c8d28-06b8-41f4-95e1-ffc8de966b39'),
    ('calendar', '05336b09-09d7-42f5-9e5b-9b2c5368e87d'),
    ('admin', '28afaa07-37f0-425b-9c7a-4c29f075f771');

INSERT INTO "role_users" ("role_id", "user_id")
VALUES
    ((SELECT DISTINCT "roles"."id" FROM "roles" WHERE "roles"."name" = 'tutor'), '315c8d28-06b8-41f4-95e1-ffc8de966b39'),
    ((SELECT DISTINCT "roles"."id" FROM "roles" WHERE "roles"."name" = 'calendar'), '05336b09-09d7-42f5-9e5b-9b2c5368e87d'),
    ((SELECT DISTINCT "roles"."id" FROM "roles" WHERE "roles"."name" = 'admin'), '28afaa07-37f0-425b-9c7a-4c29f075f771');

INSERT INTO "event_types" ("name", "color", "id")
VALUES
    ('test', 'test', '991092b1-e49b-4dfb-9297-cbe32fbe80d7');

INSERT INTO "events" ("name", "created_at", "start_at", "end_at", "event_type_events", "id")
VALUES
    ('test', '2022-03-08 21:07:36.439421+00', '2022-03-08 22:00:00+00', '2022-03-08 23:00:00+00', '991092b1-e49b-4dfb-9297-cbe32fbe80d7', 'd6d9a46a-8021-46b4-aba7-96d0fbd26f3a');
