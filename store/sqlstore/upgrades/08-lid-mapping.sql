-- v8 (compatible with v8+): Add tables for LID<->JID mapping
CREATE TABLE pakaiwa.lid_map (
	lid TEXT PRIMARY KEY,
	pn  TEXT UNIQUE NOT NULL
);
