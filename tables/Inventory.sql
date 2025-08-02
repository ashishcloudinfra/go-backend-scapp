drop table if exists Item;
drop table if exists Equipment;

CREATE TABLE Equipment (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name TEXT UNIQUE NOT NULL,
  img TEXT NOT NULL,
  type TEXT NOT NULL,
  instructions TEXT NOT NULL,
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Item (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  status TEXT NOT NULL CHECK (status IN ('available', 'maintenance', 'decommissioned')),
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  equipmentId UUID REFERENCES Equipment(id) ON DELETE CASCADE,
  companyId UUID NOT NULL REFERENCES company(id) ON DELETE CASCADE
);
