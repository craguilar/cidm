-- noinspection SqlNoDataSourceInspectionForFile
-----------------------------------------------------------------------------------------------------------------------
--- Cleanup data base
-----------------------------------------------------------------------------------------------------------------------
set search_path = t_cidm;
-- Cleanup database Users
drop sequence s_users_id ;
drop table    user_tokens;
drop table    users;
-- Cleanup Authentication
drop sequence s_authentication_id;
drop table    logins;
drop table    authentications;
drop table    resources;
drop table    user_resources;
drop table    configurations;
drop schema   t_cidm;
commit;
-----------------------------------------------------------------------------------------------------------------------
--- Create data base
-----------------------------------------------------------------------------------------------------------------------

-- Create  DB Schemas
create schema t_cidm;
commit;
--- Sequence creation
create sequence s_authentication_id START 100;
create sequence s_users_id START 1001;

-- CREATE TABLES
CREATE TABLE users (
	id        varchar(255),
	created_on   date,
	created_by   varchar(255),
	constraint pk_users primary key (id)
);
comment on table users is 'TABLE users, needed to store  email for particular user';

create table user_tokens(
	id 			 varchar(2000), -- this is the access token
	user_id      varchar(255)  references users(id),-- fk for users
	token_type   varchar(1000),-- oauth, user/password
	expiration   date,         -- expiration
	-- Control attributes
	created_on   date,
	created_by   varchar(255),
	constraint pk_user_tokens primary key (id,user_id)
);
comment on table user_tokens is 'TABLE user_tokens associated with a particular user.';

create table logins(
	id    NUMERIC,
	title VARCHAR(1000) NOT NULL UNIQUE,
	is_default BOOLEAN DEFAULT false,
	TYPE       VARCHAR(255),
	enabled    BOOLEAN ,
	canonical_url VARCHAR(512),
	-- Control attributes
	modified_on  date,
	modified_by  varchar(255),
	created_on   date,
	created_by   varchar(255),
	constraint pk_logins primary key (id)
);
comment on table logins is 'TABLE logins, store available logins for the idm service.';

create table authentications(
	id numeric,
	authentication_type varchar(1000),    -- oauth, user/password
	authentication_provider varchar(1000),-- google , git-hub , etc.
	login_url varchar(360)           ,    -- /api/v1/idm/cidmlogin, etc..
	redirect_url varchar(360)        ,    -- /api/v1/idm/cidmcallback, logout
	oauth_client_id     varchar(1000),  -- 655418251283-g86mvkobac60g8jtrgrl6hv3q096fn2s.apps.googleusercontent.com
	oauth_client_secret varchar(1000),  -- lckknE8RFegCRZep85iIizpR
	enabled             boolean,
	constraint pk_authentication primary key(id)
);
comment on table logins is 'TABLE authentications, configurations for instance.';

create table configurations(
    id varchar(2000),
    value varchar(2000),
    enabled boolean,
	modified_on  date,
	modified_by  varchar(255),
	created_on   date,
	created_by   varchar(255),
	constraint pk_configuration primary key (id)
);

create table resources(
	id varchar(1000),
	description varchar(1000),
	modified_on  date,
	modified_by  varchar(255),
	created_on   date,
	created_by   varchar(255),
	constraint pk_resources primary key(id)
);


create table permissions(
  email        varchar(255),
  resource     varchar(1000),
  operation    varchar(10),
  enabled      boolean,
  modified_on  date,
  modified_by  varchar(255),
  created_on   date,
  created_by   varchar(255),
  constraint pk_permissions primary key (email,resource,operation)
);

-- Index creation creation
create unique index idx_login_title  on logins(title);
commit;


-----------------------------------------------------------------------------------------------------------------------
--- Initial setups for IDM 
--- 1 is always created by Admin
-----------------------------------------------------------------------------------------------------------------------
insert into users 
values('caruizag@gmail.com',current_timestamp,0);

insert into logins (id,title,is_default,"type",enabled,canonical_url,created_on,created_by)
values (1,'Identity Management',true,'oauth2-google',true,'/api/idm/cidmlogin',current_timestamp,'caruizag@gmail.com') ;

insert into authentications 
values(1,'OAUTH2','GOOGLE','/api/v1/idm/cidmlogin','/api/idm/cidmcallback',
      '655418251283-g86mvkobac60g8jtrgrl6hv3q096fn2s.apps.googleusercontent.com','lckknE8RFegCRZep85iIizpR',true);

insert into configurations(id,value,enabled,modified_on,modified_by,created_on,created_by)
values ('oauth2callbackurl','http://localhost:3000/dashboard',true,null,null,current_timestamp,'caruizag@gmail.com');

insert into resources (id,description,modified_on,modified_by,created_on,created_by)
values ('*','All resources',null,null,current_timestamp,':admin_user');

insert into permissions(email,resource,operation,modified_on,modified_by,created_on,created_by)
values ('caruizag@gmail.com','*','CRUD',null,null,current_timestamp,'caruizag@gmail.com');

commit;


