drop table if exists EventParticipant;
drop table if exists ScheduledEvent;
drop table if exists Event;
drop table if exists EventRoom;

CREATE TABLE EventRoom (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name TEXT NOT NULL,
  capacity INTEGER CHECK (capacity > 0),
  location TEXT,
  isUnderMaintenance BOOLEAN DEFAULT FALSE,
  startTime TEXT,
  endTime TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  companyId UUID NOT NULL REFERENCES company(id) ON DELETE CASCADE
);

CREATE TABLE Event (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name TEXT NOT NULL,
  description TEXT,
  organiserId UUID NOT NULL REFERENCES iam(id) ON DELETE CASCADE,
  eventRoomId UUID,
  startDate TEXT NOT NULL,
  endDate TEXT NOT NULL,
  startTime TEXT NOT NULL,
  endTime TEXT NOT NULL,
  isAllDayEvent BOOLEAN DEFAULT FALSE,
  isRecurring BOOLEAN DEFAULT FALSE,
  recurrenceType TEXT,
  status VARCHAR(255) NOT NULL CHECK (status IN ('review', 'authorized')),
  metadata JSON NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  companyId UUID NOT NULL REFERENCES company(id) ON DELETE CASCADE
);

CREATE TABLE ScheduledEvent (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  startDate DATE NOT NULL,
  endDate DATE NOT NULL,
  startTime TIME NOT NULL,
  endTime TIME NOT NULL,
  eventId UUID NOT NULL REFERENCES Event(id) ON DELETE CASCADE
);

CREATE TABLE EventParticipant (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  hasRegistered BOOLEAN DEFAULT FALSE,
  hasAttended BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  scheduledEventId UUID REFERENCES ScheduledEvent(id) ON DELETE CASCADE,
  attendeeId UUID NOT NULL REFERENCES iam(id) ON DELETE CASCADE
);
