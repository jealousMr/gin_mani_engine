create table task(
    id bigint(20) auto_increment,
    task_id varchar(32) not null,
    rule_id varchar(32) not null,
    operator varchar(32) default null,
    execute_state tinyint(4) unsigned not null,
    output_name varchar(128) not null,
    output_url varchar(256) not null,
    output_state tinyint(4) unsigned not null,
    create_at timestamp default CURRENT_TIMESTAMP,
    update_at timestamp default CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE (`task_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;