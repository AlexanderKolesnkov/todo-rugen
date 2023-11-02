CREATE TABLE tasks (
                       id INTEGER NOT NULL PRIMARY KEY,
                       title VARCHAR(100) NOT NULL,
                       content TEXT NOT NULL,
                       created TIMESTAMP NOT NULL,
                       status BOOLEAN
);

CREATE INDEX idx_tasks_created ON tasks(created);

INSERT INTO tasks(id, title, content, created, status) VALUES (
                                                                  COALESCE((SELECT MAX(id+1) FROM tasks), 1),
                                                                  'Сделать приложение для финансов',
                                                                  E'Написать на golang приложение для доходов и расходов!',
                                                                  CURRENT_TIMESTAMP(0),
                                                                  FALSE
                                                              );

INSERT INTO tasks(id, title, content, created, status) VALUES (
                                                                  COALESCE((SELECT MAX(id+1) FROM tasks), 1),
                                                                  'Пройти Zelda TotK',
                                                                  E'Пройти уже наконецто Зельду!',
                                                                  CURRENT_TIMESTAMP(0),
                                                                  FALSE
                                                              );

INSERT INTO tasks(id, title, content, created, status) VALUES (
                                                                  COALESCE((SELECT MAX(id+1) FROM tasks), 1),
                                                                  'Устроиться на работу Golang разработчиком',
                                                                  E'Выучить Golang на базавом уровне и найти работу golang разработчиком,\n чтобы быстрее прокачивать скилы',
                                                                  CURRENT_TIMESTAMP(0),
                                                                  FALSE
                                                              );