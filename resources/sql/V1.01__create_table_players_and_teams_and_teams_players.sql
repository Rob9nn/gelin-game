create table players (
	player_id serial primary key,
	firstname text not null,
	lastname text not null,
	phone integer not null,
	created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp
);

create table teams (
	team_id serial primary key,
	name text not null,
	created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp
);

create table teams_players (
	player_id integer not null,
	team_id integer not null,
	foreign key (player_id) references players(player_id),
	foreign key (team_id) references teams(team_id),
	created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp
);
