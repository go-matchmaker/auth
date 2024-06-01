
-- Create the users table
CREATE TABLE departments (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users (
    id VARCHAR(255) PRIMARY KEY,
    role VARCHAR(20) NOT NULL,
    name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone_number VARCHAR(20),
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    department_id VARCHAR(255) NOT NULL REFERENCES departments(id)
);

-- Add a constraint to check the role values
ALTER TABLE users
ADD CONSTRAINT role_check
CHECK (role IN ('super_admin', 'admin', 'employee', 'user'));


-- Create the user_attributes table
CREATE TABLE user_attributes (
    id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL REFERENCES users(id),
    attribute VARCHAR(255) NOT NULL,
    view BOOLEAN DEFAULT FALSE,
    search BOOLEAN DEFAULT FALSE,
    detail BOOLEAN DEFAULT FALSE,
    add BOOLEAN DEFAULT FALSE,
    update BOOLEAN DEFAULT FALSE,
    delete BOOLEAN DEFAULT FALSE,
    export BOOLEAN DEFAULT FALSE,
    upload BOOLEAN DEFAULT FALSE,
    can_see_price BOOLEAN DEFAULT FALSE
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
);


-- Function to generate generic ID
CREATE OR REPLACE FUNCTION generate_generic_id()
RETURNS TRIGGER AS $$
DECLARE
    date_part VARCHAR(8);
    seq_id INT;
    seq_name TEXT;
BEGIN
    -- Get the date part in YYYYMMDD format
    date_part := TO_CHAR(NEW.created_at, 'YYYYMMDD');
    
    -- Dynamically construct the sequence name based on the table name
    seq_name := TG_TABLE_NAME || '_id_seq';
    
    -- Get the next value from the dynamically constructed sequence name
    EXECUTE 'SELECT NEXTVAL(''' || seq_name || ''')' INTO seq_id;
    
    -- Combine the date part with the sequence ID without leading zeros
    NEW.id := date_part || seq_id::TEXT;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Generate sequences and triggers for tables with id as primary key
DO $$
DECLARE
    table_record RECORD;
    sequence_name TEXT;
    trigger_sql TEXT;
    sequence_sql TEXT;
BEGIN
    FOR table_record IN 
        SELECT tc.table_name 
        FROM information_schema.table_constraints tc
        JOIN information_schema.constraint_column_usage AS ccu
        USING (constraint_schema, constraint_name)
        WHERE tc.constraint_type = 'PRIMARY KEY' 
        AND ccu.column_name = 'id' 
        AND tc.table_schema = 'public'
    LOOP
        -- Generate the sequence name based on the table name
        sequence_name := table_record.table_name || '_id_seq';

        -- Create the sequence for the table
        sequence_sql := 'CREATE SEQUENCE IF NOT EXISTS ' || quote_ident(sequence_name) || ';';
        EXECUTE sequence_sql;

        -- Create the trigger for the table
        trigger_sql := 'CREATE TRIGGER before_insert_' || quote_ident(table_record.table_name) ||
                       ' BEFORE INSERT ON ' || quote_ident(table_record.table_name) ||
                       ' FOR EACH ROW EXECUTE FUNCTION generate_generic_id();';
        EXECUTE trigger_sql;
    END LOOP;
END $$;
