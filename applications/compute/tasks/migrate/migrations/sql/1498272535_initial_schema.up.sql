CREATE TABLE computations (
    id BIGSERIAL PRIMARY KEY,
    root_operator_id BIGINT NOT NULL,
    operator_map JSON NOT NULL,
    resource_id VARCHAR(255) NOT NULL
);

CREATE TABLE computation_statuses (
    id BIGSERIAL PRIMARY KEY,
    computation_id BIGINT NOT NULL,
    status SMALLINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    FOREIGN KEY (computation_id) REFERENCES computations(id)
);
CREATE INDEX computation_statuses_indices ON computation_statuses USING btree (computation_id, created_at, status);
