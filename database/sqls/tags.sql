DO $$
BEGIN

  IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'users') THEN
    CREATE TABLE IF NOT EXISTS tags (
      id uuid PRIMARY KEY NOT NULL,
      name text,
      public boolean,
      qrid bigint REFERENCES qrs(id),
      iduser uuid REFERENCES public.users(id)
    );
  ELSE
    RAISE NOTICE 'The public.users table does not exist. The tags table was not created.';
  END IF;
END $$;