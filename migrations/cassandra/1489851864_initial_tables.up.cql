CREATE TABLE IF NOT EXISTS qubit.node (
			node_id uuid,
			node_version int,
			data text,
			digest text,
			labels map<text, text>,
			inputs list<uuid>,
			outputs list<uuid>,
			PRIMARY KEY (node_id, node_version)
) WITH CLUSTERING ORDER BY (node_version ASC);
