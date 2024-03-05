-- Define custom ENUM types
CREATE TYPE position_employee AS ENUM('ADMIN', 'EMPLOYEE', 'GA');
CREATE TYPE room_status AS ENUM('AVAILABLE', 'BOOKED');
CREATE TYPE transaction_status AS ENUM('PENDING', 'ACCEPT', 'DECLINE');

-- Table for employee data
CREATE TABLE employee
(
    employee_id  UUID PRIMARY KEY,
    full_name    VARCHAR(250),
    division     VARCHAR(250),
    phone_number VARCHAR(20),
    position     position_employee,
    username     VARCHAR(250),
    password     VARCHAR(250)
);

-- Table for room details
CREATE TABLE room_details
(
    room_details_id UUID PRIMARY KEY,
    room_type       VARCHAR(250),
    capacity        INTEGER,
    facility        TEXT[]
);

-- Table for room data
CREATE TABLE room
(
    room_id         UUID PRIMARY KEY,
    room_details_id UUID,
    name            VARCHAR(250),
    status          room_status,
    FOREIGN KEY (room_details_id) REFERENCES room_details (room_details_id)
);

-- Table for transactions
CREATE TABLE transactions
(
    transaction_id UUID PRIMARY KEY,
    employee_id    UUID,
    room_id        UUID,
    start_date     DATE NOT NULL,
    end_date       DATE NOT NULL,
    description    VARCHAR(250) NOT NULL,
    status         transaction_status DEFAULT 'PENDING', -- Defaulting to PENDING
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (employee_id) REFERENCES employee (employee_id),
    FOREIGN KEY (room_id) REFERENCES room (room_id)
);

CREATE TABLE transaction_logs
(
    transaction_log_id UUID PRIMARY KEY,
    transaction_id UUID,
    approved_by UUID,
    approval_status transaction_status, -- Defaulting to PENDING
    description varchar(100), -- description for reject or accept
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (transaction_id) REFERENCES transactions (transaction_id),
    FOREIGN KEY (approved_by) REFERENCES employee (employee_id)
);