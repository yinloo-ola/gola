package structs

/*
desc users;
+------------+--------------+------+------+---------+----------------+
| Field      | Type         | Null | Key  | Default | Extra          |
+------------+--------------+------+------+---------+----------------+
| id         | int          | NO   | PRI  |         | auto_increment |
| name       | varchar(255) | NO   |      | ""      |                |
| email      | varchar(255) | NO   | MUL  | ""      |                |
| created_at | int unsigned | NO   |      | "0"     |                |
| updated_at | int unsigned | NO   |      | "0"     |                |
+------------+--------------+------+------+---------+----------------+
*/

type RowDesc struct {
	Field, Type, Null, Key, Default, Extra string
}

/*
show index from users;
+-------+------------+----------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+---------+------------+
| Table | Non_unique | Key_name | Seq_in_index | Column_name | Collation | Cardinality | Sub_part | Packed | Null | Index_type | Comment | Index_comment | Visible | Expression |
+-------+------------+----------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+---------+------------+
| users |          0 | email    |            1 | email       | NULL      |           0 |     NULL | NULL   |      | BTREE      |         |               | YES     | NULL       |
| users |          1 | name     |            1 | name        | NULL      |           0 |     NULL | NULL   |      | BTREE      |         |               | YES     | NULL       |
+-------+------------+----------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+---------+------------+
*/

type IndexDesc struct {
	Table, Key_name, Column_name, Collation, Sub_part, Packed, Null, Index_type, Comment, Index_comment, Visible, Expression string
	Non_unique, Seq_in_index, Cardinality                                                                                    int
}

func filterBy[T any](items []*T, isNeeded func(item *T) bool) []*T {
	result := make([]*T, 0, len(items))

	for _, item := range items {
		if isNeeded(item) {
			result = append(result, item)
		}
	}

	return result
}