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
    version BIGINT NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    operators JSON NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id)
);
CREATE TABLE scene_logs (
    id BIGSERIAL PRIMARY KEY,
    scene_id BIGINT NOT NULL,
    scene_version BIGINT NOT NULL,
    previous_log_id BIGINT NOT NULL,
    change_data JSON NOT NULL,
    FOREIGN KEY (scene_id) REFERENCES scenes(id),
    FOREIGN KEY (previous_log_id) REFERENCES scene_logs(id),
    UNIQUE (scene_id, scene_version)
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
CREATE TABLE render_parameters (
    id BIGSERIAL PRIMARY KEY,
    scene_id BIGINT NOT NULL,
    scene_version BIGINT NOT NULL,
    operator_id BIGINT NOT NULL,
    time DOUBLE PRECISION NOT NULL, -- TODO: BIGINT?
    input_ids BIGINT[] NOT NULL,
    configuration JSON NOT NULL
);
CREATE INDEX idx_render_parameters__scene_id__scene_version__operator_id__time
    ON render_parameters
    USING btree (scene_id, scene_version, operator_id, time);
--------------------------------
-- End Denormalized Operators
--------------------------------
