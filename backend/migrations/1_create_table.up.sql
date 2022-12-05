CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS card (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text NOT NULL default '',
    attack integer NOT NULL default 0,
    defence integer NOT NULL default 0,
    intelligence integer NOT NULL default 0,
    agility integer NOT NULL default 0,
    resilience integer NOT NULL default 0,
    created_at timestamp with time zone NOT NULL default now(),
    updated_at timestamp with time zone NOT NULL default now()
);

CREATE TABLE IF NOT EXISTS deck (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL default now(),
    updated_at timestamp with time zone NOT NULL default now()
);

CREATE TABLE IF NOT EXISTS deck_cards (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    deck_id uuid NOT NULL references deck(id),
    card_id uuid NOT NULL references card(id),
    created_at timestamp with time zone NOT NULL default now(),
    updated_at timestamp with time zone NOT NULL default now()
);

CREATE TABLE IF NOT EXISTS match (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    player_deck_id uuid NOT NULL references deck(id),
    npc_deck_id uuid NOT NULL references deck(id),
    victory boolean default false,
    created_at timestamp with time zone NOT NULL default now(),
    updated_at timestamp with time zone NOT NULL default now()
);

CREATE TYPE attribute AS ENUM (
    'ATTACK',
    'DEFENCE',
    'INTELLIGENCE',
    'AGILITY',
    'RESILIENCE'
);

CREATE TABLE IF NOT EXISTS round (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    match_id uuid NOT NULL references match(id),
    player_card_id uuid NOT NULL references card(id),
    com_card_id uuid NOT NULL references card(id),
    victory boolean default false,
    attribute attribute NOT NULL,
    created_at timestamp with time zone NOT NULL default now(),
    updated_at timestamp with time zone NOT NULL default now()
);
