drop table if exists personalDetail;

CREATE TABLE personalDetail (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  firstName VARCHAR(255) NOT NULL,
  lastName VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  phone VARCHAR(255) NOT NULL,
  address VARCHAR(255),
  city VARCHAR(255),
  state VARCHAR(255),
  zip VARCHAR(255),
  country VARCHAR(255),
  dob VARCHAR(255),
  gender VARCHAR(255),
  metadata TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  iamId UUID NOT NULL REFERENCES iam(id) ON DELETE CASCADE
);
