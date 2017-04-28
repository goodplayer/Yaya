-- Table: public.post

-- DROP TABLE public.post;

CREATE TABLE public.post
(
  id character(32) NOT NULL, -- 64 bit time + 32 bit node + 32 bit seq
  "user" character(32) NOT NULL, -- user id. related to user table
  status bigint NOT NULL DEFAULT 0, -- status
  createtime timestamp without time zone,
  updatetime timestamp without time zone,
  content character varying(65535),
  rev bigint, -- version
  parentid character(32), -- aggregation of posts
  type bigint, -- post type
  title character varying(10000), -- title 10000
  CONSTRAINT post_pri_key_id PRIMARY KEY (id)
)
WITH (
OIDS=FALSE
);
ALTER TABLE public.post
  OWNER TO meidomx;
COMMENT ON COLUMN public.post.id IS '64 bit time + 32 bit node + 32 bit seq';
COMMENT ON COLUMN public.post."user" IS 'user id. related to user table';
COMMENT ON COLUMN public.post.status IS 'status';
COMMENT ON COLUMN public.post.rev IS 'version';
COMMENT ON COLUMN public.post.parentid IS 'aggregation of posts';
COMMENT ON COLUMN public.post.type IS 'post type';
COMMENT ON COLUMN public.post.title IS 'title 10000';


-- Index: public.post_idx_user

-- DROP INDEX public.post_idx_user;

CREATE INDEX post_idx_user
  ON public.post
  USING btree
  ("user" COLLATE pg_catalog."default");

------------------------------------------------------------

-- Table: public.postmeta

-- DROP TABLE public.postmeta;

CREATE TABLE public.postmeta
(
  id character(32) NOT NULL, -- the same as id of post
  meta1 jsonb,
  CONSTRAINT postmeta_pri_key_id PRIMARY KEY (id)
)
WITH (
OIDS=FALSE
);
ALTER TABLE public.postmeta
  OWNER TO meidomx;
COMMENT ON COLUMN public.postmeta.id IS 'the same as id of post';


-- Index: public.postmeta_idx_meta1

-- DROP INDEX public.postmeta_idx_meta1;

CREATE INDEX postmeta_idx_meta1
  ON public.postmeta
  USING gin
  (meta1);

------------------------------------------------------------

-- Table: public.task

-- DROP TABLE public.task;

CREATE TABLE public.task
(
  id character(32) NOT NULL,
  type bigint,
  title character varying(1000), -- title 1000
  content character varying(5000), -- content 5000
  start timestamp without time zone, -- start time
  "end" timestamp without time zone, -- end time
  createtime timestamp without time zone,
  updatetime timestamp without time zone,
  CONSTRAINT task_pri_key_id PRIMARY KEY (id)
)
WITH (
OIDS=FALSE
);
ALTER TABLE public.task
  OWNER TO meidomx;
COMMENT ON COLUMN public.task.title IS 'title 1000';
COMMENT ON COLUMN public.task.content IS 'content 5000';
COMMENT ON COLUMN public.task.start IS 'start time';
COMMENT ON COLUMN public.task."end" IS 'end time';


-- Index: public.task_idx_end

-- DROP INDEX public.task_idx_end;

CREATE INDEX task_idx_end
  ON public.task
  USING btree
  ("end" DESC);

-- Index: public.task_idx_start

-- DROP INDEX public.task_idx_start;

CREATE INDEX task_idx_start
  ON public.task
  USING btree
  (start DESC);

------------------------------------------------------------

-- Table: public.topic

-- DROP TABLE public.topic;

CREATE TABLE public.topic
(
  id character varying(32) NOT NULL, -- topic id
  postid character varying(32), -- post or task id
  type bigint, -- 1 - post...
  categoryid character varying(32),
  createtime timestamp without time zone,
  CONSTRAINT topic_pri_key_id PRIMARY KEY (id)
)
WITH (
OIDS=FALSE
);
ALTER TABLE public.topic
  OWNER TO meidomx;
COMMENT ON COLUMN public.topic.id IS 'topic id';
COMMENT ON COLUMN public.topic.postid IS 'post or task id';
COMMENT ON COLUMN public.topic.type IS '1 - post
2 - task';


