CREATE TABLE organizations (
    id BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL
);
CREATE TABLE projects (
    id BIGSERIAL PRIMARY KEY,
    organization_id BIGINT NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    FOREIGN KEY (organization_id) REFERENCES organizations(id)
);

CREATE TABLE image_sequences (
    id BIGSERIAL PRIMARY KEY,
    project_id BIGINT NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id)
);
CREATE TABLE images (
    id BIGSERIAL PRIMARY KEY,
    image_sequence_id BIGINT NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    FOREIGN KEY (image_sequence_id) REFERENCES image_sequences(id)
);


CREATE TABLE scene_events (
    id BIGSERIAL PRIMARY KEY,
    scene_id BIGINT NOT NULL,
    down_version INT,
    down_change_data JSON NOT NULL,
    up_version INT NOT NULL,
    up_change_data JSON NOT NULL,
    FOREIGN KEY (scene_id) REFERENCES scenes(id),
    -- TODO: Should change data be included in these indices? I think it will blow
    -- TODO: out the index sizes too much to be worth it.
    -- SELECT down_version, down_change_data WHERE scene_id=$1 ORDER BY down_version DESC
    UNIQUE (scene_id, down_version),
    -- SELECT up_version, up_change_data WHERE scene_id=$1 ORDER BY up_version ASC
    UNIQUE (scene_id, up_version)
);
CREATE TABLE scene_snapshots (
    id BIGSERIAL PRIMARY KEY,
    scene_id BIGINT NOT NULL,
    version INT NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    operators JSON NOT NULL,
    FOREIGN KEY (scene_id) REFERENCES scenes(id),
    UNIQUE (scene_id, version)
);

CREATE TABLE scenes (
    id BIGSERIAL PRIMARY KEY,
    project_id BIGINT NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id)
);

CREATE TABLE operators (
    id BIGSERIAL PRIMARY KEY,
    scene_id BIGINT NOT NULL,
    context VARCHAR(100) NOT NULL,
    input_ids BIGINT ARRAY,
    "name" VARCHAR(2048) NOT NULL,
    "type" VARCHAR(100) NOT NULL,
    parameters JSON NOT NULL,
    FOREIGN KEY (scene_id) REFERENCES scenes(id)
);

create TABLE connections (
    id BIGSERIAL PRIMARY KEY,
    input_id BIGINT NOT NULL,
    input_index INT NOT NULL,
    output_id BIGINT NOT NULL,
    output_index INT NOT NULL,
    FOREIGN KEY (input_id) REFERENCES operators(id),
    FOREIGN KEY (output_id) REFERENCES operators(id)
);
