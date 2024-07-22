-- +goose Up
-- +goose StatementBegin
create table chats(
    id         serial primary key,
    chat_name  varchar(255) not null unique,
    created_at timestamp default now() not null
);

create table chat_users(
    chat_id    int not null,
    user_id    int not null,
    user_name  varchar(255) not null,
    created_at timestamp default now() not null,
    primary key (chat_id, user_id)
);

alter table chat_users add foreign key (chat_id) references chats(id) on delete cascade;
create index in_chat_users_chat_id on chat_users(chat_id);

create table chat_messages (
    id serial  primary key,
    chat_id    int not null,
    user_id    int not null,
    text       varchar(2000) not null,
    created_at timestamp default now() not null
);

alter table chat_messages add foreign key (chat_id) references chats (id) on delete cascade;
alter table chat_messages add foreign key (chat_id, user_id) references chat_users (chat_id, user_id);
create index in_chat_messages_user_id on chat_messages(chat_id, user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table chat_messages;
drop table chat_users;
drop table chats;
-- +goose StatementEnd
