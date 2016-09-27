create table requests (
  id varchar(16) not null primary key, -- default uuid()
  request_ts timestamp not null default current_timestamp,
  request_data LONGTEXT,
  response_ts timestamp null,
  response_data LONGTEXT null
);
