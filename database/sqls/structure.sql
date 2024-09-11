create TABLE users (
  id bigint  primary key generated always as identity,
  name text,
  premium bool default false,
  email text,
  active bool default true,
  created_at timestamp with time zone default now()
);

create table tags (
  id bigint primary key generated always as identity,
  name text,
  public boolean
);

CREATE TABLE IF NOT EXISTS qrs (
  id bigint primary key generated always as identity,
  qr_code text,
  url_text text,
  premium bool DEFAULT false null,
  id_tag bigint references tags (id) ON DELETE SET NULL,
  id_user bigint references users (id) ON DELETE SET NULL,
  created_at timestamp DEFAULT now()
);

