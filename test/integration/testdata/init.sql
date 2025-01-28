-- test/integration/testdata/init.sql

CREATE TABLE IF NOT EXISTS players (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    sport TEXT NOT NULL,
    team TEXT NOT NULL,
    profile_image_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS player_descriptions (
    id UUID PRIMARY KEY,
    player_id UUID REFERENCES players(id),
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS media (
    id UUID PRIMARY KEY,
    player_id UUID REFERENCES players(id),
    source TEXT NOT NULL,
    url TEXT NOT NULL,
    title TEXT NOT NULL,
    content TEXT,
    published_at TIMESTAMP WITH TIME ZONE NOT NULL,
    thumbnail_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);