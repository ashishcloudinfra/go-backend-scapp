drop table if exists membershipPlan;

CREATE TABLE membershipPlan (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(255) NOT NULL,
  price NUMERIC(10, 2) NOT NULL,
  duration VARCHAR(50) NOT NULL CHECK (duration IN ('monthly', 'quarterly', 'annually')),
  discount VARCHAR(255),
  features TEXT[],
  status VARCHAR(255) NOT NULL,
  cancellation_policy TEXT[],
  CONSTRAINT unique_name_duration UNIQUE (name, duration),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  companyId UUID NOT NULL REFERENCES company(id) ON DELETE CASCADE
);
