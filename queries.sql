SELECT t1.task_id,
        EVERY(CASE WHEN t2.completed IS NULL THEN TRUE
                   ELSE t2.completed
              END)
FROM tasks AS t1
LEFT JOIN task_dependencies AS td ON t1.task_id = td.task_id
LEFT JOIN tasks AS t2 ON td.depends_on_task_id = t2.task_id
GROUP BY t1.task_id;


SELECT td.task_id FROM tasks AS dependent WHERE dependent.task_id = 1
JOIN task_dependencies as td ON dependent.task_id = td.depends_on_task_id