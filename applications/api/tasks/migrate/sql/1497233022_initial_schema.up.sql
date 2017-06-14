CREATE TABLE organizations (
    id BIGINT PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL
);

CREATE TABLE projects (
    id BIGINT PRIMARY KEY,
    organization_id BIGINT NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    FOREIGN KEY (organization_id) REFERENCES organizations(id)
);

CREATE TABLE scenes (
    id BIGINT PRIMARY KEY,
    project_id BIGINT NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id)
);

CREATE TABLE image_sequences (
    id BIGINT PRIMARY KEY,
    project_id BIGINT NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id)
);

CREATE TABLE images (
    id BIGINT PRIMARY KEY,
    image_sequence_id BIGINT NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    FOREIGN KEY (image_sequence_id) REFERENCES image_sequences(id)
);

CREATE TABLE operators (
    id BIGINT PRIMARY KEY,
    scene_id BIGINT NOT NULL,
    context VARCHAR(10) NOT NULL,
    type VARCHAR(50) NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    inputs BIGINT[] NOT NULL,
    parameters JSON,
    FOREIGN KEY (scene_id) REFERENCES scenes(id)
);
