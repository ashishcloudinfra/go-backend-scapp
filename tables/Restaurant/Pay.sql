CREATE TABLE Payment (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), -- Unique identifier for the payment
  amount NUMERIC(10, 2) NOT NULL,                -- Total amount paid
  currency VARCHAR(3) NOT NULL DEFAULT 'INR',    -- Currency code (e.g., INR)
  paymentMethod VARCHAR(50) NOT NULL,            -- Payment method (e.g., Credit Card, UPI)
  paymentStatus VARCHAR(20) NOT NULL,            -- Status of the payment (e.g., Success, Failed)
  paymentDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp of the payment
  transactionId VARCHAR(255),                    -- Transaction ID from the payment gateway
  customerPhone VARCHAR(15),                     -- Customer's phone for verification
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Record creation timestamp
  updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  companyId UUID REFERENCES company(id) ON DELETE CASCADE
);

CREATE TABLE PaymentOrders (
  paymentId UUID REFERENCES Payment(id) ON DELETE CASCADE,
  orderId UUID REFERENCES Orders(id) ON DELETE CASCADE,
  PRIMARY KEY (paymentId, orderId)
);
