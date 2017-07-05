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

--------------------------------
-- Begin Normalized Operators
--------------------------------
CREATE TABLE scenes (
    id BIGSERIAL PRIMARY KEY,
    project_id BIGINT NOT NULL,
    version INT NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    operators JSON NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id)
);
CREATE TABLE scene_events (
    id BIGSERIAL PRIMARY KEY,
    scene_id BIGINT NOT NULL,
    down_version INT NOT NULL,
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
--------------------------------
-- End Normalized Operators
--------------------------------

--------------------------------
-- Begin Denormalized Operators
--------------------------------
-- Denormalize and duplicate data from operators, to
-- read-optimize for fast response-time to compute to render.
-- Not using foreign keys to normalized operator tables,
-- because the whole point is to query independently.
--------------------------------
CREATE TABLE renders_operators (
    id BIGSERIAL PRIMARY KEY,
    scene_id BIGINT NOT NULL,
    scene_version INT NOT NULL,
    -- These only must be unique in the context of a scene; could probably be a UUID
    operator_id VARCHAR(32) NOT NULL,
    time BIGINT NOT NULL, -- Using BIGINT because the value needs to exactly match
    -- TODO: Bounding box
    start_x INT NOT NULL,
    start_y INT NOT NULL,
    stop_x INT NOT NULL,
    stop_y INT NOT NULL,
    input_ids VARCHAR(32)[] NOT NULL,
    parameters JSON NOT NULL
);
-- SELECT input_ids, parameters FROM render_operators
-- WHERE scene_id = $1 AND scene_version = $2 AND operator_id = $3 AND time = $4
--  AND start_x <= $5 AND start_y <= $5 AND stop_x >= $6 AND stop_y >= $6
-- Constraint names have a max length of 63 bytes, after which it is truncated
CREATE INDEX idx_render_operators__scene_id__scene_version__operator_id__tim
    ON renders_operators
    USING btree (scene_id, scene_version, operator_id, time, start_x, start_y, stop_x, stop_y);
--------------------------------
-- End Denormalized Operators
--------------------------------
