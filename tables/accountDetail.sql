drop table if exists accountDetail;

CREATE TABLE accountDetail (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  role VARCHAR(50) NOT NULL,
  permissions TEXT NOT NULL,
  plans TEXT NOT NULL,
  status VARCHAR(50),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  companyId UUID REFERENCES company(id) ON DELETE CASCADE,
  iamId UUID NOT NULL REFERENCES iam(id) ON DELETE CASCADE
);
