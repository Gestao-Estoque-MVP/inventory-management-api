CREATE TABLE companies (
  id UUID PRIMARY KEY,
  name VARCHAR(255),
  document VARCHAR(20),
  address_id UUID NOT NULL REFERENCES address(id),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);