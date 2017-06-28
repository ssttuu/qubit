CREATE TABLE computations (
    id BIGSERIAL PRIMARY KEY,
    operator_key VARCHAR(255) NOT NULL,
    time DOUBLE PRECISION NOT NULL,
    bounding_box_2d JSON NOT NULL,
    resource_id VARCHAR(255) NOT NULL
);

CREATE TABLE computation_statuses (
    id BIGSERIAL PRIMARY KEY,
    computation_id BIGINT NOT NULL,
    status SMALLINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    FOREIGN KEY (computation_id) REFERENCES computations(id)
);

CREATE INDEX idx_computation_statuses__computation_id__created_at__status
    ON computation_statuses
    USING btree (computation_id, created_at, status);
