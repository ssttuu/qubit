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

CREATE TABLE scenes (
    id BIGSERIAL PRIMARY KEY,
    project_id BIGINT NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id)
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

CREATE TABLE operators (
    id BIGSERIAL PRIMARY KEY,
    scene_id BIGINT NOT NULL,
    context VARCHAR(10) NOT NULL,
    type VARCHAR(50) NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    parameter_root JSON NOT NULL,
    FOREIGN KEY (scene_id) REFERENCES scenes(id)
);

CREATE TABLE operator_inputs (
    id BIGSERIAL PRIMARY KEY,
    operator_id BIGINT NOT NULL,
    input_id BIGINT NOT NULL,
    FOREIGN KEY (operator_id) REFERENCES operators(id),
    FOREIGN KEY (input_id) REFERENCES operators(id)
);
