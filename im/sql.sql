CREATE TABLE IF NOT EXISTS `im_user` (
  `uid` VARCHAR(128) NOT NULL,
  `type` INT(1) NOT NULL DEFAULT 1,
  `ts` timestamp default current_timestamp,
  PRIMARY KEY (`uid`)
);

create table if not exists im_friends(
  uid VARCHAR(128) not null,
  fid VARCHAR(128) not null,
  `ts` timestamp default current_timestamp,
  index (uid),
  PRIMARY KEY (`uid`, `fid`),
  foreign key (uid) references im_user(uid) on delete cascade
);

CREATE TABLE if not exists `im_room` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(1024) NOT NULL,
  `uid` VARCHAR(128) NOT NULL,
  `ts` timestamp default current_timestamp,
  PRIMARY KEY (`id`),
  foreign key (uid) references im_user(uid) on delete cascade
);

CREATE TABLE if not exists `im_room_users` (
  `uid` VARCHAR(128) NOT NULL,
  `rid` INT NOT NULL,
  `ts` timestamp default current_timestamp,
  PRIMARY KEY (`uid`, `rid`),
  foreign key (uid) references im_user(uid) on delete cascade,
  foreign key (rid) references im_room(id) on delete cascade
);
