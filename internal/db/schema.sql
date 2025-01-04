-- World table
CREATE TABLE worlds (
    id SERIAL PRIMARY KEY,
    kind VARCHAR(50) NOT NULL
);

-- Character table
CREATE TABLE characters (
    id SERIAL PRIMARY KEY,
    health INT NOT NULL,
    experience INT NOT NULL,
    level INT NOT NULL,
    food INT NOT NULL,
    water INT NOT NULL,
    energy INT NOT NULL
);

-- User table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    age INT,
    email VARCHAR(150) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    character_id INT,
    world_id INT,
    FOREIGN KEY (character_id) REFERENCES characters(id) ON DELETE CASCADE,
    FOREIGN KEY (world_id) REFERENCES worlds(id) ON DELETE CASCADE
);

-- NonPlayableCharacter table
CREATE TABLE non_playable_characters (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    health INT NOT NULL,
    experience INT NOT NULL,
    level INT NOT NULL,
    food INT NOT NULL,
    water INT NOT NULL,
    energy INT NOT NULL,
    relation INT NOT NULL,
    world_id INT NOT NULL,
    FOREIGN KEY (world_id) REFERENCES worlds(id) ON DELETE CASCADE
);

-- Habit table
CREATE TABLE habits (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    positive BOOLEAN NOT NULL,
    counter INT NOT NULL
);

-- UserHabits relational table (many-to-many between users and habits)
CREATE TABLE user_habits (
    user_id INT NOT NULL,
    habit_id INT NOT NULL,
    PRIMARY KEY (user_id, habit_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (habit_id) REFERENCES habits(id) ON DELETE CASCADE
);

-- Item table
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    itype VARCHAR(50) NOT NULL,
    rarity VARCHAR(50) NOT NULL,
    min_level INT NOT NULL,
    spawn_rate FLOAT NOT NULL
);

-- CharacterItems relational table (many-to-many between characters and items)
CREATE TABLE character_items (
    character_id INT NOT NULL,
    item_id INT NOT NULL,
    PRIMARY KEY (character_id, item_id),
    FOREIGN KEY (character_id) REFERENCES characters(id) ON DELETE CASCADE,
    FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE
);

-- Effects table
CREATE TABLE effects (
    id SERIAL PRIMARY KEY,
    effect_type VARCHAR(50) NOT NULL,
    value INT NOT NULL,
    duration INT NOT NULL
);

-- ItemEffects relational table (many-to-many between items and effects)
CREATE TABLE item_effects (
    item_id INT NOT NULL,
    effect_id INT NOT NULL,
    PRIMARY KEY (item_id, effect_id),
    FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE,
    FOREIGN KEY (effect_id) REFERENCES effects(id) ON DELETE CASCADE
);

-- NonPlayableCharacterItems relational table (many-to-many between NPCs and items)
CREATE TABLE npc_items (
    npc_id INT NOT NULL,
    item_id INT NOT NULL,
    PRIMARY KEY (npc_id, item_id),
    FOREIGN KEY (npc_id) REFERENCES non_playable_characters(id) ON DELETE CASCADE,
    FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE
);
