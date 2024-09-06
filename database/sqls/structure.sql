CREATE TABLE IF NOT EXISTS qrs (
  id bigint primary key generated always as identity,
  qr_code text,
  userId text,
  url_text text,
  premium bool DEFAULT false null,
  created_at timestamp DEFAULT now()
);

create table users (
  id uuid primary key not null,
  name text,
  premium bool default false,
  email text,
  created_at timestamp with time zone default now()
);