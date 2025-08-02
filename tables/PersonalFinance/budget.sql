drop table if exists BudgetCategoryType;
drop table if exists BudgetItem;
drop table if exists BudgetCategory;

CREATE TABLE BudgetCategoryType (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  type VARCHAR(255) UNIQUE NOT NULL,
  bgColor VARCHAR(255) NOT NULL,
  textColor VARCHAR(255) NOT NULL
);

CREATE TABLE BudgetCategory (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  categoryName VARCHAR(100) NOT NULL,
  categoryDescription VARCHAR(255),
  month INT NOT NULL,
  year INT NOT NULL,
  categoryTypeId UUID NOT NULL REFERENCES BudgetCategoryType(id) ON DELETE CASCADE ON UPDATE CASCADE,
  iamId UUID NOT NULL REFERENCES iam(id) ON DELETE CASCADE,
  parentId UUID REFERENCES BudgetCategory(id) ON DELETE CASCADE ON UPDATE CASCADE,

  -- Add a UNIQUE constraint for (categoryName, month, year)
  CONSTRAINT unique_categoryName_month_year_iamId
    UNIQUE (categoryName, month, year, iamId)
);

CREATE TABLE BudgetItem (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  categoryId UUID REFERENCES BudgetCategory(id) ON DELETE CASCADE ON UPDATE CASCADE,
  itemName VARCHAR(255) NOT NULL,
  description TEXT,
  allocatedAmount TEXT,
  actualAmount TEXT,
  currencyCode TEXT NOT NULL DEFAULT 'USD',
  status VARCHAR(20) DEFAULT 'active',
  month INT NOT NULL,
  year INT NOT NULL,
  iamId UUID NOT NULL REFERENCES iam(id) ON DELETE CASCADE,
  createdAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE AssetType (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(255) NOT NULL
);

CREATE TABLE AssetItem (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(255) NOT NULL,
  assetType VARCHAR(255) NOT NULL,
  code VARCHAR(255) DEFAULT '',
  avgBuyValue TEXT,
  currentValue TEXT,
  totalUnits TEXT,
  pctIncPerYear TEXT,
  iamId UUID NOT NULL REFERENCES iam(id) ON DELETE CASCADE,
  createdAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
