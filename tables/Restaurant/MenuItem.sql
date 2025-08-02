drop table if exists MenuItemCategory;
drop table if exists MenuItemPricing;
drop table if exists MenuItem;

CREATE TABLE MenuItemCategory (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(50) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  companyId UUID REFERENCES company(id) ON DELETE CASCADE,
  UNIQUE (name, companyId)
);

CREATE TABLE MenuItem (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(100) NOT NULL,
  description TEXT,
  cookingTime VARCHAR(255),
  photo VARCHAR(255),
  isVeg BOOLEAN NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  categoryId UUID REFERENCES MenuItemCategory(id) ON DELETE CASCADE,
  companyId UUID REFERENCES company(id) ON DELETE CASCADE
);

CREATE TABLE MenuItemPricing (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  varietyType VARCHAR(255) NOT NULL,
  price VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  menuItemId UUID REFERENCES MenuItem(id) ON DELETE CASCADE,
);

CREATE TABLE AddOnCategory (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(50) NOT NULL,
  canSelectMultiple BOOLEAN NOT NUll,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  companyId UUID REFERENCES company(id) ON DELETE CASCADE,
  UNIQUE (name, companyId)
);

CREATE TABLE AddOnItem (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(100) NOT NULL,
  description TEXT,
  isVeg BOOLEAN NOT NULL,
  price VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  categoryId UUID REFERENCES AddOnCategory(id) ON DELETE CASCADE,
  companyId UUID REFERENCES company(id) ON DELETE CASCADE
);

CREATE TABLE MenuItemAddOn (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  menuItemId UUID REFERENCES MenuItem(id) ON DELETE CASCADE,
  addOnCategoryId UUID REFERENCES AddOnCategory(id) ON DELETE CASCADE,
  companyId UUID REFERENCES company(id) ON DELETE CASCADE,
  UNIQUE (menuItemId, addOnCategoryId, companyId)
);
