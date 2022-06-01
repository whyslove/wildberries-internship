CREATE TABLE IF NOT EXISTS event(
    id SERIAL PRIMARY KEY,
    user_id INT,
    date DATE,
    descr VARCHAR(255)
);


-- create or replace procedure add_event(
--     id INT,
--     user_id INT,
--     date DATE,
--     descr VARCHAR(255)
-- )
-- LANGUAGE plpgsql    
-- AS $$
-- BEGIN
--     INSERT INTO event  
--     VALUES
--         (id, user_id, date, descr);
-- END$$;


-- create or replace procedure update_event(
--     id INT,
--     user_id INT,
--     date DATE,
--     descr VARCHAR(255)
-- )
-- LANGUAGE plpgsql    
-- AS $$
-- BEGIN
--     UPDATE event  
--     SET
--         event.id = id, event.user_id = user_id, event.date = date, event.descr = descr
--     WHERE event.id = id;
-- END$$;


-- create or replace procedure delete_event(
--     id INT
-- )
-- LANGUAGE plpgsql    
-- AS $$
-- BEGIN
--     DELETE FROM event  
--     WHERE event.id = id;
-- END$$;


