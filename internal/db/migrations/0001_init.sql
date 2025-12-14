DROP TABLE IF EXISTS objects;

CREATE TABLE objects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT,
    year TEXT,
    type TEXT,
    course TEXT,
    study_form TEXT,
    exploration_note TEXT,
    goals TEXT,
    prerequisites TEXT,
    postrequisites TEXT,
    group_id JSONB,
    contacts JSONB,
    internet_resources JSONB,
    technical_resources JSONB,
    rup_type TEXT,
    disciplines JSONB,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
