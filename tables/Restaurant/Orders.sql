drop table if exists OrderItem;
drop table if exists Orders;

CREATE TABLE Orders (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  orderNumber SERIAL NOT NULL,
  tableNumber INT NOT NULL,
  isDineIn BOOLEAN NOT NULL,
  status VARCHAR(50) NOT NULL,
  phone VARCHAR(15),
  email VARCHAR(255),
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  companyId UUID REFERENCES company(id) ON DELETE CASCADE,
  UNIQUE (orderNumber, companyId)
);

CREATE TABLE OrderItem (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  quantity INT NOT NULL DEFAULT 1,
  menuItemPricingId UUID REFERENCES MenuItemPricing(id) ON DELETE CASCADE,
  orderId UUID REFERENCES Orders(id) ON DELETE CASCADE
);
