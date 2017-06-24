CREATE TABLE computations (
    id BIGSERIAL PRIMARY KEY,
    status VARCHAR(20) NOT NULL,
    root_operator_id BIGINT NOT NULL,
    operator_map JSON NOT NULL,
    resource_id VARCHAR(255) NOT NULL
);
