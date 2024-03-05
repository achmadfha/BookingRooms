-- Insert data into employee table
INSERT INTO employee (employee_id, full_name, division, phone_number, position, username, password)
VALUES
    ('11111111-1111-1111-1111-111111111111', 'John Doe', 'Marketing', '1234567890', 'ADMIN', 'admin', '$2a$12$kFwAXX9ZyMrgl.NQcqkT6.5pdhippVeiP5lDUTpPI897FO6cK4NOe'),
    ('22222222-2222-2222-2222-222222222222', 'Jane Smith', 'HR', '9876543210', 'EMPLOYEE', 'employee', '$2a$12$1UJ.J1qySbNdL/cOUeiuTO6gm9mZvj.VVPtaQ/J8xxhuPgkhzc79y'),
    ('33333333-3333-3333-3333-333333333333', 'Michael Johnson', 'Finance', '5556667777', 'GA', 'globaladmin', '$2a$12$RyechOUkkmKFde5OvwiMKOHsVGkm4zKW703TMhcSA4z9tsAKM8iPa');

-- Insert data into room_details table
INSERT INTO room_details (room_details_id, room_type, capacity, facility)
VALUES
    ('44444444-4444-4444-4444-444444444444', 'Meeting Room', 2, ARRAY['Projector', 'Whiteboard', 'Conference Phone']),
    ('55555555-5555-5555-5555-555555555555', 'Office Cubicle', 4, ARRAY['Desk', 'Chair', 'Computer']),
    ('66666666-6666-6666-6666-666666666666', 'Conference Room', 3, ARRAY['Projector', 'Whiteboard', 'Conference Table']);

-- Insert data into room table
INSERT INTO room (room_id, room_details_id, name, status)
VALUES
    ('77777777-7777-7777-7777-777777777777', '44444444-4444-4444-4444-444444444444', 'Room 101', 'AVAILABLE'),
    ('88888888-8888-8888-8888-888888888888', '55555555-5555-5555-5555-555555555555', 'Room 201', 'BOOKED'),
    ('99999999-9999-9999-9999-999999999999', '66666666-6666-6666-6666-666666666666', 'Room 301', 'AVAILABLE');

-- Insert data into transactions table
INSERT INTO transactions (transaction_id, employee_id, room_id, start_date, end_date, description)
VALUES
    ('12345678-1234-1234-1234-123456789012', '11111111-1111-1111-1111-111111111111', '77777777-7777-7777-7777-777777777777', '2024-03-05', '2024-03-07', 'Meeting with clients'),
    ('23456789-2345-2345-2345-234567890123', '22222222-2222-2222-2222-222222222222', '88888888-8888-8888-8888-888888888888', '2024-03-08', '2024-03-10', 'Team building event'),
    ('34567890-3456-3456-3456-345678901234', '33333333-3333-3333-3333-333333333333', '99999999-9999-9999-9999-999999999999', '2024-03-11', '2024-03-13', 'Training session');

-- Insert data into transaction_logs table
INSERT INTO transaction_logs (transaction_log_id, transaction_id, approved_by, approval_status, description)
VALUES
    ('98765432-9876-9876-9876-987654321098', '12345678-1234-1234-1234-123456789012', '22222222-2222-2222-2222-222222222222', 'ACCEPT', 'Approved meeting request'),
    ('87654321-8765-8765-8765-876543210987', '23456789-2345-2345-2345-234567890123', '11111111-1111-1111-1111-111111111111', 'DECLINE', 'Event postponed'),
    ('76543210-7654-7654-7654-765432109876', '34567890-3456-3456-3456-345678901234', '33333333-3333-3333-3333-333333333333', 'PENDING', 'Awaiting approval');
